package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/infra/repos"
)

type API struct {
	config         app.Config
	logger         *log.Logger
	appVersion     string
	repos          repos.Repos
	signingKeyPair app.SigningKeyPair
	httpServer     *http.Server
}

func NewAPI(config app.Config, logger *log.Logger, appVersion string, repos repos.Repos, signing app.SigningKeyPair) *API {

	api := API{config, logger, appVersion, repos, signing, nil}
	httpSrv := http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      api.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	api.httpServer = &httpSrv
	return &api
}

func (api *API) Serve() error {

	api.logger.Printf("Listening for HTTP requests on port %s", api.httpServer.Addr)
	return api.httpServer.ListenAndServe()
}

func (api *API) Shutdown(gracefulCtx context.Context) error {

	return api.httpServer.Shutdown(gracefulCtx)
}
