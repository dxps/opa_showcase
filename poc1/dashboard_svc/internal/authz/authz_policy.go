package authz

// RuleInputSpec specifies the attributes of `subject` and `context`
// that are being used within an `AuthzRule` for evaluation.
type RuleInputSpec struct {
	SubjectAttributes []string // The attributes of `subject` part of the query `input`.
	ContextAttributes []string // The attributes of `context` part of the query `input`.
}

// RuleInputData is the concrete input that is passed to a Rule for evaluating it.
type RuleInputData struct {
	SubjectData map[string]string
	ContextData map[string]string
}

// Rule represents an authorization rule,
// whose name is being referred when setting up the PEPs.
type Rule struct {

	// The name of the rule, part of a policy, being referred to on PEPs config.
	// This is being constructed based on the policy name and rule name within the policy.
	Name string

	// The path used for querying to get the decision.
	QueryPath string

	// The policy that this rule is part of.
	Policy *Policy

	Input RuleInputSpec
}

// Policy is the core of the PBAC model.
// It includes one or more Rules. And evaluations of these rules
// are used for rendering the authorization decisions.
type Policy struct {
	ID        string
	Name      string
	QueryPath string
	Rules     []Rule
	Version   string
}
