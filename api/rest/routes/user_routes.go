package routes

import (
	"go-project/api/rest/controllers"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(app *fiber.App, userController *controllers.UserController) {
	api := app.Group("/api")
	api.Post("/users", userController.CreateUser)
}
