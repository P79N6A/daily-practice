package main

import (
	"fmt"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/database"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"

	api "github.com/rockdragon/daily-practice/code/golang/mod_proj/api"
)

func main() {
	config := utils.GetConfig()
	fmt.Println(config)

	database.SetConnectionPool(config.Mysql.MaxIdleConns, config.Mysql.MaxOpenConns)

	api.Serve(1323)
}
