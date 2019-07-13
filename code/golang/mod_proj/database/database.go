package database

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"
)

func GetConnection() *gorm.DB {
	config := utils.GetConfig().Mysql
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.User, config.Password, config.Address, config.Port, config.DB)

	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func SetConnectionPool(maxIdleConns int, maxOpenConns int) {
	log.Printf("SetMaxIdleConns: %d, SetMaxOpenConns: %d", maxIdleConns, maxOpenConns)
	conn := GetConnection()
	conn.DB().SetMaxIdleConns(maxIdleConns)
	conn.DB().SetMaxOpenConns(maxOpenConns)
	conn.DB().SetConnMaxLifetime(time.Hour)
}
