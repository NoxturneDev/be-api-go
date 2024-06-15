package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(c *fiber.Ctx) error {

	var products []model.Products
	database.DB.Find(&products)
	return c.JSON(fiber.Map{
		"data": products,
	})
}

func GetProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product model.Products
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(fiber.Map{
		"data": product,
	})
}

func CreateProduct(c *fiber.Ctx) error {

	var product model.Products
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	database.DB.Create(&product)
	return c.JSON(fiber.Map{"data": product})

}

func UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product model.Products
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	database.DB.Save(&product)
	return c.JSON(fiber.Map{"data": product})
}

func DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var product model.Products
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	database.DB.Delete(&product)
	return c.JSON(fiber.Map{"data": product})
}
