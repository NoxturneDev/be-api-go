package app

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Serve() {
	// Call the function
	fmt.Println("Server started...")

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	Routes(app)
	app.Listen(":3001")
}
