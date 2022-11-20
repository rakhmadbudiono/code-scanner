package controller

import "github.com/rakhmadbudiono/code-scanner/internal/orm"

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
	return nil
}

func (c *Controller) GetAllResultsByRepositoryID(ID string) ([]orm.Result, error) {
	return []orm.Result{}, nil
}
