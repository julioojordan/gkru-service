package routes

import (
	"gkru-service/middlewares"
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
	app.Get("/wealth/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		wealthController := ctx.Locals("controllers").(controller.Controllers).WealthController
		return wealthController.GetTotal(ctx)
	})
	app.Get("/history/getTotalIncome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalIncome(ctx)
	})
	app.Get("/history/getTotalOutcome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalOutcome(ctx)
	})
	// =========== SETUP ROUTE ===============
}
