package api

import (
	"context"
	"net/http"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/domain"
)

type contextKey string

const subjectCtxKey = contextKey("user")

func (api *API) contextSetSubject(r *http.Request, subj *domain.Subject) *http.Request {
	ctx := context.WithValue(r.Context(), subjectCtxKey, subj)
	return r.WithContext(ctx)
}

func (api *API) contextGetSubject(r *http.Request) *domain.Subject {
	subj, ok := r.Context().Value(subjectCtxKey).(*domain.Subject)
	if !ok {
		panic("missing subject value in request context")
	}

	return subj
}
