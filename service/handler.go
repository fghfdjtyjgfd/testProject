package service

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	u "mariadb/User"
	m "mariadb/model"
	rp "mariadb/repository"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register (c *fiber.Ctx, db *gorm.DB) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := u.CreateUser(db, user)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{"message": "registed user successful"})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	user := new(m.User)
	if err := c.BodyParser(user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	token, err := u.LoginUser(&gorm.DB{}, user)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 72),
		HTTPOnly: true,
	})

	return c.JSON(fiber.Map{"message": "login successful"})
}

func (h *Handler) GetBeers(c *fiber.Ctx, db *gorm.DB) error {
	var beers []m.Beer

	sql := "SELECT * FROM testdb.beers"

	if name := c.Query("name"); name != "" {
		sql = fmt.Sprintf("%s WHERE Name LIKE '%%%s%%' ", sql, name)
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	perPage := 10
	var total int64

	db.Raw(sql).Count(&total)

	sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	db.Raw(sql).Scan(&beers)

	return c.JSON(fiber.Map{
		"data":     beers,
		"total":    total,
		"page":     page,
		"lastPage": math.Ceil(float64(total / int64(perPage))),
	})
}

func (h *Handler) PostBeer(c *fiber.Ctx, db *gorm.DB) error {
	beer := new(m.Beer)
	if err := c.BodyParser(beer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := rp.CreateBeer(db, beer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "created beer successful"})
}

func (h *Handler) PutBeer(c *fiber.Ctx, db *gorm.DB) error {
	beer := new(m.Beer)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := c.BodyParser(beer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	beer.ID = id
	if err := rp.UpdateBeer(db, beer); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "updated beer successful"})
}

func (h *Handler) DeleteBeer (c *fiber.Ctx, db *gorm.DB) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	if err := rp.DeleteBeer(db, id); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(fiber.Map{"message": "deleted beer successful"})
}