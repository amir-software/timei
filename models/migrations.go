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

	authentication.AuthenticationMigrations() // migrations of authentication package

}
