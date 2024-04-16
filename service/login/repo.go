package login

import (
	m "mariadb/model"
	"mariadb/repository"

	"gorm.io/gorm"
)

type RepoInterface interface {
	FindUserOne(db *gorm.DB, email string, id int) (*m.User, error)
}

type Repo struct {
	repository.Repo
}

func NewRepo() *Repo {
	return &Repo{
		repository.Repo{},
	}
}

func (r *Repo) FindUserOne(db *gorm.DB, email string, id int) (*m.User, error) {
	model := m.User{}
	err := db.Where("email = ? or id = ?", email, id).First(&model)
	if err.Error != nil {
		return nil, err.Error
	}

	return &model, nil
}
