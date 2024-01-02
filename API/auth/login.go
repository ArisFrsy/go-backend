package auth

import (
	"log"
	"time"

	db "github.com/ArisFrsy/go-backend/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Username string
	Email    string
	Password []byte // Change the type to []byte for storing hashed password
}

func LoginHandler(c *fiber.Ctx) error {
	username := c.FormValue("username")
	pass := c.FormValue("pass")

	var user User

	// Find the user based on username
	if result := db.DB.Where("username = ?", username).First(&user); result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	// Compare the hashed password with the input password
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(pass)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid username or password",
		})
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name":     user.Name,
		"username": user.Username,
		"email":    user.Email,
		"admin":    true,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	log.Println(user)

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}

func RegisterHandler(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	username := c.FormValue("username")
	pass := c.FormValue("pass")

	// Validate that name, email, and pass are not empty
	if name == "" || email == "" || pass == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name, email, and password must not be empty",
		})
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash the password",
		})
	}

	user := User{
		Name:     name,
		Email:    email,
		Username: username,
		Password: hashedPassword,
	}

	db.DB.Create(&user)

	return c.JSON(user)
}
