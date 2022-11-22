package controller

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/confluentinc/confluent-kafka-go/kafka"

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
	RunScanner()
}

type Controller struct {
	Config *configs.Config
	ORM    orm.IORM
	Pub    IPublisher
	Sub    ISubscriber
}

type IPublisher interface {
	Produce(msg *kafka.Message, deliveryChan chan kafka.Event) error
}

type ISubscriber interface {
	Subscribe(topic string, rebalanceCb kafka.RebalanceCb) error
	ReadMessage(timeout time.Duration) (*kafka.Message, error)
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

func WithPublisher() ControllerDependencyOption {
	return func(c *Controller) {
		cfg := &kafka.ConfigMap{
			"bootstrap.servers": c.Config.Kafka.Servers,
		}
		pub, err := kafka.NewProducer(cfg)
		if err != nil {
			log.Fatalf("Couldn't connect to kafka: %s", err)
		}
		log.Println("Connected to kafka...")

		c.Pub = pub
	}
}

func WithSubscriber() ControllerDependencyOption {
	return func(c *Controller) {
		cfg := &kafka.ConfigMap{
			"bootstrap.servers": c.Config.Kafka.Servers,
			"group.id":          "myGroup",
		}
		sub, err := kafka.NewConsumer(cfg)
		if err != nil {
			log.Fatalf("Couldn't connect to kafka: %s", err)
		}
		log.Println("Connected to kafka...")

		c.Sub = sub
	}
}
