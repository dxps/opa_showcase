package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (api *API) Routes() *httprouter.Router {

	router := httprouter.New()

	router.NotFound = http.HandlerFunc(api.notFoundResponse)

	// Registering the handlers per methods and URL patterns.

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.HealthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/subjects", api.RegisterUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/authenticate", api.authenticateHandler)

	return router
}
