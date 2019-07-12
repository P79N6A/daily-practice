package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitMysql(connectionString string) (err error) {
	db, err := gorm.Open("mysql", connectionString)
	defer db.Close()
}
