package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model

	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (User, error) {
	var (
		user User
		err  error
	)
	fmt.Println(username, password)
	err = db.Where(User{Username: username}).First(&user).Error

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	return user, err
}

func CheckUserExists(username string) bool {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func CreateUser(data map[string]interface{}) (User, error) {
	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}

	err := db.Model(&User{}).Create(&user).Error

	return user, err
}
