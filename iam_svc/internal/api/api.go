package api

import (
	"log"

	"github.com/dxps/opa_showcase/iam_svc/internal/app"
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
