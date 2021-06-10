package authz

import "github.com/Jeffail/gabs/v2"

// SubjectID is the initial input about the subject.
// It can be a string or a number.
type SubjectID interface{}

// Input is the data that is passed as the `input.subject` part of a query decision.
type Input interface{} // The concrete type used is *gabs.Container.

func NewInput() Input {
	return gabs.New()
}

// InputContext is the data that is passed as the `input.context` part of a query decision.
type InputContext map[string]string

// InputSpec specifies the attributes of `subject` and `context`
// that are being used within a `Rule` when evaluating it.
type InputSpec struct {
	SubjectAttrs []string // The attributes of `subject` part of the query `input`.
	ContextAttrs []string // The attributes of `context` part of the query `input`.
}

// Rule represents an authorization rule.
// Its name is being referred when setting up the PEPs (Policy Enforcement Points).
type Rule struct {

	// The name of the rule, part of a policy, being referred to on PEPs config.
	// This is being constructed as `<rule-name>@<policy-name>` to be unique.
	Name string

	// The URL (OPA endpoint + path) used for querying for the decision.
	QueryURL string

	// The policy that this rule is part of.
	Policy *Policy `json:"-"`

	// The specification of the input, that is what parts of the `input`
	// are being used during the rule evaluation.
	InputSpec *InputSpec
}

// Policy is at the core of the PBAC model.
// It includes one or more Rules. And evaluations of these rules
// are used for rendering the authorization decisions.
type Policy struct {
	ID        string
	Name      string
	QueryPath string
	Rules     []Rule
	Version   string
}
