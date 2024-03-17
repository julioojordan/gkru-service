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

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	db := db.NewDB()
	validate := validator.New()

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

	routes.SetupRoutes(app, userController)
	err := app.Listen("localhost:3001")
	helper.PanicIfError(err);
}