package main

import (
	"flag"
	"log"
	"os"

	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/api"
	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/app"
	"github.com/dxps/opa_showcase/poc1/dashboard_svc/internal/authz"
)

const (
	APP_NAME    = "dashboard" // The application/service identifier.
	APP_VERSION = "1.0.0-DEV" // The application/service version. At build time, it gets a different value.
)

func main() {

	var cfg app.Config

	// CLI Flags definition and parsing.
	flag.IntVar(&cfg.Port, "port", 3002, "HTTP Listening Port of the API Server")
	flag.Parse()

	// Logger init: sending the entries to standard output.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	policiesFetchURL := "_"
	policyAgentURL := "http://localhost:8181"
	subjectAttributesFetchURL := "_"

	authz, err := authz.NewAuthzFacade(APP_NAME, policiesFetchURL, policyAgentURL, subjectAttributesFetchURL, logger)
	if err != nil {
		logger.Println("[main] Startup error, cannot init the Authz:", err)
		return
	}

	api := api.NewAPI(cfg, authz, logger, APP_VERSION)

	err = api.Serve()
	logger.Fatal(err)
}
