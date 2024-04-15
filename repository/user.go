package repository

import (
	"gorm.io/gorm"

	m "mariadb/model"
)

func FindUserOne(db *gorm.DB, email string, id int) (*m.User, error) {
	model := m.User{}
	err := db.Where("email = ? or id = ?", email, id).First(&model)
	if err.Error != nil {
		return nil, err.Error
	}

	return &model, nil
}
