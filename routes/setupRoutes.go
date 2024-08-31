package routes

import (
	"fmt"
	"gkru-service/controller"
	"gkru-service/middlewares"

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

	// =========== KELUARGA ===============
	app.Get("/Keluarga/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.GetTotalKeluarga(ctx)
	})
	app.Post("/Keluarga/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.AddKeluarga(ctx)
	})
	app.Patch("/Keluarga/:idKeluarga/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.UpdateDataKeluarga(ctx)
	})
	app.Get("/Keluarga/:idKeluarga", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		fmt.Println("masuk /Keluarga/:idKeluarga")
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.FindOne(ctx)
	})

	// =========== WEALTH ===============
	app.Get("/wealth/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		wealthController := ctx.Locals("controllers").(controller.Controllers).WealthController
		return wealthController.GetTotal(ctx)
	})

	// =========== HISTORY ===============
	app.Get("/history/getTotalIncome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalIncome(ctx)
	})
	app.Get("/history/getTotalOutcome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		// to do need to test this with id_wilayah and id_lingkungan query params
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalOutcome(ctx)
	})

	// =========== ANGGOTA ===============
	app.Patch("/anggota/:idAnggota/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.UpdateAnggota(ctx)
	})
	app.Post("/anggota/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.AddAnggota(ctx)
	})
	app.Get("/anggota/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.GetTotalAnggota(ctx)
	})
	// =========== SETUP ROUTE ===============
}
