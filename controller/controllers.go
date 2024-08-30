package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindOne(ctx *fiber.Ctx) error
}

type DataKeluargaController interface {
	FindOne(ctx *fiber.Ctx) error
	AddKeluarga(ctx *fiber.Ctx) error
	GetTotalKeluarga(ctx *fiber.Ctx) error
}

type WealthController interface {
	GetTotal(ctx *fiber.Ctx) error
}

type DataAnggotaController interface {
	AddAnggota(ctx *fiber.Ctx) error
	UpdateAnggota(ctx *fiber.Ctx) error
	GetTotalAnggota(ctx *fiber.Ctx) error
}

type TransactionHistoryController interface {
	GetTotalIncome(ctx *fiber.Ctx) error
	GetTotalOutcome(ctx *fiber.Ctx) error
}

type DataLingkunganController interface {
	// FindOneWithId(ctx *fiber.Ctx, id int32) error
}

type Controllers struct {
	UserController               UserController
	DataLingkunganController     DataLingkunganController
	DataKeluargaController       DataKeluargaController
	WealthController             WealthController
	TransactionHistoryController TransactionHistoryController
	DataAnggotaController        DataAnggotaController
}
