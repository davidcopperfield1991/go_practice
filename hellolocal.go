package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"encoding/json"
	"io/ioutil"
	"log"
)

type Info struct {
	// gorm.Model
	ID    uint `gorm:"primary_key"`
	Visit uint
}

type Visitsay struct {
	db *gorm.DB
}

type Person struct {
	Name string `json:"name"`
}

func hellosay(c echo.Context) error {
	person := Person{}
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed reading the request body : %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(b, &person)
	if err != nil {
		log.Printf("failed unmarshaling ... : %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("hello %#v ", person)
	return c.String(http.StatusOK, "we got your name")
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
	e.POST("/hellosay", hellosay)

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
