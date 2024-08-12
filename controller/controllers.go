package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindOne(ctx *fiber.Ctx) error
}

type DataKeluargaController interface {
	FindOne(ctx *fiber.Ctx) error
}

type DataLingkunganController interface {
	// FindOneWithId(ctx *fiber.Ctx, id int32) error
}

type Controllers struct {
	UserController           UserController
	DataLingkunganController DataLingkunganController
	DataKeluargaController   DataKeluargaController
}