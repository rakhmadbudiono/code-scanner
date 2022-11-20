package controller

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rakhmadbudiono/code-scanner/configs"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
)

type IController interface {
	GetAllRepositories() ([]orm.Repository, error)
	CreateRepository(repo *orm.Repository) (*orm.Repository, error)
	DeleteRepository(ID string) error
	GetRepositoryByID(ID string) (*orm.Repository, error)
	UpdateRepository(repo *orm.Repository) (*orm.Repository, error)
	ScanRepository(ID string) error
	GetAllResultsByRepositoryID(ID string) ([]orm.Result, error)
}

type Controller struct {
	Config *configs.Config
	ORM    orm.IORM
}

type ControllerDependencyOption func(*Controller)

func NewController(cfg *configs.Config, opts ...ControllerDependencyOption) *Controller {
	ctrl := &Controller{
		Config: cfg,
	}

	for _, opt := range opts {
		opt(ctrl)
	}

	return ctrl
}

func WithDatabase() ControllerDependencyOption {
	return func(c *Controller) {
		db, err := gorm.Open(postgres.Open(c.Config.Database.DSN), &gorm.Config{})
		if err != nil {
			log.Fatalf("Couldn't connect to database: %s", err)
		}
		log.Println("Connected to database...")

		c.ORM = &orm.ORM{DB: db}
	}
}