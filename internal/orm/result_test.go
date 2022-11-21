package orm_test

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/jackc/pgtype"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestGetAllResultsByRepositoryID(t *testing.T) {
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
				// TODO: mock where method to return orm.IDatabaseConnection
				mockDB.On("Where", mock.Anything, mock.Anything).Return(db).Once()
				mockDB.On("Find", mock.Anything).Return(db).Once()

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
				// TODO: mock where method to return orm.IDatabaseConnection
				mockDB.On("Where", mock.Anything, mock.Anything).Return(db).Once()
				mockDB.On("Find", mock.Anything).Return(db).Once()

				return mockDB
			}(),
			err: errors.New("database error"),
		},
	}

	for _, tc := range cases {
		_ = &orm.ORM{DB: tc.database}
		// TODO: test skipped for now due to incomplete mocking
		// _, err := o.GetAllResultsByRepositoryID(tc.input)

		// assert.Equal(t, tc.err, err)
	}
}

func TestCreateResult(t *testing.T) {
	cases := []struct {
		input    orm.Result
		database orm.IDatabaseConnection
		err      error
	}{
		{
			input: orm.Result{
				RepositoryID: "uuid",
				Status:       orm.Queued,
				Findings:     pgtype.JSONB{},
				QueuedAt:     time.Now(),
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
			input: orm.Result{
				RepositoryID: "uuid",
				Status:       orm.Queued,
				Findings:     pgtype.JSONB{},
				QueuedAt:     time.Now(),
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
		_, err := o.CreateResult(tc.input)

		assert.Equal(t, tc.err, err)
	}
}

func TestUpdateResult(t *testing.T) {
	cases := []struct {
		input    orm.Result
		database orm.IDatabaseConnection
		err      error
	}{
		{
			input: orm.Result{
				ID:           "uuid",
				RepositoryID: "uuid",
				Status:       orm.Success,
				Findings:     pgtype.JSONB{},
				QueuedAt:     time.Now(),
				ScanningAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				FinishedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
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
			input: orm.Result{
				ID:           "uuid",
				RepositoryID: "uuid",
				Status:       orm.Success,
				Findings:     pgtype.JSONB{},
				QueuedAt:     time.Now(),
				ScanningAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
				FinishedAt: sql.NullTime{
					Time:  time.Now(),
					Valid: true,
				},
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
		_, err := o.UpdateResult(tc.input)

		assert.Equal(t, tc.err, err)
	}
}
