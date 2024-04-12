package main

import (
	conn "mariadb/connection"
	m "mariadb/model"
)

func main() {
	db, err := conn.NewDB()
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&m.Beer{})

}
