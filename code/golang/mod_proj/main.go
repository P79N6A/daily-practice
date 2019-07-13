package main

import (
	"fmt"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"

	api "github.com/rockdragon/daily-practice/code/golang/mod_proj/api"
)

func main() {
	config := utils.GetConfig()
	fmt.Println(config)

	api.Serve(1323)
}
