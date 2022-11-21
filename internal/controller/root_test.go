package controller_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
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
		input *configs.Config
	}{
		{
			input: cfg,
		},
	}

	for _, tc := range cases {
		ctrl := controller.NewController(tc.input, controller.WithDatabase())

		assert.NotNil(t, ctrl.ORM)
	}
}

func TestWithPublisher(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	cases := []struct {
		input *configs.Config
	}{
		{
			input: cfg,
		},
	}

	for _, tc := range cases {
		ctrl := controller.NewController(tc.input, controller.WithPublisher())

		assert.NotNil(t, ctrl.Pub)
	}
}
