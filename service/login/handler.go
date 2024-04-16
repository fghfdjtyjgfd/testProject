package login

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Interface interface {
	LoginUser(c *fiber.Ctx) error
}

type Endpoint struct {
	service  ServiceInterface
	database *gorm.DB
}

func NewEndpoint(db *gorm.DB) Interface {
	return &Endpoint{
		service:  NewService(),
		database: db,
	}
}

func (e *Endpoint) LoginUser(c *fiber.Ctx) error {
	form := LoginForm{}
	if err := c.BodyParser(&form); err != nil {
		return err
	}
	out, err := e.service.LoginUser(e.database, form)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// c.Cookie(&fiber.Cookie{
	// 	Name:     "jwt",
	// 	Value:    token,
	// 	Expires:  time.Now().Add(time.Hour * 72),
	// 	HTTPOnly: true,
	// })

	return c.JSON(out)
}
