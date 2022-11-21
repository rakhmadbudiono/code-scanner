package configs

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

var root Config

// Config wraps all config types
type Config struct {
	Server
	Database
	Kafka
}

// New load and return config object
func New() *Config {
	if err := envconfig.Process("SERVER", &root.Server); err != nil {
		log.Fatal(err.Error())
	}
	if err := envconfig.Process("DB", &root.Database); err != nil {
		log.Fatal(err.Error())
	}
	root.Database.DSN = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		root.Database.Host,
		root.Database.Port,
		root.Database.User,
		root.Database.Name,
		root.Database.Password,
	)
	if err := envconfig.Process("KAFKA", &root.Kafka); err != nil {
		log.Fatal(err.Error())
	}

	return &root
}
