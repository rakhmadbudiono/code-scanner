package orm

import (
	"github.com/google/uuid"
)

type Repository struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `gorm:"type:varchar(100);not null" json:"name"`
	Link string `gorm:"type:varchar(150);not null" json:"link"`
}

func (orm *ORM) GetAllRepositories() ([]Repository, error) {
	var repos []Repository
	tx := orm.DB.Find(&repos)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return repos, nil
}

func (orm *ORM) CreateRepository(repo Repository) (*Repository, error) {
	repo.ID = uuid.NewString()
	tx := orm.DB.Create(&repo)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &repo, nil
}

func (orm *ORM) DeleteRepository(ID string) error {
	repo := &Repository{ID: ID}
	tx := orm.DB.Delete(repo)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (orm *ORM) GetRepositoryByID(ID string) (*Repository, error) {
	repo := &Repository{ID: ID}
	tx := orm.DB.First(repo)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return repo, nil
}

func (orm *ORM) UpdateRepository(repo Repository) (*Repository, error) {
	tx := orm.DB.Save(&repo)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &repo, nil
}
