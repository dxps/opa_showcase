package authz

// AuthzRuleInput specifies the attributes of `subject` and `context`
// that are being used within an `AuthzRule`.
type AuthzRuleInput struct {
	SubjectAttributes []string // The attributes of `subject` part of the query `input`.
	ContextAttributes []string // The attributes of `context` part of the query `input`.
}

// AuthzRule represents an authorization rule,
// whose name is being referred when setting up the PEPs.
type AuthzRule struct {

	// The name of the rule, part of a policy, being referred to on PEPs config.
	// This is being constructed based on the policy name and rule name within the policy.
	Name string

	// The path used for querying to get the decision.
	QueryPath string

	// The policy that this rule is part of.
	Policy *AuthzPolicy

	Input AuthzRuleInput
}

// AuthzPolicy is the core of the PBAC model.
// It includes one or more rules, and evaluations of such rules
// render the authorization decisions.
type AuthzPolicy struct {
	ID        string
	Name      string
	QueryPath string
	Rules     []AuthzRule
	Version   string
}
