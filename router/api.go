package router

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	u "mariadb/User"
	mid "mariadb/middleWare"
	m "mariadb/model"
	rp "mariadb/repository"
)

func NewApiRouter(db *gorm.DB) {
	app := fiber.New()
	app.Get("/beers/all", func(c *fiber.Ctx) error {
		return c.JSON(rp.GetBeers(db))
	})

	app.Use("/beers", mid.AuthRequired)

	app.Get("/beers", func(c *fiber.Ctx) error {
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
	})

	app.Post("/beers", func(c *fiber.Ctx) error {
		beer := new(m.Beer)
		if err := c.BodyParser(beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := rp.CreateBeer(db, beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "created beer successful"})
	})

	app.Put("/beers/:id", func(c *fiber.Ctx) error {
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
	})

	app.Delete("/beers/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := rp.DeleteBeer(db, id); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "deleted beer successful"})
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		user := new(u.User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err := u.CreateUser(db, user)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{"message": "registed user successful"})
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(u.User)
		if err := c.BodyParser(user); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		token, err := u.LoginUser(db, user)
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
	})

	app.Listen(":8000")
}
