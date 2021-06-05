package app

type Config struct {
	Port     int    // The HTTP listening port.
	EnvStage string // The environment stage (dev|qa|prod).
	Db       struct {
		DSN string // The data source name.
	}
}
