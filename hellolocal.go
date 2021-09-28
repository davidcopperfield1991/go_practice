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

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Info{})
	//   db.Create(&Info{visit_count: ++1})
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/hello", hello)

	fmt.Println("inja")
	// db.Create(&Info{Visit: 1})
	// var info Info
	// db.First(&info, 1)
	row := db.Model(Info{}).Select("Visit").Where("id = ?", 1)
	jow := db.First(row)
	fmt.Println(jow)
	for i := 0; i < 2; i++ {
		fmt.Println(row)
	}
	// var num2 = db.First(&info)
	// fmt.Println(num2.First())
	//  visit_calculator := 4
	// a := &Info{
	// 	ID:    1,
	// 	Visit: uint(func visit_calculator(i )
	//
	// db.Save(a)

	// var num = db.First(info, 1)
	// yam := &num
	// fmt.Println(num)
	// db.Model(&info).Updates(Info{Visit: uint(visit_calculator(1))})
	// result := Info{}.visit_count
	// fmt.Println(result)
	// fmt.Println(Info{})

	// Start server
	e.Logger.Fatal(e.Start(":9020"))
}

// Handle
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "hi")

}

func visit_calculator(i int) int {
	bebin := i + 1
	return bebin

}
