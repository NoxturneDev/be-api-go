package model

import (
	"be-api-go/database"
	"fmt"
	"gorm.io/gorm"
)

var db = database.DB

type User struct {
	*gorm.Model
	Username string `json:"name"`
	Password string `json:"password"`
}

type Customer struct {
	*gorm.Model
	PhoneNumber string `json:"phone_number"`
}

func migrate() error {
	fmt.Println("Migrating the schema...")

	err := db.AutoMigrate(&User{}, &Customer{})

	if err != nil {
		fmt.Println("Error migrating the schema")
		return err
	}

	fmt.Println("Schema migrated successfully")
	return nil
}
