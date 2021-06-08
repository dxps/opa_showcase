package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/domain"
	"github.com/pascaldekloe/jwt"
)

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

		api.logger.Println("[api processJWT] token =", token)

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
func (api *API) authorizeOnPolicy(policy string, next http.HandlerFunc) http.HandlerFunc {

	fn := func(w http.ResponseWriter, r *http.Request) {

		subj := api.contextGetSubject(r)

		api.logger.Printf("[api authorizeOnPolicy] %+v", *subj)

		// Handle AuthZ through OPA.

		// permissions, err := app.models.Permissions.GetAllForUser(user.ID)
		// if err != nil {
		// 	api.serverErrorResponse(w, r, err)
		// 	return
		// }

		// if !permissions.Include(code) {
		// 	api.notPermittedResponse(w, r)
		// 	return
		// }

		next.ServeHTTP(w, r)
	}

	return fn
}
