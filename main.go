package main

import (
	u "mariadb/User"
	conn "mariadb/connection"
	m "mariadb/model"
	router "mariadb/router"
)

func main() {
	db, err := conn.NewDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&m.Beer{}, &u.User{})

	router.NewApiRouter(db)
}
