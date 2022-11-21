package configs_test

import (
	"os"
	"strings"
	"testing"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	envs := os.Environ()
	cases := []struct {
		exportEnvs func()
		expected   configs.Config
	}{
		{

			exportEnvs: func() {
				os.Clearenv()
			},
			expected: configs.Config{
				Server: configs.Server{
					Port: "8000",
				},
				Database: configs.Database{
					Host:     "localhost",
					Port:     "5432",
					Name:     "code-scanner",
					User:     "postgres",
					Password: "postgres",
					DSN:      "host=localhost port=5432 user=postgres dbname=code-scanner sslmode=disable password=postgres",
				},
				Kafka: configs.Kafka{
					Servers:       "localhost",
					ScanRepoTopic: "code-scanner.repository.scan",
				},
			},
		},
	}

	for _, tc := range cases {
		tc.exportEnvs()
		cfg := configs.New()

		assert.Equal(t, tc.expected, *cfg)
	}

	// return envs
	for _, env := range envs {
		pair := strings.SplitN(env, "=", 2)
		os.Setenv(pair[0], pair[1])
	}
}
