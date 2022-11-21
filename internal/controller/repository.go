package controller

import (
	"errors"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
)

func (c *Controller) GetAllRepositories() ([]orm.Repository, error) {
	return c.ORM.GetAllRepositories()
}

func (c *Controller) CreateRepository(repo *orm.Repository) (*orm.Repository, error) {
	return c.ORM.CreateRepository(*repo)
}

func (c *Controller) DeleteRepository(ID string) error {
	return c.ORM.DeleteRepository(ID)
}

func (c *Controller) GetRepositoryByID(ID string) (*orm.Repository, error) {
	return c.ORM.GetRepositoryByID(ID)
}

func (c *Controller) UpdateRepository(repo *orm.Repository) (*orm.Repository, error) {
	return c.ORM.UpdateRepository(*repo)
}

func (c *Controller) ScanRepository(ID string) error {
	repo, err := c.ORM.GetRepositoryByID(ID)
	if err != nil {
		return err
	}
	if repo == nil {
		return errors.New("repository not found")
	}

	result := orm.Result{
		RepositoryID: ID,
		Status:       orm.Queued,
		QueuedAt:     time.Now(),
	}
	if err = result.Findings.Set([]string{}); err != nil {
		return err
	}
	if _, err := c.ORM.CreateResult(result); err != nil {
		return err
	}

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &c.Config.Kafka.ScanRepoTopic, Partition: kafka.PartitionAny},
		Value:          []byte(ID),
	}
	if err := c.Pub.Produce(message, nil); err != nil {
		return err
	}

	return nil
}

func (c *Controller) GetAllResultsByRepositoryID(ID string) ([]orm.Result, error) {
	return c.ORM.GetAllResultsByRepositoryID(ID)
}
