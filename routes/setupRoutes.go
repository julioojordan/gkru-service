package routes

import (
	"gkru-service/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, controller controller.UserController) {
    app.Get("/login", controller.FindOne)
}
