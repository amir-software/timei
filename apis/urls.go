package apis

import (
	"golang-temp/authentication"
	"golang-temp/logging"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	router := server.Group("/")
	router.Use(logging.Logging)

	mustAuthenticate := server.Group("/")
	// mustAuthenticate.Use(authentication.Authenticate)
	mustAuthenticate.GET("/get_services", GetServices)
	mustAuthenticate.GET("/get_service", GetServiceByIdService)
	mustAuthenticate.POST("/create_service", CreateService)
	mustAuthenticate.PUT("/update_service", UpdateServices)
	mustAuthenticate.DELETE("/delete_service", UpdateServices)

	//
	router.POST("/sign_up", authentication.SignUp)
	router.POST("/login", authentication.Login)
}
