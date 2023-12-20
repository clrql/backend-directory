package crud

import (
	"fmt"
	"strconv"
	"time"

	"github.com/clrql/backend-directory/database"
	"github.com/clrql/backend-directory/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// UsersCrud defines CRUD operations for users.
func UsersCrud(r fiber.Router) {

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Auto-migrate User model
	db.AutoMigrate(&models.UserModel{})

	// Handle OPTIONS request for user resource (MODEL)
	r.Options("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"name":        "string",
			"birth":       "time.Time",
			"address":     "string",
			"description": "*string",
		})
	})

	// Handle GET request to fetch all users (LIST)
	r.Get("/", func(c *fiber.Ctx) error {
		var users []models.UserModel
		if err := db.Find(&users).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to list.",
			})
		}
		return c.Status(fiber.StatusOK).JSON(users)
	})

	// Handle POST request to create a new user (CREATE)
	r.Post("/", func(c *fiber.Ctx) error {
		var body struct {
			Name        string    `json:"name"`
			Birth       time.Time `json:"birth"`
			Address     string    `json:"address"`
			Description *string   `json:"description"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body.",
			})
		}

		var newUser models.UserModel
		newUser.Name = body.Name
		newUser.Address = body.Address
		newUser.Birth = body.Birth
		newUser.Description = body.Description

		if err := db.Create(&newUser).Error; err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create.",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(newUser)
	})

	// Handle DELETE request to delete all users (DROP)
	r.Delete("/", func(c *fiber.Ctx) error {
		if err := db.Delete(&models.UserModel{}, "1 = 1").Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to drop.",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})

	// Group users by ID
	byId := r.Group("/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid urlparam id.",
			})
		}
		c.Locals("id", id)
		return c.Next()
	})

	// Handle GET request for a specific user by ID (RETRIEVE)
	byId.Get("/", func(c *fiber.Ctx) error {
		id := c.Locals("id").(uint64)
		var user models.UserModel
		err := db.First(&user, id).Error
		switch err {
		case nil:
			break
		case gorm.ErrRecordNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found.",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve.",
			})
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})

	// Handle PUT request to update a specific user by ID (UPDATE)
	byId.Put("/", func(c *fiber.Ctx) error {
		var body struct {
			Name        string    `json:"name"`
			Birth       time.Time `json:"birth"`
			Address     string    `json:"address"`
			Description *string   `json:"description"`
		}

		if err := c.BodyParser(&body); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body.",
			})
		}

		id := c.Locals("id").(uint64)
		var user models.UserModel

		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to put.",
			})
		}

		user.Name = body.Name
		user.Address = body.Address
		user.Birth = body.Birth
		user.Description = body.Description

		if err := db.Save(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to save.",
			})
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	// Handle DELETE request for a specific user by ID (DELETE)
	byId.Delete("/", func(c *fiber.Ctx) error {
		id := c.Locals("id").(uint64)
		if err := db.Delete(&models.UserModel{}, "ID = ?", id).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to delete.",
			})
		}
		return c.SendStatus(fiber.StatusNoContent)
	})

}
