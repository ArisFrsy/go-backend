package fiberMain

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	auth "github.com/ArisFrsy/go-backend/API/auth"
	protected "github.com/ArisFrsy/go-backend/API/protected"
	"github.com/ArisFrsy/go-backend/API/public"
)

func FiberHandler() {
	app := fiber.New()

	// Public route
	app.Post("/login", auth.LoginHandler)
	app.Post("/register", auth.RegisterHandler)

	// Unauthenticated route
	app.Get("/", public.AccessibleHandler)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Restricted Routes
	app.Get("/restricted", protected.RestrictedHandler)

	// Get Users
	app.Get("/users", protected.UsersHandler)
	app.Put("/users", protected.UpdateUserHandler)
	app.Delete("/users/:id", protected.DeleteUserHandler)

	// Upload File
	app.Post("/upload", protected.FileUploadHandler)

	app.Listen(":3000")
}
