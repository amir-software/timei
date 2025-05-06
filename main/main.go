package main

import (
	"fmt"
	"golang-temp/apis"
	"golang-temp/database"
	"golang-temp/models"
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	database.InitDataBase("timei.db")
	models.MigrateTables()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	serverPort := os.Getenv("SERVER_PORT")
	server := gin.Default()

	apis.RegisterRoutes(server)

	sugar.Info("Server is going to run")

	// apis.SendMain()
	server.Run(fmt.Sprintf(":%s", serverPort))

}
