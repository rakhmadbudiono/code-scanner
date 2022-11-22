package controller_test

import (
	"errors"
	"testing"

	"github.com/rakhmadbudiono/code-scanner/internal/controller"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRunScanner(t *testing.T) {
	cases := []struct {
		controller *controller.Controller
	}{
		{
			controller: func() *controller.Controller {
				mockSub := new(mocks.ISubscriber)
				mockSub.On("Subscribe", mock.Anything, mock.Anything).Return(errors.New("error subscribe"))

				return &controller.Controller{
					Sub: mockSub,
				}
			}(),
		},
	}

	for _, tc := range cases {
		assert.Panics(t, tc.controller.RunScanner)
	}
}

func TestProcessMessage(t *testing.T) {
	cases := []struct {
		input      string
		controller *controller.Controller
	}{
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, errors.New(("error fetching result"))).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(nil, nil).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, nil).Once()
				mockORM.On("UpdateResult", mock.Anything).Return(&orm.Result{}, errors.New("error update result status"))

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, nil).Once()
				mockORM.On("UpdateResult", mock.Anything).Return(&orm.Result{}, nil)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{}, errors.New("error fetching repo")).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, nil).Once()
				mockORM.On("UpdateResult", mock.Anything).Return(&orm.Result{}, nil)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(nil, nil).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, nil).Once()
				mockORM.On("UpdateResult", mock.Anything).Return(&orm.Result{}, nil)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{Link: ""}, nil).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
		{
			input: "uuid",
			controller: func() *controller.Controller {
				mockORM := new(mocks.IORM)
				mockORM.On("GetResultByID", mock.Anything).Return(&orm.Result{}, nil).Once()
				mockORM.On("UpdateResult", mock.Anything).Return(&orm.Result{}, nil)
				mockORM.On("GetRepositoryByID", mock.Anything).Return(&orm.Repository{Link: "https://github.com/rakhmadbudiono/duck-pic-service"}, nil).Once()

				return &controller.Controller{
					ORM: mockORM,
				}
			}(),
		},
	}

	for _, tc := range cases {
		tc.controller.ProcessMessage(tc.input)
	}
}
