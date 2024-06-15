package main

import (
	"be-api-go/app"
	"be-api-go/database"
)

func main() {
	database.ConnectionDB()
	database.Migrate()
	app.Serve()
}
