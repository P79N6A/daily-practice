package database

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"
)

func GetConnection() *gorm.DB {
	config := utils.GetConfig().Mysql
	connStr := fmt.Sprintf("%s:%s@/%s:%d/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Address, config.Port, config.DB)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
