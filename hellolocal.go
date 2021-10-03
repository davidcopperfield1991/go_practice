package main

import (
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

type Visitsay struct {
	Count int
}

func main() {

	var i Info
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.GET("/visit", i.visitshow)

	// Start server
	e.Logger.Fatal(e.Start(":9020"))

}

// Handle
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")

}

func (i Info) visitshow(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Info{})
	db.FirstOrCreate(&Info{Visit: 1, ID: 1})

	row := db.Where("id = ?", 1).First(&i)
	row.Updates(Info{Visit: i.Visit + 1})

	return c.JSON(http.StatusOK, i.Visit)

}
