package main

import (
	"log"
	"regexp"

	"github.com/asaskevich/govalidator"
)

type person struct {
	ID     int
	Email  string `valid:"email"`
	Mobile string `valid:"mobile"`
}

func init() {
	pattern := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	regex := regexp.MustCompile(pattern)

	govalidator.CustomTypeTagMap.Set("mobile",
		govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
			switch v := context.(type) {
			case person:
				return regex.MatchString(v.Mobile)
			}
			return false
		}))
}

func main() {
	p := person{2, "hello.com", "1234533423"}
	if _, err := govalidator.ValidateStruct(p); err != nil {
		log.Fatal(err)
	}
}
