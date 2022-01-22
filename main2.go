package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Mohtawa struct {
	ID   int    `json:"id"`
	Kind string `json:"kind"`
}

type Audience int

const (
	PassengerAudience Audience = 0
	DriverAudience    Audience = 1
)

func first(c *fiber.Ctx) error {
	sk := Mohtawa{}

	if err := c.BodyParser(&sk); err != nil {
		return err
	}
	if sk.ID == 0 {
		sk.Kind = "passenger"
	} else if sk.ID == 1 {
		sk.Kind = "driver"
	} else {
		fmt.Println("some err")
	}

	return c.JSON(sk)
}

func second(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "0" {
		fmt.Println("passenger")
	} else if id == "1" {
		fmt.Println("driver")
	}
	return nil
}

func Routers(app *fiber.App) {
	app.Post("/first", first)
	app.Post("/second/:id", second)
}

func main() {
	app := fiber.New()
	Routers(app)

	app.Listen(":3000")
}
