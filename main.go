package main

import (
	"be-api-go/app"
	"be-api-go/database"
)

func main() {
	database.ConnectionDB()
	app.Serve()
}
