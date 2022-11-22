package orm

import "gorm.io/gorm"

type IORM interface {
	GetAllRepositories() ([]Repository, error)
	CreateRepository(repo Repository) (*Repository, error)
	DeleteRepository(ID string) error
	GetRepositoryByID(ID string) (*Repository, error)
	UpdateRepository(repo Repository) (*Repository, error)
	GetAllResultsByRepositoryID(ID string) ([]Result, error)
	CreateResult(res Result) (*Result, error)
	UpdateResult(res Result) (*Result, error)
	GetResultByID(ID string) (*Result, error)
}

type ORM struct {
	DB IDatabaseConnection
}

type IDatabaseConnection interface {
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Create(value interface{}) *gorm.DB
	Delete(value interface{}, conds ...interface{}) *gorm.DB
	First(dest interface{}, conds ...interface{}) *gorm.DB
	Save(value interface{}) *gorm.DB
	Where(query interface{}, args ...interface{}) *gorm.DB
}
