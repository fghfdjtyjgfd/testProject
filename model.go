package main

import (
	"log"

	"gorm.io/gorm"
)

type Beer struct {
	ID       int
	Name     string
	Type     string
	Detail   string
	ImageURL string
}

func CreateBeer(db *gorm.DB, beer *Beer) error {
	result := db.Create(beer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetBeers(db *gorm.DB) []Beer {
	var beers []Beer
	result := db.Find(&beers)
	if result.Error != nil {
		log.Fatalf("Error to get beer: %v", result.Error)
	}
	return beers
}

func GetBeer(db *gorm.DB, id int) *Beer {
	var beer Beer
	result := db.First(&beer, id)
	if result.Error != nil {
		log.Fatalf("Error to get beer: %v", result.Error)
	}
	return &beer
}

func UpdateBeer(db *gorm.DB, beer *Beer) error {
	result := db.Save(&beer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteBeer(db *gorm.DB, id int) error {
	var beer Beer
	result := db.Delete(&beer, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func SearchBeer(db *gorm.DB, beerName string) *Beer {
	var beer Beer
	result := db.Where("name = ?", beerName).First(&beer)
	if result.Error != nil {
		log.Fatalf("Error to search beer: %v", result.Error)
	}
	return &beer

}
