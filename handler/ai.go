package handler

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func TestAIconnection(c *fiber.Ctx) error {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyANaDL7jL7ZhbbAzz05JEBAibYvyzVuK7c"))
	if err != nil {
		fmt.Println(err)
	}

	defer client.Close()
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a AI and magic"))
	if err != nil {
		fmt.Println("an error occurred")
		fmt.Println(err)
	}

	return c.JSON(fiber.Map{"message": resp})
}
