package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
)

var cfg *configs.Config = configs.New()

func TestNewController(t *testing.T) {
	cases := []struct {
		input    *configs.Config
		expected controller.Controller
	}{
		{
			input: cfg,
			expected: controller.Controller{
				Config: cfg,
			},
		},
	}

	for _, tc := range cases {
		ctrl := controller.NewController(tc.input)

		assert.Equal(t, tc.expected.Config.DSN, ctrl.Config.DSN)
	}
}

func TestWithDatabase(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cases := []struct {
		input struct {
			config *configs.Config
			orm    orm.IORM
		}
		expected controller.Controller
	}{
		{
			input: struct {
				config *configs.Config
				orm    orm.IORM
			}{
				config: cfg,
			},
			expected: controller.Controller{
				Config: cfg,
			},
		},
	}

	for _, tc := range cases {
		ctrl := controller.NewController(tc.input.config, controller.WithDatabase())

		assert.Equal(t, tc.expected.Config, ctrl.Config)
		assert.NotNil(t, ctrl.ORM)
	}
}
