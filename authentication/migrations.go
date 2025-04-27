package authentication

import "golang-temp/database"

func AuthenticationMigrations() {
	err := database.DataBase.AutoMigrate(&User{})
	if err != nil {
		panic("Faild to migrate User model")
	}

}
