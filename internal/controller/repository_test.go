package controller_test

import (
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
			input:      "uuid",
			controller: controller.Controller{},
			err:        nil,
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
			input:      "uuid",
			controller: controller.Controller{},
			expected:   []orm.Result{},
			err:        nil,
		},
	}

	for _, tc := range cases {
		results, err := tc.controller.GetAllResultsByRepositoryID(tc.input)

		assert.Equal(t, tc.expected, results)
		assert.Equal(t, tc.err, err)
	}
}
