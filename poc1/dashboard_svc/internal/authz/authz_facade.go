package authz

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Jeffail/gabs/v2"
	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/app"
)

type AuthzFacade struct {
	config           authzFacadeConfig
	policies         []Policy          // The policies being used.
	Rules            map[string]*Rule  // The rules that are being referred to in the PEP places, mapped by name.
	inputContextData map[string]string // Common (app wide) `input.context` data, that policies may refer to.
	logger           *log.Logger       // For its own logging needs.
}

type authzFacadeConfig struct {
	productName               string // Application/service (UU)ID.
	subjectAttributesFetchURL string // URL for GETting the attributes of a subject.
	policiesFetchURL          string // URL for GETting the policies used by this app.
	policyAgentURL            string // URL of OPA's REST API.
}

// NewAuthzFacade is the facade object, supporting the authorization processing needs
// in the PEP (Policy Enforcement Point) places.
func NewAuthzFacade(
	productName string,
	policiesFetchURL string,
	policyAgentURL string,
	subjectAttributesFetchURL string,
	logger *log.Logger,
) (*AuthzFacade, error) {

	a := AuthzFacade{
		config: authzFacadeConfig{
			productName:               productName,
			policiesFetchURL:          policiesFetchURL,
			policyAgentURL:            policyAgentURL,
			subjectAttributesFetchURL: subjectAttributesFetchURL,
		},
		inputContextData: make(map[string]string),
		logger:           logger,
	}
	// Initializations of policies and rules.
	if err := a.initPolicies(); err != nil {
		return nil, err
	}
	a.initRules()
	// a.logRules()

	return &a, nil
}

// Initializing the policies by fetching them from PAP
// (Policy Administration Point) and keep them in memory.
func (a *AuthzFacade) initPolicies() error {

	// Simulating the policies fetching and getting back this result:
	policies := []Policy{
		{
			ID:        "001",
			Name:      "products_enablement",
			QueryPath: "products_enablement",
			Version:   "202106091735",
			Rules: []Rule{
				{
					Name: "subject_has_product",
					InputSpec: &InputSpec{
						SubjectAttrs: []string{"products"},
						ContextAttrs: []string{"product"},
					},
				},
			},
		},
		{
			ID:        "002",
			Name:      "rbac",
			QueryPath: "rbac",
			Version:   "202106100018",
			Rules: []Rule{
				{
					Name: "subject_has_support_role",
					InputSpec: &InputSpec{
						SubjectAttrs: []string{"memberOf"},
						ContextAttrs: []string{"product"},
					},
				},
			},
		},
	}
	a.policies = policies
	return nil
}

// Initializing the rules, based on the fetched policies.
func (a *AuthzFacade) initRules() {

	a.Rules = make(map[string]*Rule, len(a.policies))
	for _, p := range a.policies {
		for _, r := range p.Rules {
			ar := Rule{
				Name:      fmt.Sprintf("%s:%s", p.Name, r.Name),
				QueryURL:  fmt.Sprintf("%s/v1/data/%s/%s", a.config.policyAgentURL, p.QueryPath, r.Name),
				Policy:    &p,
				InputSpec: r.InputSpec,
			}
			a.Rules[ar.Name] = &ar
		}
	}
}

func (a *AuthzFacade) logRules() {

	out := "[authz.AuthzFacade] Loaded rules: [ "
	for _, rule := range a.Rules {
		ruleJson, _ := json.Marshal(rule)
		out = fmt.Sprintf("%s \n\t %+v, ", out, string(ruleJson))
	}
	out = fmt.Sprintf("%s]\n", out)
	a.logger.Println(out)
}

// SetInputContextData is setting data into the common (app wide) `input.context` part, that policies may refer to.
func (a *AuthzFacade) SetInputContextData(name, value string) {
	a.inputContextData[name] = value
}

// QueryDecision does a query to the policy agent for getting an authorization decision.
// In this case, just the rule name and any common and previously set input context data is being used.
func (a *AuthzFacade) QueryDecision(ruleName string, subjectID SubjectID) (bool, error) {

	r, ok := a.Rules[ruleName]
	if !ok {
		a.logger.Printf("[authz.QueryDecision] rule named '%s' not found.", ruleName)
		return false, app.ErrNotFound
	}

	input, err := a.buildInput(subjectID)
	if err != nil {
		a.logger.Println("[authz.QueryDecision] buildInput error:", err)
		return false, err
	}

	res, err := a.queryOPA(r.QueryURL, input)
	if err != nil {
		a.logger.Println("[authz.QueryDecision] queryOPA error:", err)
		return false, err
	}
	return res.Result, nil
}

// buildInput is building the input for the Rule that is being evaluated as part of the query.
func (a *AuthzFacade) buildInput(subjectID SubjectID) (Input, error) {

	input := NewInput()

	if err := a.addSubjectAttributes(input, subjectID); err != nil {
		return nil, err
	}
	if err := a.addContextAttributes(input); err != nil {
		return nil, err
	}

	return input, nil
}

func (a *AuthzFacade) addSubjectAttributes(input Input, subjectID SubjectID) error {

	inputJson, _ := input.(*gabs.Container) // The concrete type being used.
	_, _ = inputJson.Array("input", "subject", "products")
	_ = inputJson.ArrayAppend("product_1", "input", "subject", "products")
	_ = inputJson.ArrayAppend("dashboard", "input", "subject", "products")

	_, _ = inputJson.Array("input", "subject", "memberOf")
	_ = inputJson.ArrayAppend("officer@product_1", "input", "subject", "memberOf")
	_ = inputJson.ArrayAppend("support@dashboard", "input", "subject", "memberOf")

	return nil
}

func (a *AuthzFacade) addContextAttributes(input Input) error {

	inputJson, _ := input.(*gabs.Container) // The concrete type being used.
	_, _ = inputJson.Set(a.config.productName, "input", "context", "product")

	return nil
}

// queryOPA contacts the OPA on its REST API, querying it based on that `url` path and `input`.
func (a *AuthzFacade) queryOPA(url string, input Input) (*QueryResult, error) {

	inputJson, ok := input.(*gabs.Container) // The concrete type being used.
	if !ok {
		return nil, app.ErrInternal
	}
	inputStr := inputJson.String()
	a.logger.Println("[authz.queryOPA] inputStr:", inputStr)

	postBody := bytes.NewBuffer([]byte(inputStr))
	resp, err := http.Post(url, "application/json", postBody)
	if err != nil {
		return nil, err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	qr := &QueryResult{}
	if err := json.Unmarshal(respBody, qr); err != nil {
		return nil, err
	}
	return qr, nil
}
