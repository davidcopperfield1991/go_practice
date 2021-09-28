package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Info struct {
	// gorm.Model
	ID    uint `gorm:"primary_key"`
	Visit uint
}

func data() int {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Info{})

	db.Create(&Info{Visit: 1, ID: 1})
	var i Info
	row := db.Where("id = ?", 1).First(&i)

	row.Updates(Info{Visit: i.Visit + 1})
	dd := i.Visit

	return int(dd)

}

func main() {

	// Echo instance
	e := echo.New()

	fmt.Println(data())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.GET("/visit", visit)

	// Start server
	e.Logger.Fatal(e.Start(":9020"))

}

// Handle
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")

}

func visit(c echo.Context) error {
	dd := data()
	// ds := string(dd)

	return c.JSON(http.StatusOK, dd)

}

