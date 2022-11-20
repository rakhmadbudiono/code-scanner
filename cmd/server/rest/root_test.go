package rest_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rakhmadbudiono/code-scanner/cmd/server/rest"
	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/controller"
	"github.com/rakhmadbudiono/code-scanner/mocks"
)

var cfg *configs.Config = configs.New()

func TestStart(t *testing.T) {
	cases := []struct {
		server *rest.Server
		err    error
	}{
		{
			server: func() *rest.Server {
				mockEcho := new(mocks.IEcho)
				mockEcho.On("Start", mock.Anything).Return(nil)

				return &rest.Server{
					Echo:       mockEcho,
					Controller: controller.NewController(cfg),
					Config:     cfg,
				}
			}(),
			err: nil,
		},
		{
			server: func() *rest.Server {
				mockEcho := new(mocks.IEcho)
				mockEcho.On("Start", mock.Anything).Return(errors.New("couldn't start server"))

				return &rest.Server{
					Echo:       mockEcho,
					Controller: controller.NewController(cfg),
					Config:     cfg,
				}
			}(),
			err: errors.New("couldn't start server"),
		},
	}

	for _, tc := range cases {
		if tc.err != nil {
			assert.Panics(t, tc.server.Start)
		} else {
			assert.NotPanics(t, tc.server.Start)
		}
	}
}
