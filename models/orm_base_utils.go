package models

import "golang-temp/database"

func GetServiceById(params map[string]interface{}) Service {
	var services []Service
	database.DataBase.Where(params).Find(&services)

	return services[0]
}

func GetServices(params map[string]interface{}) []Service {
	var services []Service

	orderField := params["order"]

	if orderField != nil{
		database.DataBase.Where(params).Find(&services).Order(orderField)	
	}else{
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
