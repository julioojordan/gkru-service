package routes

import (
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

	// =========== SETUP STATIC ==========
	app.Static("/uploads", "./uploads")

	// =========== SETUP ROUTE ===============
	app.Post("/login", func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.FindOne(ctx)
	})
	app.Get("/user", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.FindAll(ctx)
	})
	app.Post("/user/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.Add(ctx)
	})
	app.Patch("/user/:idUser/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.Update(ctx)
	})
	app.Delete("/user/:idUser/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		userController := ctx.Locals("controllers").(controller.Controllers).UserController
		return userController.DeleteOne(ctx)
	})

	// =========== KELUARGA ===============
	app.Get("/keluarga/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataKeluargaController := ctx.Locals("controllers").(controller.Controllers).DataKeluargaController
		idWilayah := ctx.Query("idWilayah")
		idLingkungan := ctx.Query("idLingkungan")
		if idWilayah != "" || idLingkungan != "" {
			return dataKeluargaController.GetTotalKeluargaWithFilter(ctx)
		} else {
			return dataKeluargaController.GetTotalKeluarga(ctx)
		}
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
	app.Post("/history/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.Add(ctx)
	})
	app.Post("/history/addIuran", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.AddBatch(ctx)
	})
	app.Get("/history/getTotalIncome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalIncome(ctx)
	})
	app.Get("/history/getTotalOutcome", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.GetTotalOutcome(ctx)
	})
	app.Get("/history", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		idKeluarga := ctx.Query("idKeluarga")
		if idKeluarga != "" {
			return transactionHistoryController.FindAllWithIdKeluarga(ctx)
		} else {
			return transactionHistoryController.FindAll(ctx)
		}
	})
	app.Get("/historyByGroup", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.FindByGroup(ctx)
	})
	app.Get("/historyWithContext", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.FindAllWithKeluargaContext(ctx)
	})
	app.Get("/historyWithTimeFilter", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.FindAllHistoryWithTimeFilter(ctx)
	})
	app.Get("/historySetoran", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.FindAllSetoran(ctx)
	})
	app.Patch("/history/:idTh/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.Update(ctx)
	})
	app.Delete("/history/:idTh/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.Delete(ctx)
	})
	app.Get("/history/:idTh", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		transactionHistoryController := ctx.Locals("controllers").(controller.Controllers).TransactionHistoryController
		return transactionHistoryController.FindOne(ctx)
	})

	// =========== ANGGOTA ===============
	app.Patch("/anggota/:idAnggota/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.UpdateAnggota(ctx)
	})
	app.Delete("/anggota/:idAnggota/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.DeleteOneAnggota(ctx)
	})
	app.Post("/anggota/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.AddAnggota(ctx)
	})
	app.Get("/anggota/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataAnggotaController := ctx.Locals("controllers").(controller.Controllers).DataAnggotaController
		return dataAnggotaController.GetTotalAnggota(ctx)
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
		idKeluarga := ctx.Query("idKeluarga")
		if idKeluarga != "" {
			return dataAnggotaController.FindAllWithIdKeluarga(ctx)
		} else {
			return dataAnggotaController.FindAll(ctx)
		}
	})

	// =========== LINGKUNGAN ===============
	app.Patch("/lingkungan/:idLingkungan/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.Update(ctx)
	})
	app.Delete("/lingkungan/:idLingkungan/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.DeleteOne(ctx)
	})
	app.Post("/lingkungan/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.Add(ctx)
	})
	app.Get("/lingkungan/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.GetTotalLingkungan(ctx)
	})
	app.Get("/lingkungan/:idLingkungan", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindOneWithParam(ctx)
	})
	app.Get("/lingkungan", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindAll(ctx)
	})
	app.Get("/lingkunganWithTotalKeluarga", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataLingkunganController := ctx.Locals("controllers").(controller.Controllers).DataLingkunganController
		return dataLingkunganController.FindAllWithTotalKeluarga(ctx)
	})

	// =========== WILAYAH ===============
	app.Patch("/wilayah/:idWilayah/update", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.Update(ctx)
	})
	app.Delete("/wilayah/:idWilayah/delete", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.DeleteOne(ctx)
	})
	app.Get("/wilayah/getTotal", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.GetTotalWilayah(ctx)
	})
	app.Post("/wilayah/add", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.Add(ctx)
	})
	app.Get("/wilayah/:idWilayah", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.FindOne(ctx)
	})
	app.Get("/wilayah", middlewares.AuthMiddleware, func(ctx *fiber.Ctx) error {
		dataWilayahController := ctx.Locals("controllers").(controller.Controllers).DataWilayahController
		return dataWilayahController.FindAll(ctx)
	})

	// ========== STATIS ============
	app.Static("/uploads", "./uploads")
	// =========== SETUP ROUTE ===============
}
