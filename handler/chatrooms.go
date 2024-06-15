package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetChatRooms(c *fiber.Ctx) error {
	var chatRooms []model.ChatRooms

	database.DB.Find(&chatRooms)

	return c.JSON(chatRooms)
}

func GetChatRoom(c *fiber.Ctx) error {
	id := c.Params("id")
	var chatRoom model.ChatRooms

	if err := database.DB.First(&chatRoom, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat room not found"})
	}

	return c.JSON(chatRoom)
}

func GetChatRoomBySellerId(c *fiber.Ctx) error {
	id := c.Params("seller_id")
	var chatRoom []model.ChatRooms

	if err := database.DB.Where("seller_id = ?", id).Find(&chatRoom).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Chat room not found"})
	}

	return c.JSON(chatRoom)
}

func CreateChatRoom(c *fiber.Ctx) error {
	var req struct {
		SellerId    int    `json:"seller_id"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newChatRoom := model.ChatRooms{
		PhoneNumber: req.PhoneNumber,
		SellerId:    req.SellerId,
	}

	if err := database.DB.Create(&newChatRoom).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "Chat room created successfully"})
}
