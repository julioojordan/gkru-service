package service

import (
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	FindOne(ctx *fiber.Ctx) (interface{}, error)
}

type WealthService interface {
	GetTotal(ctx *fiber.Ctx) (interface{}, error)
}

type DataAnggotaService interface {
	AddAnggota(ctx *fiber.Ctx) (interface{}, error)
	UpdateAnggota(ctx *fiber.Ctx) (interface{}, error)
	GetTotalAnggota(ctx *fiber.Ctx) (interface{}, error)
	DeleteOneAnggota(ctx *fiber.Ctx) (interface{}, error)
	DeleteBulkAnggota(ctx *fiber.Ctx) (interface{}, error)
	FindAll(ctx *fiber.Ctx) (interface{}, error)
	FindOne(ctx *fiber.Ctx) (interface{}, error)
}

type TransactionHistoryService interface {
	GetTotalIncome(ctx *fiber.Ctx) (interface{}, error)
	GetTotalOutcome(ctx *fiber.Ctx) (interface{}, error)
}

type DataKeluargaService interface {
	FindOne(ctx *fiber.Ctx) (interface{}, error)
	FindAll(ctx *fiber.Ctx) (interface{}, error)
	AddKeluarga(ctx *fiber.Ctx) (interface{}, error)
	GetTotalKeluarga(ctx *fiber.Ctx) (interface{}, error)
	UpdateDataKeluarga(ctx *fiber.Ctx) (interface{}, error)
	DeleteDataKeluarga(ctx *fiber.Ctx) (interface{}, error)
}

type DataLingkunganService interface {
	FindOneWithParam(ctx *fiber.Ctx) (interface{}, error)
	FindAll(ctx *fiber.Ctx) (interface{}, error)
	Add(ctx *fiber.Ctx) (interface{}, error)
	Update(ctx *fiber.Ctx) (interface{}, error)
	DeleteOne(ctx *fiber.Ctx) (interface{}, error)
}

type Services struct {
	UserService           UserService
	DataLingkunganService DataLingkunganService
	DataKeluargaService   DataKeluargaService
}
