package app

type Config struct {
	Port int    // The HTTP listening port.
	Env  string // The environment stage (dev|qa|prod).
}
