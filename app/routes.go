package app

import (
	"be-api-go/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"net/http"
)

func Routes(app *fiber.App) {
	app.Get("/api/users", handler.GetUsers)
	app.Get("/api/user/:id", handler.GetUser)
	app.Post("/api/user", handler.CreateUser)
	app.Put("/api/user", handler.UpdateUser)
	app.Delete("/api/user/:id", handler.DeleteUser)

	//app.Get("/api/ai/test", handler.TestAIconnection)

	//	ws
	//app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
	//	handler.AiWebsocketHandler(c)
	//}))

	app.Get("/ws/ai", func(c *fiber.Ctx) error {
		var writer http.ResponseWriter
		var request *http.Request
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer = w
			request = r
		})(c.Context())

		err := handler.AiChatHandler(writer, request)
		if err != nil {
			return err
		}

		return nil
	})

	// sellers
	app.Get("/api/sellers", handler.GetSellers)
	app.Get("/api/seller/chat/:id", handler.GetSellerChatRooms)
	app.Put("/api/seller", handler.UpdateSeller)
	app.Delete("/api/seller/:id", handler.DeleteSeller)
	app.Get("/api/seller/:id", handler.GetSeller)
	app.Post("/api/seller/create", handler.CreateSeller)

	// products
	app.Get("/api/products", handler.GetProducts)
	app.Get("/api/product/:id", handler.GetProduct)
	app.Post("/api/product", handler.CreateProduct)
	app.Put("/api/product", handler.UpdateProduct)
	app.Delete("/api/product/:id", handler.DeleteProduct)
}
