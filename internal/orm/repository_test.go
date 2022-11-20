package orm_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/mocks"
)

func TestGetAllRepositories(t *testing.T) {
	cases := []struct {
		database orm.IDatabaseConnection
		err      error
	}{
		{
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: nil,
				}
				mockDB.On("Find", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: nil,
		},
		{
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: errors.New("database error"),
				}
				mockDB.On("Find", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		o := &orm.ORM{DB: tc.database}
		_, err := o.GetAllRepositories()

		assert.Equal(t, tc.err, err)
	}
}

func TestCreateRepository(t *testing.T) {
	cases := []struct {
		input    orm.Repository
		database orm.IDatabaseConnection
		expected *orm.Repository
		err      error
	}{
		{
			input: orm.Repository{
				Name: "test-repo",
				Link: "github.com/rakhmadbudinono/test-repo",
			},
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: nil,
				}
				mockDB.On("Create", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: nil,
		},
		{
			input: orm.Repository{
				Name: "test-repo",
				Link: "github.com/rakhmadbudinono/test-repo",
			},
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: errors.New("database error"),
				}
				mockDB.On("Create", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		o := &orm.ORM{DB: tc.database}
		_, err := o.CreateRepository(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
func TestDeleteRepository(t *testing.T) {
	cases := []struct {
		input    string
		database orm.IDatabaseConnection
		err      error
	}{
		{
			input: "uuid",
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: nil,
				}
				mockDB.On("Delete", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: nil,
		},
		{
			input: "uuid",
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: errors.New("database error"),
				}
				mockDB.On("Delete", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		o := &orm.ORM{DB: tc.database}
		err := o.DeleteRepository(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
func TestGetRepositoryByID(t *testing.T) {
	cases := []struct {
		input    string
		database orm.IDatabaseConnection
		err      error
	}{
		{
			input: "uuid",
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: nil,
				}
				mockDB.On("First", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: nil,
		},
		{
			input: "uuid",
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: errors.New("database error"),
				}
				mockDB.On("First", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		o := &orm.ORM{DB: tc.database}
		_, err := o.GetRepositoryByID(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
func TestUpdateRepository(t *testing.T) {
	cases := []struct {
		input    orm.Repository
		database orm.IDatabaseConnection
		err      error
	}{
		{
			input: orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudinono/test-repo",
			},
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: nil,
				}
				mockDB.On("Save", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: nil,
		},
		{
			input: orm.Repository{
				ID:   "uuid",
				Name: "test-repo",
				Link: "github.com/rakhmadbudinono/test-repo",
			},
			database: func() orm.IDatabaseConnection {
				mockDB := new(mocks.IDatabaseConnection)
				db := &gorm.DB{
					Error: errors.New("database error"),
				}
				mockDB.On("Save", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		o := &orm.ORM{DB: tc.database}
		_, err := o.UpdateRepository(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
