package api

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/infra/repos"
)

type API struct {
	config     app.Config
	logger     *log.Logger
	appVersion string
	repos      repos.Repos
	signing    signingPair
}

type signingPair struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

func NewAPI(config app.Config, logger *log.Logger, appVersion string, repos repos.Repos) *API {

	privKey, pubKey := generateECDSAKeys()
	return &API{
		config, logger, appVersion, repos, signingPair{privKey, pubKey},
	}
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
