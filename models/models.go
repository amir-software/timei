package models

import "gorm.io/gorm"

type Center struct {
	gorm.Model
	Code        int
	Title       string
	Address     string
	PhoneNumber string
	Avater      string `validate:"required"`
	Description string
}

type Service struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"custom_description"`
	IsActive    bool   `json:"-"`
	CenterId    uint
}

type ServiceInstance struct {
	gorm.Model
	Title     string
	Image     string
	ServiceId int
	Service   Service `gorm:"foreignKey:ServiceId"`
}

// `gorm:"TYPE:integer REFERENCES Center"`
