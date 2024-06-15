package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetSellers(c *fiber.Ctx) error {

	var sellers []model.Sellers
	database.DB.Find(&sellers)
	return c.JSON(fiber.Map{
		"data": sellers,
	})
}

func GetSellerChatRooms(c *fiber.Ctx) error {
	id := c.Params("id")
	var seller model.Sellers

	if err := database.DB.Preload("Products").Preload("ChatRooms").First(&seller, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Seller not found"})
	}
	return c.JSON(seller)
}

func UpdateSeller(c *fiber.Ctx) error {
	id := c.Params("id")
	var seller model.Sellers

	if err := database.DB.First(&seller, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Seller not found"})
	}
	if err := c.BodyParser(&seller); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	database.DB.Save(&seller)
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   seller,
	})
}

func DeleteSeller(c *fiber.Ctx) error {
	id := c.Params("id")
	var seller model.Sellers
	if err := database.DB.First(&seller, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Seller not found"})
	}
	database.DB.Delete(&seller)
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   seller,
	})
}
func GetSeller(c *fiber.Ctx) error {
	id := c.Params("id")
	var seller model.Sellers
	if err := database.DB.First(&seller, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Seller not found"})
	}
	return c.JSON(fiber.Map{
		"data": seller,
	})
}

func CreateSeller(c *fiber.Ctx) error {
	var seller model.Sellers
	if err := c.BodyParser(&seller); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	database.DB.Create(&seller)
	return c.JSON(seller)
}
