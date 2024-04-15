package model

type User struct {
	ID       int
	Email    string `gorm:"unique"`
	Password string
}
