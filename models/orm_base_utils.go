package models

import (
	"fmt"
	"golang-temp/database"
)

func GetServiceById(params map[string]interface{}) Service {
	var services []Service
	database.DataBase.Where(params).Find(&services)

	return services[0]
}

func GetServices(params map[string]interface{}) []Service {
	var services []Service

	orderField := params["order"]

	if orderField != nil {
		database.DataBase.Where(params).Find(&services).Order(orderField)
	} else {
		delete(params, "order")
		database.DataBase.Where(params).Find(&services)
	}
	return services
}

func CreateService(service *Service) uint {
	newService := Service{Title: service.Title, Description: service.Description, IsActive: service.IsActive}
	database.DataBase.Create(&newService)

	return newService.ID
}

func UpdateService(newService, oldService Service) uint {
	database.DataBase.Model(&oldService).Updates(newService)

	return oldService.ID
}

func DeleteService(id int) string {
	database.DataBase.Delete(&Service{}, id)

	return "Service was deleted"
}

/// Center orm function

func GetCenterById(params map[string]interface{}) Center {
	var centers []Center
	database.DataBase.Where(params).Find(&centers)

	return centers[0]
}

func GetCenters(params map[string]interface{}) []Center {
	var centers []Center

	orderField := params["order"]

	if orderField != nil {
		database.DataBase.Where(params).Find(&centers).Order(orderField)
	} else {
		delete(params, "order")
		database.DataBase.Where(params).Find(&centers)
	}
	return centers
}

func CreateCenter(center *Center) uint {
	newCenter := Center{Code: center.Code, Title: center.Title,
		Address: center.Address, PhoneNumber: center.PhoneNumber, Avater: center.Avater}

	fmt.Println(newCenter)
	database.DataBase.Create(&newCenter)
	return newCenter.ID
}

func UpdateCenter(newCenter, oldCenter Center) uint {
	database.DataBase.Model(&oldCenter).Updates(newCenter)

	return oldCenter.ID
}

func DeleteCenter(id int) string {
	database.DataBase.Delete(&Service{}, id)

	return "Service was deleted"
}
