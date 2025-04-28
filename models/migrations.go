package models

import (
	"golang-temp/authentication"
	"golang-temp/database"
	"golang-temp/logging"
)

func MigrateTables() {
	err := database.DataBase.AutoMigrate(&Service{})

	if err != nil {
		panic("Faild to migrate Service")
	}

	err = database.DataBase.AutoMigrate(&Center{})

	if err != nil {
		panic("Faild to migrate Center")
	}

	err = database.DataBase.AutoMigrate(&logging.ApiLog{})

	if err != nil {
		panic("Faild to migrate ApiLog")
	}

	err = database.DataBase.AutoMigrate(&ServiceInstance{})

	if err != nil {
		panic("Faild to migrate ServiceInstance")
	}

	authentication.AuthenticationMigrations() // migrations of authentication package

}
