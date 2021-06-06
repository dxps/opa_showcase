package api

import (
	"log"
	"net/http"

	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/infra/repos"
	"github.com/julienschmidt/httprouter"
)

type API struct {
	config     app.Config
	logger     *log.Logger
	appVersion string
	repos      repos.Repos
}

func NewAPI(config app.Config, logger *log.Logger, appVersion string, repos repos.Repos) *API {
	return &API{
		config, logger, appVersion, repos,
	}
}

func (api *API) Routes() *httprouter.Router {

	router := httprouter.New()

	// Registering the handlers per methods and URL patterns.

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", api.HealthcheckHandler)

	router.HandlerFunc(http.MethodPost, "/v1/subjects", api.RegisterUserHandler)

	return router
}
