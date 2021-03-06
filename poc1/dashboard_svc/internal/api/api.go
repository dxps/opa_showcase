package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/authz"
)

type API struct {
	config     app.Config
	authz      *authz.AuthzFacade
	logger     *log.Logger
	appVersion string
}

func NewAPI(config app.Config, authz *authz.AuthzFacade, logger *log.Logger, appVersion string) *API {

	return &API{config, authz, logger, appVersion}
}

func (api *API) Serve() error {

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", api.config.Port),
		Handler:      api.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	api.logger.Printf("Listening for HTTP requests on port %s", srv.Addr)
	return srv.ListenAndServe()
}
