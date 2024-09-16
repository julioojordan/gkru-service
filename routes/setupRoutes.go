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
	app.Get("/keluarga/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.GetTotalKeluarga(ctx)
	})
	app.Post("/keluarga/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.AddKeluarga(ctx)
	})
	app.Patch("/keluarga/:idKeluarga/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.UpdateDataKeluarga(ctx)
	})
	app.Patch("/keluarga/:idKeluarga/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.DeleteDataKeluarga(ctx)
	})
	app.Get("/keluarga/:idKeluarga", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		fmt.Println("masuk /Keluarga/:idKeluarga")
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.FindOne(ctx)
	})
	app.Get("/keluarga", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		return dataKeluargaController.FindAll(ctx)
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
	app.Patch("/anggota/:idAnggota/update", func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.UpdateAnggota(ctx)
	})
	app.Delete("/anggota/:idAnggota/delete", func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.DeleteOneAnggota(ctx)
	})
	app.Post("/anggota/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.AddAnggota(ctx)
	})
	app.Post("/anggota/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.DeleteBulkAnggota(ctx)
	})
	app.Get("/anggota/:idAnggota", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.FindOne(ctx)
	})
	app.Get("/anggota", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.FindAll(ctx)
	})

	// =========== LINGKUNGAN ===============
	app.Patch("/lingkungan/:idLingkungan/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.Update(ctx)
	})
	app.Delete("/anggota/:idLingkungan/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.DeleteOne(ctx)
	})
	app.Get("/lingkungan/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindAll(ctx)
	})
	app.Post("/lingkungan/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.Add(ctx)
	})
	app.Get("/lingkungan/:idLingkungan", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindOneWithParam(ctx)
	})
	app.Get("/lingkungan", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindAll(ctx)
	})
	// =========== SETUP ROUTE ===============
}
