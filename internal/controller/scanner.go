package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/jackc/pgtype"
	"github.com/rakhmadbudiono/code-scanner/internal/orm"
	"github.com/rakhmadbudiono/code-scanner/internal/scanner"
)

func (c *Controller) RunScanner() {
	if err := c.Sub.Subscribe(c.Config.Kafka.ScanRepoTopic, nil); err != nil {
		log.Panicf("error registering topic to subscribe: %s (%s)\n", err, c.Config.Kafka.ScanRepoTopic)
	}

	for {
		msg, err := c.Sub.ReadMessage(-1)
		if err != nil {
			log.Printf("consumer error: %s (%s)\n", err, msg)
			continue
		}

		ID := string(msg.Value)
		log.Printf("consuming message: %s", ID)
		go c.ProcessMessage(ID)
	}
}

func (c *Controller) ProcessMessage(ID string) {
	result, err := c.ORM.GetResultByID(ID)
	if err != nil {
		log.Printf("processing message, error get result by id: %s (%s)", err, ID)
		return
	}
	if result == nil {
		log.Printf("processing message, result not found: %s (%s)", err, ID)
		return
	}

	if err = c.markResultAsInProgress(result); err != nil {
		return
	}

	repo, err := c.ORM.GetRepositoryByID(result.RepositoryID)
	if err != nil {
		log.Printf("processing message, error get repository by id: %s (%s)", err, result.RepositoryID)
		c.markResultAsFailed(result)
		return
	}
	if repo == nil {
		log.Printf("processing message, repository not found: %s (%s)", err, result.RepositoryID)
		c.markResultAsFailed(result)
		return
	}

	sc := &scanner.ScanProcessor{FileSystem: memfs.New()}
	findings, err := sc.Process(repo.Link)
	if err != nil {
		log.Println(err)
		c.markResultAsFailed(result)
		return
	}

	convertedFindings := pgtype.JSONB{}
	if err = convertedFindings.Set(findings); err != nil {
		log.Printf("processing message, failed to assign findings to JSONB object: %s", err)
		return
	}
	result.Findings = pgtype.JSONB(convertedFindings)
	c.markResultAsSuccess(result)
}

func (c *Controller) markResultAsInProgress(result *orm.Result) error {
	result.Status = orm.InProgress
	result.ScanningAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if _, err := c.ORM.UpdateResult(*result); err != nil {
		log.Printf("processing message, error update result: %s (%s)", err, result.ID)
		c.markResultAsFailed(result)
		return err
	}
	return nil
}

func (c *Controller) markResultAsFailed(result *orm.Result) {
	result.Status = orm.Failure
	result.FinishedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if _, err := c.ORM.UpdateResult(*result); err != nil {
		log.Printf("processing message, error marking result as fail: %s (%s)", err, result.ID)
		return
	}
}

func (c *Controller) markResultAsSuccess(result *orm.Result) {
	result.Status = orm.Success
	result.FinishedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if _, err := c.ORM.UpdateResult(*result); err != nil {
		log.Printf("processing message, error marking result as success: %s (%s)", err, result.ID)
		return
	}
}
