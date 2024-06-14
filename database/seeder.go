package database

import (
	"be-api-go/model"
	"fmt"
)

func Migrate() error {
	var err error
	fmt.Println("Migrating the schema...")

	err = DB.AutoMigrate(&model.Users{})
	err = DB.AutoMigrate(&model.Customers{})

	if err != nil {
		fmt.Printf("Error migrating the schema%v", err)
		return err
	}

	fmt.Println("Schema migrated successfully")
	return nil
}
