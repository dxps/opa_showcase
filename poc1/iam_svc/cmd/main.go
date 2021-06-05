package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"

	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/api"
	"github.com/dxps/opa_showcase/poc1/iam_svc/internal/app"
)

// App version. At build time, it gets a different value.
const APP_VERSION = "1.0.0-DEV"

func main() {

	var cfg app.Config

	// Defining the CLI flags and their default values, plus parsing the startup call.
	flag.IntVar(&cfg.Port, "port", 3001, "HTTP Listening Port of the API Server")
	flag.StringVar(&cfg.EnvStage, "env", "DEV", "Environment stage (DEV|QA|PROD)")
	flag.StringVar(&cfg.Db.DSN, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	// Logger init: sending the entries to standard output.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := app.New(cfg, logger, APP_VERSION)

	if err := app.Init(); err != nil {
		logger.Fatal(err)
	}

	api := api.NewAPI(cfg, logger, APP_VERSION)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", app.Config.Port),
		Handler:      api.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Listening for HTTP requests on port %s", srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
