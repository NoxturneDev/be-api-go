package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetCustomers(c *fiber.Ctx) error {
	var customers []model.Customers
	database.DB.Find(&customers)
	return c.JSON(customers)
}

func GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer model.Customers

	if err := database.DB.First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": customer,
	})
}

func DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer model.Customers
	if err := database.DB.First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Customer not found"})
	}
	if err := database.DB.Delete(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Customer not deleted"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Customer deleted successfully"})

}
