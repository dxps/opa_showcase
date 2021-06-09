package authz

import "fmt"

type AuthzFacade struct {
	config   authzFacadeConfig
	policies []Policy // The policies being used.
	rules    []Rule   // The rules that are being referred to in the PEP places.
}

type authzFacadeConfig struct {
	appID                     string // The application's (UU)ID.
	subjectAttributesFetchURL string // URL for GETting the attributes of a subject.
	policiesFetchURL          string // URL for GETting the policies used by this app.
}

// NewAuthzFacade is the facade object, supporting the authorization processing needs
// in the PEP (Policy Enforcement Point) places.
func NewAuthzFacade(appID string, policiesFetchURL string, subjectAttributesFetchURL string) (*AuthzFacade, error) {

	a := AuthzFacade{
		config: authzFacadeConfig{
			appID:                     appID,
			policiesFetchURL:          policiesFetchURL,
			subjectAttributesFetchURL: subjectAttributesFetchURL,
		},
	}
	// Initializations of policies and rules.
	if err := a.initPolicies(); err != nil {
		return nil, err
	}
	a.initRules()

	return &a, nil
}

// Initializing the policies by fetching them from PAP
// (Policy Administration Point) and keep them in memory.
func (a *AuthzFacade) initPolicies() error {

	// Simulating the policies fetching and having back the result.
	policies := []Policy{
		{
			ID:        "001",
			Name:      "products_enablement",
			QueryPath: "products_enablement/",
			Version:   "202106091735",
			Rules: []Rule{
				{
					Name: "subject_has_product",
					Input: RuleInputSpec{
						SubjectAttributes: []string{"products"},
						ContextAttributes: []string{"product"},
					},
				},
			},
		},
		{
			ID:        "002",
			Name:      "standard_rbac",
			QueryPath: "rbac/",
			Version:   "202106100018",
			Rules: []Rule{
				{
					Name: "subject_is_support",
					Input: RuleInputSpec{
						SubjectAttributes: []string{"memberOf"},
						ContextAttributes: []string{"group"},
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

	a.rules = make([]Rule, len(a.policies))
	for _, p := range a.policies {
		for r := range p.Rules {
			ar := Rule{
				Name:      fmt.Sprintf(p.Name, ":", r),
				QueryPath: fmt.Sprintf(p.QueryPath, "/", r),
				Policy:    &p,
			}
			a.rules = append(a.rules, ar)
		}

	}
}
