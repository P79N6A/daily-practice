package database

import (
	"github.com/jinzhu/gorm"
	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"
)

type User struct {
	gorm.Model
	Name string
	Password string
	Status utils.UserStatus
}

type Product struct {
	gorm.Model
	Name string
	Price uint32
	Description string
	PicUrl string
}