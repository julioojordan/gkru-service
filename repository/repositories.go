package repository

//interface repository

import (
	"database/sql"
	"gkru-service/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error)
}

type WealthRepository interface {
	GetTotal(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalWealth, error)
}

type TransactionHistoryRepository interface {
	GetTotalIncome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error)
	GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error)
	// GetHistory(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error)
}

type DataKeluargaRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaFinal, error)
	GetTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalKeluarga, error)
	AddKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaRaw, error)
	UpdateDataKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.UpdatedDataKeluarga, error)
}

type DataAnggotaRepository interface {
	GetTotalAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalAnggota, error)
	AddAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggota, error)
	UpdateAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithStatus, error)
	UpdateKeteranganAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithKeteranganOnly, error)
}

type DataAnggotaKeluargaRelRepository interface {
	FindKeluargaAnggotaRel(id int32, tx *sql.Tx) ([]entity.DataAnggotaWithKeluargaRel, error)
}

type DataLingkunganRepository interface {
	FindOneById(id int32, tx *sql.Tx) (entity.DataLingkungan, error)
}

type DataWilayahRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error)
}

type Repositories struct {
	DataLingkunganRepository          DataLingkunganRepository
	DataAnggotaRepository             DataAnggotaRepository
	DataAnggotaKeluargaRelRepository DataAnggotaKeluargaRelRepository
}
