package main

import (
	"github.com/clrql/backend-directory/crud"
	"github.com/clrql/backend-directory/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Establish a connection to the database
	_, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// Initialize Fiber app
	app := fiber.New(fiber.Config{})

	// Define CRUD operations for user endpoints
	crud.UsersCrud(app.Group("/users"))

	// Start the server on port 3030
	app.Listen(":3030")
}
