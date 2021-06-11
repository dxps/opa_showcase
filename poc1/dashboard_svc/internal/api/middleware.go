package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/domain"
	"github.com/pascaldekloe/jwt"
)

// processJWT is processing the JWT, if present in the "Authorization" header,
// and it sets the `domain.Subject` with the ID based on 'sub' JWT claim.
func (api *API) processJWT(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authzHeader := r.Header.Get("Authorization")
		if authzHeader == "" {
			r = api.contextSetSubject(r, domain.AnonymousSubject)
			next.ServeHTTP(w, r)
			return
		}
		headerParts := strings.Split(authzHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			api.invalidAuthenticationTokenResponse(w, r)
			return
		}

		token := headerParts[1]
		// api.logger.Println("[api processJWT] token =", token)

		// Minimal token validation, for now.

		// For now, we won't perform any signing validation
		// (using jwt.ECDSACheck([]byte(token), publicKey) function).

		claims, err := jwt.ParseWithoutCheck([]byte(token))
		if err != nil {
			api.invalidAuthenticationTokenResponse(w, r)
			return
		}
		if !claims.Valid(time.Now()) {
			api.invalidAuthenticationTokenResponse(w, r)
			return
		}

		// Get subject attributes ...
		subject := domain.Subject{ID: claims.Subject}

		// Add the subject to the request context and continue as normal.
		r = api.contextSetSubject(r, &subject)
		next.ServeHTTP(w, r)
	})
}

// func (api *API) checkAuthorization(code string, next http.HandlerFunc) http.HandlerFunc {
func (api *API) authorizeByRule(ruleName string, next http.HandlerFunc) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {

		subjectID := api.contextGetSubject(r)

		allowed, err := api.authz.QueryDecision(ruleName, subjectID)

		if err != nil {
			api.serverErrorResponse(w, r, err)
			return
		}
		if !allowed {
			api.notPermittedResponse(w, r)
		}

		subj := api.contextGetSubject(r)

		api.logger.Printf("[api.authorizeByRule] %+v", *subj)

		next.ServeHTTP(w, r)
	}

	return fn
}
