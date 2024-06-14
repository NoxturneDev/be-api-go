package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	var err error
	dsn := "host=localhost port=5432 dbname=technospace_test user=postgres password=root connect_timeout=10 sslmode=prefer"
	dialect := postgres.Open(dsn)
	DB, err = gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		fmt.Printf("Failed to connect to the database: %v\n", err)
		panic("Failed to connect to the database")
	}
}
