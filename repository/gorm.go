package repository

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"

	m "mariadb/model"
)

type GormDB struct {
	db *gorm.DB
}

func NewGorm(db *gorm.DB) *GormDB {
	return &GormDB{db: db}
}

func (g *GormDB) FindUserOne(db *gorm.DB, email string, id int) (*m.User, error) {
	model := m.User{}
	err := db.Where("email = ? or id = ?", email, id).First(&model)
	if err.Error != nil {
		return nil, err.Error
	}

	return &model, nil
}

func (g *GormDB) CreateBeer(db *gorm.DB, beer *m.Beer) error {
	for i := 0; i < 50; i++ {
		db.Create(&m.Beer{
			Name:     faker.Word(),
			Type:     faker.Word(),
			Detail:   faker.Paragraph(),
			ImageURL: fmt.Sprintf("http://test.com/%s", faker.UUIDDigit()),
		})
	}
	return nil
}

func (g *GormDB) UpdateBeer(db *gorm.DB, beer *m.Beer) error {
	result := db.Save(&beer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GormDB) DeleteBeer(db *gorm.DB, id int) error {
	var beer m.Beer
	result := db.Delete(&beer, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GormDB) SearchBeer(db *gorm.DB, beerName string) *m.Beer {
	var beer m.Beer
	result := db.Where("name = ?", beerName).First(&beer)
	if result.Error != nil {
		log.Fatalf("Error to search beer: %v", result.Error)
	}
	return &beer
}
