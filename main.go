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

	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	routes.SetupRoutes(app, userController, logger)
	err := app.Listen("localhost:3001")
	helper.PanicIfError(err);
}