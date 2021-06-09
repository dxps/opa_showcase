package authz

// AuthzRuleInput specifies the attributes of `subject` and `data`
// that are being used within an `AuthzRule`.
type AuthzRuleInput struct {
	SubjectAttributes []string
	DataAttributes    []string
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
