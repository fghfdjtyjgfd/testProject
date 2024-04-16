package service

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	m "mariadb/model"
	repo "mariadb/repository"
)

type Databaser interface {
	FindUserOne(db *gorm.DB, email string, id int) (*m.User, error)
	CreateBeer(db *gorm.DB, beer *m.Beer) error
	GetBeers(db *gorm.DB) []m.Beer
	GetBeer(db *gorm.DB, id int) *m.Beer
	UpdateBeer(db *gorm.DB, beer *m.Beer) error
	DeleteBeer(db *gorm.DB, id int) error
	SearchBeer(db *gorm.DB, beerName string) *m.Beer
}

type Servicer interface {
	AuthRequired(c *fiber.Ctx) error
	CreateUser(db *gorm.DB, user *m.User) error
	LoginUser(db *gorm.DB, user *m.User) (string, error)
}

type Service struct {
	database Databaser
}

func NewService(database Databaser) *Service {
	return &Service{database: database}
}

func (s *Service) AuthRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("jwtSecretKey")

	token, err := jwt.ParseWithClaims(cookie, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}

func (s *Service) CreateUser(db *gorm.DB, user *m.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) LoginUser(db *gorm.DB, user *m.User) (string, error) {
	selectedUser, err := repo.FindUserOne(db, user.Email, user.ID)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("jwtSecretKey")))
	if err != nil {
		return "", err
	}
	return t, nil

}
