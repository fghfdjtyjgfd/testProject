package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Beer struct {
	ID uint
	Name string
	Type string
	Detail string
	ImageURL string
}

func CreateBeer(db *gorm.DB, beer *Beer) {
	result := db.Create(beer)
	if result.Error != nil {
		log.Fatalf("Error to create beer: %v", result.Error)
	}
	fmt.Println("Created beer successful")
}

func GetBeer(db *gorm.DB, id uint) *Beer{
	var beer Beer
	result := db.First(&beer, id)
	if result.Error != nil {
		log.Fatalf("Error to get beer: %v", result.Error)
	}
	return &beer
}

func UpdateBeer (db *gorm.DB, beer *Beer) {
	
	result := db.Save(&beer)
	if result.Error != nil {
		log.Fatalf("Error to update beer: %v", result.Error)
	}
	fmt.Println("updated successful")
}

func DeleteBeer (db *gorm.DB, id uint){
	var beer Beer
	result := db.Delete(&beer ,id)
	if result.Error != nil {
		log.Fatalf("Error to delete beer: %v", result.Error)
	}
	fmt.Println("deleted successful")
}

func SearchBeer (db *gorm.DB, beerName string) *Beer{
	var beer Beer
	result := db.Where("name = ?", beerName).First(&beer)
	if result.Error != nil {
		log.Fatalf("Error to search beer: %v", result.Error)
	}
	return &beer
	
}