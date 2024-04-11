package main

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	

)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	db.AutoMigrate(&Beer{})
	fmt.Println("connected successful")

	// newBeer := &Beer{
	// 	ID: 1,
	// 	Name: "cbg",
	// 	Type: "good",
	// 	Detail: "test",
	// 	ImageURL: "google",
	// }
	// CreateBeer(db, newBeer)
	currentBeer := SearchBeer(db, "cbg")
	fmt.Println(currentBeer)
}