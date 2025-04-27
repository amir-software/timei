package authentication

import (
	"fmt"
	"golang-temp/database"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (user *User) CreateUser() int {

	hash_password, err := HashPassword(user.Password)
	if err == nil {
		user.Password = hash_password
		database.DataBase.Create(&user)
		return user.ID
	} else {
		fmt.Print(err)
		return 0
	}
}

func (user *User) ValidateCredentials() bool {
	var users []User
	email := user.Email

	database.DataBase.Where("email = ?", email).Find(&users)

	return CheckPassword(user.Password, users[0].Password)
}
