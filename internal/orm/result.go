package orm

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type ResultStatus string

const (
	Queued     ResultStatus = "QUEUED"
	InProgress ResultStatus = "IN PROGRESS"
	Success    ResultStatus = "SUCCESS"
	Failure    ResultStatus = "FAILURE"
)

type Result struct {
	ID           string       `json:"id"`
	RepositoryID string       `json:"repository_id"`
	Status       ResultStatus `gorm:"type:enum_status;not null" json:"status"`
	Findings     pgtype.JSONB `gorm:"type:jsonb" json:"findings"`
	QueuedAt     time.Time    `json:"queued_at"`
	ScanningAt   sql.NullTime `json:"scanning_at"`
	FinishedAt   sql.NullTime `json:"finished_at"`
}

func (orm *ORM) GetAllResultsByRepositoryID(ID string) ([]Result, error) {
	var results []Result
	tx := orm.DB.Where("repository_id = ?", ID).Find(&results)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return results, nil
}

func (orm *ORM) CreateResult(res Result) (*Result, error) {
	res.ID = uuid.NewString()
	tx := orm.DB.Create(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &res, nil
}

func (orm *ORM) GetResultByID(ID string) (*Result, error) {
	result := &Result{ID: ID}
	tx := orm.DB.First(result)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return result, nil
}

func (orm *ORM) UpdateResult(res Result) (*Result, error) {
	tx := orm.DB.Save(&res)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &res, nil
}
