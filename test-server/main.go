package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gofiber/fiber"
)

func main() {

	app := fiber.New()
	data, _ := ioutil.ReadFile("data.json")
	fmt.Println(data)
	jj := map[string]string{}
	err := json.Unmarshal(data, &jj)
	fmt.Println(err)
	fmt.Println(jj)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(jj)
	})

	app.Listen(":4001")

}
