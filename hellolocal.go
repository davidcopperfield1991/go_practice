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
	db *gorm.DB
}

func main() {

	// var i Info

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Info{})
	db = db.FirstOrCreate(&Info{Visit: 1, ID: 1})

	var v Visitsay
	v = Visitsay{
		db: db,
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)
	e.GET("/visit", v.visitshow)

	// Start server
	e.Logger.Fatal(e.Start(":9020"))

}

// Handle
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hello world")

}

func (v Visitsay) visitshow(c echo.Context) error {
	var i Info
	db := v.db
	row := db.Where("id = ?", 1).First(&i)
	row.Updates(Info{Visit: i.Visit + 1})

	return c.JSON(http.StatusOK, i.Visit)

}
