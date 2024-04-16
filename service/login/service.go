package login

import (
	"errors"
	"mariadb/utils"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type ServiceInterface interface {
	LoginUser(db *gorm.DB, form LoginForm) (string, error)
}

type Service struct {
	repository RepoInterface
}

func NewService() *Service {
	return &Service{
		repository: NewRepo(),
	}
}

func (s *Service) LoginUser(db *gorm.DB, form LoginForm) (string, error) {
	user, err := s.repository.FindUserOne(db, form.Email, form.ID)
	if err != nil {
		return "", err
	}

	isMatch := utils.ComparePassword(user.Password, form.Password)
	if !isMatch {
		return "", errors.New("Incorrect credential.")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte(os.Getenv("jwtSecretKey")))
	if err != nil {
		return "", err
	}

	return t, nil
}
