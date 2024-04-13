package main

import (
	"fmt"

	conn "mariadb/connection"
	m "mariadb/model"
	"mariadb/router"
)

func main() {
	db, err := conn.NewDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&m.Beer{})
	
	router.NewApiRouter(db)
}
