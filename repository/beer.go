package repository

import (
	"fmt"
	"log"

	"github.com/bxcodec/faker/v3"
	"gorm.io/gorm"

	m "mariadb/model"
)

func CreateBeer(db *gorm.DB, beer *m.Beer) error {
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


func UpdateBeer(db *gorm.DB, beer *m.Beer) error {
	result := db.Save(&beer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteBeer(db *gorm.DB, id int) error {
	var beer m.Beer
	result := db.Delete(&beer, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SearchBeer(db *gorm.DB, beerName string) *m.Beer {
	var beer m.Beer
	result := db.Where("name = ?", beerName).First(&beer)
	if result.Error != nil {
		log.Fatalf("Error to search beer: %v", result.Error)
	}
	return &beer

}
