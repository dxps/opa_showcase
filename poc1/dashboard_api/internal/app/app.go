package app

import (
	"log"
)

type App struct {
	Config  Config
	Logger  *log.Logger
	Version string
}

func New(config Config, logger *log.Logger, version string) *App {
	return &App{
		Config:  config,
		Logger:  logger,
		Version: version,
	}
}
