package handler

import (
	"be-api-go/database"
	"be-api-go/model"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []model.Users

	database.DB.Find(&users)

	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.Users

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	newUser := model.Users{
		Username: req.Username,
		Password: req.Password,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "User created successfully"})
}

func UpdateUser(c *fiber.Ctx) error {
	var req struct {
		UserId   int    `json:"user_id"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	var user model.Users
	if err := database.DB.First(&user, req.UserId).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	user.Username = req.Username
	user.Password = req.Password

	if err := database.DB.Save(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "User updated successfully"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user model.Users

	if err := database.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})

}

func LoginUser(c *fiber.Ctx) error {
	var requestLogin struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requestLogin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.DB.Where("username = ?", requestLogin.Username).First(&model.Users{}).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": "User not found",
		})
	}
	if err := database.DB.Where("username = ? AND password = ?", requestLogin.Username, requestLogin.Password).First(&model.Users{}).Error; err != nil {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login success",
	})
}

func LogoutUser(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logout success",
	})
}
