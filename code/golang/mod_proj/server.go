package main

import (
	api "github.com/rockdragon/daily-practice/code/golang/mod_proj/api"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
