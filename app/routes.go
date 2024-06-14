package app

import (
	"be-api-go/handler"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/api/users", handler.GetUsers)
	app.Get("/api/user/:id", handler.GetUser)
	app.Post("/api/user", handler.CreateUser)
	app.Put("/api/user", handler.UpdateUser)
	app.Delete("/api/user/:id", handler.DeleteUser)

	app.Get("/api/ai/test", handler.TestAIconnection)

	//	ws
	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
		handler.WebsocketHandler(c)
	}))
}
