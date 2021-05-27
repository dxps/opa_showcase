package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/dxps/opa_showcase/authz_svc/internal/api"
	"github.com/dxps/opa_showcase/authz_svc/internal/opa"
	"github.com/dxps/opa_showcase/authz_svc/internal/version"
	"github.com/sirupsen/logrus"
)

var configFile = flag.String("config", "", "set the OPA config file to load")
var verbose = flag.Bool("verbose", false, "enable verbose logging")
var versionFlag = flag.Bool("version", false, "print version and exit")

func main() {

	flag.Parse()

	if *versionFlag {
		fmt.Println("Version:", version.Version)
		fmt.Println("Vcs:", version.Vcs)
		os.Exit(0)
	}

	setupLogging()

	engine, err := opa.New(opa.Config(*configFile))
	if err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Failed to initialize OPA.")
	}

	ctx := context.Background()

	if err := engine.Start(ctx); err != nil {
		logrus.WithFields(logrus.Fields{"err": err}).Fatal("Failed to start OPA.")
	}

	if err := api.New(engine).Run(ctx); err != nil {
		logrus.Fatal(err)
	}

	logrus.Info("Shutting down.")
}

func setupLogging() {
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	logLevel := logrus.InfoLevel
	if *verbose {
		logLevel = logrus.DebugLevel
	}
	logrus.SetLevel(logLevel)
}
