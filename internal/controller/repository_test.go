package controller_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/rakhmadbudiono/code-scanner/internal/controller"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/mocks"
)

func TestGetAllRepositories(t *testing.T) {
	cases := []struct {
		controller controller.Controller
		expected   []orm.Repository
		err        error
	}{
		{
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetAllRepositories").Return([]orm.Repository{
					orm.Repository{
						ID:   "uuid",
						Name: "test-repo",
						Link: "github.com/rakhmadbudiono/test-repo",
					},
				}, nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			expected: []orm.Repository{
				orm.Repository{
					ID:   "uuid",
					Name: "test-repo",
					Link: "github.com/rakhmadbudiono/test-repo",
				},
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		repos, err := tc.controller.GetAllRepositories()

		assert.Equal(t, tc.expected, repos)
		assert.Equal(t, tc.err, err)
	}
}
func TestCreateRepository(t *testing.T) {
	cases := []struct {
		input      *orm.Repository
		controller controller.Controller
		expected   *orm.Repository
		err        error
	}{
		{
			input: &orm.Repository{
				Name: "test-repo",
				Link: "github.com/rakhmadbudiono/test-repo",
			},
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("CreateRepository", mock.Anything).Return(&orm.Repository{
					ID:   "uuid",
					Name: "test-repo",
					Link: "github.com/rakhmadbudiono/test-repo",
				}, nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			expected: &orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudiono/test-repo",
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		repo, err := tc.controller.CreateRepository(tc.input)

		assert.Equal(t, tc.expected, repo)
		assert.Equal(t, tc.err, err)
	}
}
func TestDeleteRepository(t *testing.T) {
	cases := []struct {
		input      string
		controller controller.Controller
		err        error
	}{
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("DeleteRepository", mock.Anything).Return(nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			err: nil,
		},
	}

	for _, tc := range cases {
		err := tc.controller.DeleteRepository(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
func TestGetRepositoryByID(t *testing.T) {
	cases := []struct {
		input      string
		controller controller.Controller
		expected   *orm.Repository
		err        error
	}{
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{
					ID:   "uuid",
					Name: "test-repo",
					Link: "github.com/rakhmadbudiono/test-repo",
				}, nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			expected: &orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudiono/test-repo",
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		repo, err := tc.controller.GetRepositoryByID(tc.input)

		assert.Equal(t, tc.expected, repo)
		assert.Equal(t, tc.err, err)
	}
}
func TestUpdateRepository(t *testing.T) {
	cases := []struct {
		input      *orm.Repository
		controller controller.Controller
		expected   *orm.Repository
		err        error
	}{
		{
			input: &orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudiono/test-repo",
			},
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("UpdateRepository", mock.Anything).Return(&orm.Repository{
					ID:   "uuid",
					Name: "test-repo",
					Link: "github.com/rakhmadbudiono/test-repo",
				}, nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			expected: &orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudiono/test-repo",
			},
			err: nil,
		},
	}

	for _, tc := range cases {
		repo, err := tc.controller.UpdateRepository(tc.input)

		assert.Equal(t, tc.expected, repo)
		assert.Equal(t, tc.err, err)
	}
}

func TestScanRepository(t *testing.T) {
	cases := []struct {
		input      string
		controller controller.Controller
		err        error
	}{
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, nil).Once()
				mockORM.On("CreateResult", mock.Anything).Return(&orm.Result{}, nil).Once()

				mockPub := new(mocks.IPublisher)
				mockPub.On("Produce", mock.Anything, mock.Anything).Return(nil).Once()

				return controller.Controller{
					ORM:    mockORM,
					Pub:    mockPub,
					Config: cfg,
				}
			}(),
			err: nil,
		},
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, errors.New("error when get repo")).Once()
				mockORM.On("CreateResult", mock.Anything).Return(&orm.Result{}, nil).Once()

				mockPub := new(mocks.IPublisher)
				mockPub.On("Produce", mock.Anything, mock.Anything).Return(nil).Once()

				return controller.Controller{
					ORM:    mockORM,
					Pub:    mockPub,
					Config: cfg,
				}
			}(),
			err: errors.New("error when get repo"),
		},
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(nil, nil).Once()
				mockORM.On("CreateResult", mock.Anything).Return(&orm.Result{}, nil).Once()

				mockPub := new(mocks.IPublisher)
				mockPub.On("Produce", mock.Anything, mock.Anything).Return(nil).Once()

				return controller.Controller{
					ORM:    mockORM,
					Pub:    mockPub,
					Config: cfg,
				}
			}(),
			err: errors.New("repository not found"),
		},
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, nil).Once()
				mockORM.On("CreateResult", mock.Anything).Return(&orm.Result{}, errors.New("error when create result")).Once()

				mockPub := new(mocks.IPublisher)
				mockPub.On("Produce", mock.Anything, mock.Anything).Return(nil).Once()

				return controller.Controller{
					ORM:    mockORM,
					Pub:    mockPub,
					Config: cfg,
				}
			}(),
			err: errors.New("error when create result"),
		},
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, nil).Once()
				mockORM.On("CreateResult", mock.Anything).Return(&orm.Result{}, nil).Once()

				mockPub := new(mocks.IPublisher)
				mockPub.On("Produce", mock.Anything, mock.Anything).Return(errors.New("error when produce event")).Once()

				return controller.Controller{
					ORM:    mockORM,
					Pub:    mockPub,
					Config: cfg,
				}
			}(),
			err: errors.New("error when produce event"),
		},
	}

	for _, tc := range cases {
		err := tc.controller.ScanRepository(tc.input)

		assert.Equal(t, tc.err, err)
	}
}

func TestGetAllResultsByRepositoryID(t *testing.T) {
	cases := []struct {
		input      string
		controller controller.Controller
		expected   []orm.Result
		err        error
	}{
		{
			input: "uuid",
			controller: func() controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetAllResultsByRepositoryID", mock.Anything).Return([]orm.Result{}, nil).Once()

				return controller.Controller{
					ORM: mockORM,
				}
			}(),
			expected: []orm.Result{},
			err:      nil,
		},
	}

	for _, tc := range cases {
		results, err := tc.controller.GetAllResultsByRepositoryID(tc.input)

		assert.Equal(t, tc.expected, results)
		assert.Equal(t, tc.err, err)
	}
}
