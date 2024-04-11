package main

import(
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

type Beer struct {
	gorm.Model
	Name string
	Type string
	Detail string
	ImageURL string
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil{
		panic("failed to connect database")
	}
	db.AutoMigrate(&Beer{})
	fmt.Println("connected successful")

}