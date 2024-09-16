package main

import (
	"gkru-service/controller"
	"gkru-service/db"
	"gkru-service/helper"
	"gkru-service/repository"
	"gkru-service/routes"
	"gkru-service/service"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	db := db.NewDB(logger)
	validate := validator.New()

	logger.Info("Database connected successfully")

	//prepare service
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)

	dataLingkunganRepository := repository.NewDataLingkunganRepository(db)
	dataLingkunganService := service.NewDataLingkunganService(dataLingkunganRepository, db, validate)
	dataLingkunganController := controller.NewDataLingkunganController(dataLingkunganService)

	dataKeluargaRepository := repository.NewDataKeluargaRepository(db)
	dataKeluargaService := service.NewDataKeluargaService(dataKeluargaRepository, db, validate)
	dataKeluargaController := controller.NewDataKeluargaController(dataKeluargaService)

	dataAnggotaRepository := repository.NewDataAnggotaRepository(db)
	dataAnggotaService := service.NewDataAnggotaService(dataAnggotaRepository, db, validate)
	dataAnggotaController := controller.NewDataAnggotaController(dataAnggotaService)

	dataAnggotaKeluargaRelRepository := repository.NewDataAnggotaKeluargaRelRepository(db)

	wealthRepository := repository.NewWealthRepository(db)
	wealthService := service.NewWealthService(wealthRepository, db, validate)
	wealthController := controller.NewWealthController(wealthService)

	transactionHistoryRepository := repository.NewTransactionHistoryRepository(db)
	transactionHistoryService := service.NewTransactionHistoryService(transactionHistoryRepository, db, validate)
	transactionHistoryController := controller.NewTransactionHistoryController(transactionHistoryService)

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	//setup context
	app.Use(func(c *fiber.Ctx) error {
		controllers := controller.Controllers{
			UserController:               userController.(*controller.UserControllerImpl),
			DataLingkunganController:     dataLingkunganController.(*controller.DataLingkunganControllerImpl),
			DataKeluargaController:       dataKeluargaController.(*controller.DataKeluargaControllerImpl),
			WealthController:             wealthController.(*controller.WealthControllerImpl),
			TransactionHistoryController: transactionHistoryController.(*controller.TransactionHistoryControllerImpl),
			DataAnggotaController:        dataAnggotaController.(*controller.DataAnggotaControllerImpl),
		}
		services := service.Services{
			DataLingkunganService: dataLingkunganService.(*service.DataLingkunganServiceImpl),
		}
		//gak tau ini kenapa ga perlu di parse ke dataLingkunganRepositoryImpl -> padahal yang atasnya bisa
		repositories := repository.Repositories{
			DataLingkunganRepository:         dataLingkunganRepository,
			DataKeluargaRepository:           dataKeluargaRepository,
			DataAnggotaKeluargaRelRepository: dataAnggotaKeluargaRelRepository,
			DataAnggotaRepository:            dataAnggotaRepository,
		}
		c.Locals("controllers", controllers)
		c.Locals("services", services)
		c.Locals("repositories", repositories)
		return c.Next()
	})

	routes.SetupRoutes(app, logger)
	err := app.Listen("localhost:3001")
	helper.PanicIfError(err)
}
