package configs_test

import (
	"os"
	"testing"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cases := []struct {
		exportEnvs func()
		expected   configs.Config
	}{
		{
			exportEnvs: func() {
				os.Setenv("SERVER_PORT", "8000")
				os.Setenv("DB_HOST", "testhost")
				os.Setenv("DB_PORT", "1234")
				os.Setenv("DB_NAME", "dummydb")
				os.Setenv("DB_USER", "postgres")
				os.Setenv("DB_PASSWORD", "postgres")
			},
			expected: configs.Config{
				Server: configs.Server{
					Port: "8000",
				},
				Database: configs.Database{
					Host:     "testhost",
					Port:     "1234",
					Name:     "dummydb",
					User:     "postgres",
					Password: "postgres",
					DSN:      "host=testhost port=1234 user=postgres dbname=dummydb sslmode=disable password=postgres",
				},
			},
		},
	}

	for _, tc := range cases {
		tc.exportEnvs()
		cfg := configs.New()

		assert.Equal(t, tc.expected, *cfg)
	}
}
