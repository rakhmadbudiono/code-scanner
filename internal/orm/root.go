package orm

import "gorm.io/gorm"

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
