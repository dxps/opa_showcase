package app

import "errors"

// Common and generic errors.
var (
	ErrSubjectIDInvalid = errors.New("subject's provided id is invalid")
	ErrNotFound         = errors.New("not found")
	ErrInternal         = errors.New("internal error")
)
