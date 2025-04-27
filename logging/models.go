package logging

import (
	"golang-temp/database"

	"gorm.io/gorm"
)

type ApiLog struct {
	gorm.Model
	Path         string
	ResponseTime float64
}

func addLog(path string, responseTime float64) {
	apiLog := ApiLog{Path: path, ResponseTime: responseTime}
	database.DataBase.Create(&apiLog)
}
