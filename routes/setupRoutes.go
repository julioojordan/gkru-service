package routes

import (
	"fmt"
	"gkru-service/authentication"
	"gkru-service/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App, controller controller.UserController) {
	//set middleware cors origin
	app.Use(cors.New())
    app.Post("/login", controller.FindOne)
	app.Get("/testAuth", func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			fmt.Println("missing Auth HEader")
			return c.SendString("unauthorized")
		}
		tokenString := auth[len("Bearer "):]
		err := authentication.VerifyToken(tokenString)
		if err != nil {
			fmt.Println("invalid token")
			return c.SendString("unauthorized 2")
		}
		
		return c.SendString("Berhasil !")
	})
}
