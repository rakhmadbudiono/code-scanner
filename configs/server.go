package configs

type Server struct {
	Port string `envconfig:"SERVER_PORT" default:"8000"`
}
