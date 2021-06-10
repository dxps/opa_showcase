package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *API) Routes() http.Handler {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(api.notFoundResponse)

	// Registering the handlers per methods and URL patterns.

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.healthcheckHandler)

	router.HandlerFunc(http.MethodGet,
		"/v1/subjects/:id/portfolio",
		api.authorizeByRule("products_enablement:subject_has_product", api.getSubjectPortfolio),
	)

	return api.processJWT(router)
}
