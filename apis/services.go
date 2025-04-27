package apis

import (
	"fmt"
	"golang-temp/models"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetServices(context *gin.Context) {
	params := context.Request.URL.Query()

	var page, pageSize int

	if params.Get("page") == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(params.Get("page"))
	}
	if params.Get("page_size") == "" {
		pageSize = 3
	} else {
		pageSize, _ = strconv.Atoi(params.Get("page_size"))
	}

	cleanParams := getCleanParams(params)

	services := models.GetServices(cleanParams)

	pagniatedService, _ := paginate(page, pageSize, services)

	context.JSON(http.StatusOK, gin.H{"services": pagniatedService})
}

func GetServiceByIdService(context *gin.Context) {

	params := context.Request.URL.Query()
	cleanParams := getCleanParams(params)

	service := models.GetServiceById(cleanParams)

	context.JSON(http.StatusOK, gin.H{"service": service})
}

func CreateService(context *gin.Context) {
	var service models.Service
	_ = context.ShouldBindJSON(&service)

	serviceObjId := models.CreateService(&service)

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Object with id %v was created", serviceObjId)})
}

func UpdateServices(context *gin.Context) {
	var inCommingService models.Service
	_ = context.ShouldBindJSON(&inCommingService)

	var oldService models.Service
	cleanParams := getCleanParams(context.Request.URL.Query())
	oldService = models.GetServiceById(cleanParams)

	updatedId := models.UpdateService(inCommingService, oldService)

	context.JSON(http.StatusOK, gin.H{"updated obj ": updatedId})

}

func DeleteService(context *gin.Context) {
	cleanParams := getCleanParams(context.Request.URL.Query())

	id := cleanParams["id"]

	if id != nil && reflect.ValueOf(id).Kind() == reflect.Int {
		iAreaId, _ := id.(int) // Very IMPORTANT
		models.DeleteService(iAreaId)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "some thing wron in deleting"})
	}

}
