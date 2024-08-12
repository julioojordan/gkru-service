package routes

import (
	"fmt"
	"gkru-service/authentication"
	"gkru-service/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, Customlogger *logrus.Logger) {
	// =========== SETUP MIDDLEWARE ===============
	//set middleware cors origin
	app.Use(cors.New())
	//logger untuk tiap endpoint
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	//setup component logger
	app.Use(func(ctx *fiber.Ctx) error {
        ctx.Locals("logger", Customlogger)
        return ctx.Next()
    })
	// =========== SETUP MIDDLEWARE ===============


	// =========== SETUP ROUTE ===============
    // app.Post("/login", controller.FindOne)
	app.Post("/login", func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.FindOne(ctx)
	})

	app.Get("/Keluarga", func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.FindOne(ctx)
	})
	app.Get("/testAuth", func(c *fiber.Ctx) error {
		auth := c.Get("Authorization")
		if auth == "" {
			fmt.Println("missing Auth Header")
			return c.SendString("unauthorized")
		}
		tokenString := auth[len("Bearer "):]
		err := authentication.VerifyToken(tokenString)
		if err != nil {
			fmt.Println("invalid token")
			return c.SendString("unauthorized 2")
		}
		logger, _ := c.Locals("logger").(*logrus.Logger)
		logger.Info("Berhasil !")
		
		return c.SendString("Berhasil !")
	})
	// =========== SETUP ROUTE ===============
}
