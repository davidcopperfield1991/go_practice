package main

import (
	"fmt"
	"io/ioutil"

	"github.com/Jeffail/gabs"
	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()
	data, _ := ioutil.ReadFile("data.json")
	fmt.Println(data)

	jsonParsed, err := gabs.ParseJSON([]byte(data))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(jsonParsed)
	})

	app.Listen(":4001")

}
