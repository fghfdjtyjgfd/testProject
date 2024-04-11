package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/testdb"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Beer{})

	app := fiber.New()

	app.Get("/beers", func(c *fiber.Ctx) error {
		return c.JSON(GetBeers(db))
	})

	app.Get("/beers/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		beer := GetBeer(db, id)
		return c.JSON(beer)
	})

	app.Post("/beers", func(c *fiber.Ctx) error {
		beer := new(Beer)
		if err := c.BodyParser(beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := CreateBeer(db, beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "created beer successful"})
	})

	app.Put("/beers/:id", func(c *fiber.Ctx) error {
		beer := new(Beer)
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := c.BodyParser(beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		beer.ID = id
		if err := UpdateBeer(db, beer); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "updated beer successful"})
	})

	app.Delete("/beers/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := DeleteBeer(db, id); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(fiber.Map{"message": "deleted beer successful"})
	})

	app.Listen(":8000")
}
