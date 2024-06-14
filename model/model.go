package model

import (
	"gorm.io/gorm"
)

type Users struct {
	*gorm.Model
	Username string `json:"name"`
	Password string `json:"password"`
}

type Customers struct {
	*gorm.Model
	PhoneNumber string `json:"phone_number"`
}
