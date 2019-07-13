package database

import (
	"fmt"

	"github.com/rockdragon/daily-practice/code/golang/mod_proj/utils"
)

func CreateUserWithProfile(json utils.CreateUserSpec) (bool, string) {
	profile := Profile{
		Description: json.Description,
		Avatar:      json.Avatar,
	}
	user := User{
		Name:     json.Name,
		Password: json.Password,
	}
	user.Profile = profile

	conn := GetConnection()
	conn.Save(&user)

	errors := conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, fmt.Sprintf("%v", errors)
	}
	return true, ""
}

func RetrieveUserByName(name string) (bool, *User) {
	user := User{}
	profile := Profile{}
	conn := GetConnection()
	conn.Where("name = ?", name).First(&user)
	conn.Model(&user).Related(&profile)

	errors := conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, nil
	}
	if user.ID == 0 { // not found
		return false, nil
	}

	user.Profile = profile
	return true, &user
}

func UpdateUserWithProfile(json utils.UpdateUserSpec) (bool, string) {
	user := User{}
	conn := GetConnection()
	conn.Where("name = ?", json.Name).First(&user)

	errors := conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, fmt.Sprintf("%v", errors)
	}
	if user.ID == 0 { // not found
		return false, "User not exists"
	}

	user.Password = json.Password

	var profile Profile
	conn.Model(&user).Related(&profile)

	profile.Description = json.Description
	profile.Avatar = json.Avatar

	user.Profile = profile
	conn.Save(&user)

	errors = conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, fmt.Sprintf("%v", errors)
	}

	return true, ""
}

func DeleteUserWithProfile(name string) (bool, string) {
	user := User{}
	conn := GetConnection()
	conn.Where("name = ?", name).First(&user)

	errors := conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, fmt.Sprintf("%v", errors)
	}
	if user.ID == 0 { // not found
		return false, "User not exists"
	}

	user.Status = utils.Deleted
	conn.Save(&user)

	conn.Delete(&user)

	errors = conn.GetErrors()
	if errors != nil && len(errors) > 0 {
		return false, fmt.Sprintf("%v", errors)
	}

	return true, ""
}
