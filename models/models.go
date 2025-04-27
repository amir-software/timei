package models

import "gorm.io/gorm"

type Center struct {
	gorm.Model
	Code        int
	Title       string
	Address     string
	PhoneNumber string
}

type Service struct {
	gorm.Model
	Title       string `json:"title" validate:"required"`
	Description string `json:"custom_description"`
	IsActive    bool   `json:"-"`
	CenterId    uint   `gorm:"TYPE:integer REFERENCES Center"`
}
