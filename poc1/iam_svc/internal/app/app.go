package app

import "log"

type App struct {
	Config  Config
	Logger  *log.Logger
	Version string
}
