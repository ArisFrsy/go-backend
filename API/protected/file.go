package protected

import (
	"crypto/rand"
	"fmt"
	"strings"

	db "github.com/ArisFrsy/go-backend/database"
	"github.com/gofiber/fiber/v2"
)

func getRandomName() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func FileUploadHandler(c *fiber.Ctx) error {

	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "No file provided",
		})
	}

	// Generate a unique file name
	fileExt := strings.Split(file.Filename, ".")[1] // Get the file extension
	newFileName := fmt.Sprintf("%s.%s", getRandomName(), fileExt)
	// Save the file to the uploads directory
	err = c.SaveFile(file, fmt.Sprintf("./file/%s", newFileName))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not save the file",
		})
	}

	// Save file details to the database
	newFile := db.File{
		FileName:  file.Filename,
		Path:      fmt.Sprintf("./file/%s", newFileName),
		Extension: fileExt,
	}
	db.DB.Create(&newFile)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Message": "File Uploaded Succesfuly",
		"File":    newFile,
	})
}
