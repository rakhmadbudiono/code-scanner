package configs

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Name     string `envconfig:"DB_NAME" default:"code-scanner"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"postgres"`
	DSN      string
}
