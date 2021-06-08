package domain

// Subject represents a human individual or a service (machine) account
// that is registered within and owned by the IAM service.
type Subject struct {
	ID         string       // The (external, UUID based) identifier.
	Attributes []*Attribute // The attributes of the subject.
}

var AnonymousSubject = &Subject{}
