package connection

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// New new database connection
func NewDB() (*gorm.DB, error) {
	dns := "root:root@tcp(127.0.0.1:3306)/testdb"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
