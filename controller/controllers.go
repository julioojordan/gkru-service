package controller

import "github.com/gofiber/fiber/v2"

type UserController interface {
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Add(ctx *fiber.Ctx) error
	DeleteOne(ctx *fiber.Ctx) error
}

type DataKeluargaController interface {
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	AddKeluarga(ctx *fiber.Ctx) error
	GetTotalKeluarga(ctx *fiber.Ctx) error
	GetTotalKeluargaWithFilter(ctx *fiber.Ctx) error
	UpdateDataKeluarga(ctx *fiber.Ctx) error
	DeleteDataKeluarga(ctx *fiber.Ctx) error
}

type WealthController interface {
	GetTotal(ctx *fiber.Ctx) error
}

type DataAnggotaController interface {
	AddAnggota(ctx *fiber.Ctx) error
	UpdateAnggota(ctx *fiber.Ctx) error
	GetTotalAnggota(ctx *fiber.Ctx) error
	DeleteOneAnggota(ctx *fiber.Ctx) error
	DeleteBulkAnggota(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindAllWithIdKeluarga(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
}

type TransactionHistoryController interface {
	GetTotalIncome(ctx *fiber.Ctx) error
	GetTotalOutcome(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindAllWithIdKeluarga(ctx *fiber.Ctx) error
	FindAllWithKeluargaContext(ctx *fiber.Ctx) error
	FindAllHistoryWithTimeFilter(ctx *fiber.Ctx) error
	FindAllSetoran(ctx *fiber.Ctx) error
	FindByGroup(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	Add(ctx *fiber.Ctx) error
	AddBatch(ctx *fiber.Ctx) error
}

type DataLingkunganController interface {
	FindOneWithParam(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	FindAllWithTotalKeluarga(ctx *fiber.Ctx) error
	Add(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	DeleteOne(ctx *fiber.Ctx) error
	GetTotalLingkungan(ctx *fiber.Ctx) error
}

type DataWilayahController interface {
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
	Add(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	DeleteOne(ctx *fiber.Ctx) error
	GetTotalWilayah(ctx *fiber.Ctx) error
}

type Controllers struct {
	UserController               UserController
	DataLingkunganController     DataLingkunganController
	DataWilayahController        DataWilayahController
	DataKeluargaController       DataKeluargaController
	WealthController             WealthController
	TransactionHistoryController TransactionHistoryController
	DataAnggotaController        DataAnggotaController
}
