package protected

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	db "github.com/ArisFrsy/go-backend/database"
)

func UsersHandler(c *fiber.Ctx) error {
	userget := c.Locals("user").(*jwt.Token)
	claims := userget.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	var user db.User

	// Find the user based on username
	if result := db.DB.Where("name = ?", name).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	return c.JSON(user)
}

func UpdateUserHandler(c *fiber.Ctx) error {
	// Parse the form values
	id := c.FormValue("id")
	name := c.FormValue("name")
	email := c.FormValue("email")

	// Fetch the user from the database
	var user db.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update the user's name and email if new values are provided in the form
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	// Save the updated user back to the database
	if err := db.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update user"})
	}

	return c.JSON(user)
}

func DeleteUserHandler(c *fiber.Ctx) error {
	// Parse the ID from the request parameters
	id := c.Params("id")

	// Fetch the user from the database
	var user db.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Delete the user
	if err := db.DB.Delete(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete user"})
	}

	return c.JSON(fiber.Map{"message": "User deleted successfully"})
}
