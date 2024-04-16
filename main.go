package main

import (
	conn "mariadb/connection"
	"mariadb/model"
	m "mariadb/model"
	router "mariadb/router"
)

func main() {
	db, err := conn.NewDB()
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&m.Beer{}, &model.User{})

	router.NewApiRouter(db)
}
