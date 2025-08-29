package main

import (
	"go-project/api/rest/controllers"
	"go-project/api/rest/routes"
	"go-project/pkg/repositories"
	"go-project/pkg/services"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize dependencies
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)

	// Create UserController instance
	userController := controllers.NewUserController(userService)

	// Register routes with dependency-injected controller
	routes.RegisterUserRoutes(app, userController)

	// Start the server
	app.Listen(":4001")
}
