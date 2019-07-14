package main

import (
	"fmt"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/services"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"

	api "github.com/rockdragon/daily-practice/code/golang/mod_proj/api"
)

func main() {
	config := utils.GetConfig()
	fmt.Println(config)

	database.SetConnectionPool(config.Mysql.MaxIdleConns, config.Mysql.MaxOpenConns)

	go func() {
		services.Serve(1551)
	}()

	api.Serve(1323)
}
