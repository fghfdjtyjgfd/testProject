package main

import (
	conn "mariadb/connection"
	m "mariadb/model"
	"mariadb/router"
)
func main() {
	db, err := conn.NewDB()
	if err != nil {
		panic("failed to connect database")
	}


	router.NewApiRouter(db)
}
