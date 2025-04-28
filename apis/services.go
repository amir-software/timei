package apis

import (
	"fmt"
	"golang-temp/models"
	"net/http"
	"path/filepath"
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

// / Center

func GetCentersAPI(context *gin.Context) {
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

	centers := models.GetCenters(cleanParams)

	pagniatedService, _ := paginate(page, pageSize, centers)

	context.JSON(http.StatusOK, gin.H{"centers": pagniatedService})
}

func GetcentersByIdCenterAPI(context *gin.Context) {

	params := context.Request.URL.Query()
	cleanParams := getCleanParams(params)

	center := models.GetCenterById(cleanParams)

	context.JSON(http.StatusOK, gin.H{"center": center})
}

func CreateCentersAPI(context *gin.Context) {

	file, err := context.FormFile("Avater")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Image file is required"})
		return
	}
	codeStr := context.PostForm("Code")

	code, _ := strconv.Atoi(codeStr)
	title := context.PostForm("Title")
	Address := context.PostForm("Address")
	PhoneNumber := context.PostForm("PhoneNumber")

	// Save the uploaded file
	avatarPath := filepath.Join("uploads", file.Filename)
	if err := context.SaveUploadedFile(file, avatarPath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save file"})
		return
	}
	center := models.Center{
		Code:        code,
		Title:       title,
		Address:     Address,
		PhoneNumber: PhoneNumber,
		Avater:      avatarPath,
	}

	centerObjId := models.CreateCenter(&center)

	context.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Center with id %v was created", centerObjId)})
}

func UpdateCenterAPI(context *gin.Context) {
	var inCommingCenter models.Center
	_ = context.ShouldBindJSON(&inCommingCenter)

	var oldCenter models.Center
	cleanParams := getCleanParams(context.Request.URL.Query())
	oldCenter = models.GetCenterById(cleanParams)

	updatedId := models.UpdateCenter(inCommingCenter, oldCenter)

	context.JSON(http.StatusOK, gin.H{"updated obj ": updatedId})

}

func DeleteCenterAPI(context *gin.Context) {
	cleanParams := getCleanParams(context.Request.URL.Query())

	id := cleanParams["id"]

	if id != nil && reflect.ValueOf(id).Kind() == reflect.Int {
		iAreaId, _ := id.(int) // Very IMPORTANT
		models.DeleteCenter(iAreaId)
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"message": "some thing wron in deleting"})
	}

}

func GetServiceInstanceAPI(context *gin.Context) {
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

	servicesInstances := models.GetServiceInstance(cleanParams)

	pagniatedServiceInstance, _ := paginate(page, pageSize, servicesInstances)

	context.JSON(http.StatusOK, gin.H{"serviceInstance": pagniatedServiceInstance})

}

func CreateServiceInstance(context *gin.Context) {
	var serviceIns models.ServiceInstance
	err := context.ShouldBindJSON(&serviceIns)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse body"})
	}
	serviceInsId := models.CreateServiceInstance(&serviceIns)

	context.JSON(http.StatusOK, gin.H{"service Ins id": serviceInsId})
}
