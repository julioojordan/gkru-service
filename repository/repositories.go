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
}

type DataKeluargaRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaFinal, error)
	GetTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalKeluarga, error)
}

type DataAnggotaRepository interface {
	// FindKeluargaAnggotaRel(relId int32, tx *sql.Tx) ([]entity.DataAnggota, error)
	GetTotalAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalAnggota, error)
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
