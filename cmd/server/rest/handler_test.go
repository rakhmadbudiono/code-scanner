package rest_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/rakhmadbudiono/code-scanner/cmd/server/rest"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllRepositories(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetAllRepositories").Return([]orm.Repository{}, nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetAllRepositories").Return(nil, errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.GetAllRepositories(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestCreateRepository(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("CreateRepository", mock.Anything).Return(&orm.Repository{}, nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPost,
			url:     "/repository",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("CreateRepository", mock.Anything).Return(nil, errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPost,
			url:     "/repository",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.CreateRepository(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestDeleteRepository(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("DeleteRepository", mock.Anything).Return(nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodDelete,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("DeleteRepository", mock.Anything).Return(errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodDelete,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.DeleteRepository(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestGetRepositoryByIDs(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetRepositoryByID", mock.Anything).Return(nil, errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.GetRepositoryByID(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestUpdateRepository(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("UpdateRepository", mock.Anything).Return(&orm.Repository{}, nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPut,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("UpdateRepository", mock.Anything).Return(nil, errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPut,
			url:     "/repository/:id",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.UpdateRepository(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestScanRepository(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("ScanRepository", mock.Anything).Return(nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPost,
			url:     "/repository/:id/scan",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("ScanRepository", mock.Anything).Return(errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodPost,
			url:     "/repository",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.ScanRepository(ctx)

		assert.Equal(t, tc.err, err)
	}
}

func TestGetAllResultsByRepositoryID(t *testing.T) {
	cases := []struct {
		server  *rest.Server
		method  string
		url     string
		reqBody io.Reader
		err     error
	}{
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetAllResultsByRepositoryID", mock.Anything).Return([]orm.Result{}, nil).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository/:id/result",
			reqBody: strings.NewReader(`{}`),
			err:     nil,
		},
		{
			server: func() *rest.Server {
				mockController := new(mocks.IController)
				mockController.On("GetAllResultsByRepositoryID", mock.Anything).Return(nil, errors.New("controller error")).Once()

				return &rest.Server{
					Controller: mockController,
				}
			}(),
			method:  http.MethodGet,
			url:     "/repository/:id/result",
			reqBody: strings.NewReader(`{}`),
			err:     errors.New("controller error"),
		},
	}

	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, tc.reqBody)
		req.Header.Set("Content-Type", "application/json")
		res := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, res)
		err := tc.server.GetAllResultsByRepositoryID(ctx)

		assert.Equal(t, tc.err, err)
	}
}
