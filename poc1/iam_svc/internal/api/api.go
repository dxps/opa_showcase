package api

import (
	"log"

	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/infra/repos"
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
