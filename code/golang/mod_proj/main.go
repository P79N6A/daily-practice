package main

import (
	"fmt"

	utils "github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"

	"github.com/labstack/echo"
	api "github.com/rockdragon/daily-practice/code/golang/mod_proj/api"
)

func main() {
	config := utils.GetConfig()
	fmt.Println(config)

	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
