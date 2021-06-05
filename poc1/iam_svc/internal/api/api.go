package api

import (
	"log"
	"net/http"

	"github.com/dxps/opa_showcase/iam_svc/internal/app"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	config     app.Config
	logger     *log.Logger
	appVersion string
}

func NewAPI(config app.Config, logger *log.Logger, appVersion string) *API {
	return &API{
		config, logger, appVersion,
	}
}

func (api *API) Routes() *httprouter.Router {

	router := httprouter.New()

	// Registering the handlers per methods and URL patterns.
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.HealthcheckHandler)

	return router
}
