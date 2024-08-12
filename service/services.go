package service

import (

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	FindOne(ctx *fiber.Ctx) (interface{}, error)
}

type DataKeluargaService interface {
	FindOne(ctx *fiber.Ctx) (interface{}, error)
}

type DataLingkunganService interface {
	FindOneById(ctx *fiber.Ctx, id int32) (interface{}, error)
}

type Services struct {
	UserService           UserService
	DataLingkunganService DataLingkunganService
	DataKeluargaService   DataKeluargaService
}