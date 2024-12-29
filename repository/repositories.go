package repository

//interface repository

import (
	"database/sql"
	"gkru-service/entity"

	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.User, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.User, error)
	Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error)
	DeleteOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error)
}

type WealthRepository interface {
	GetTotal(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalWealth, error)
}

type TransactionHistoryRepository interface {
	GetTotalIncome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error)
	GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error)
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.ThFinal, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error)
	FindAllWithIdKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error)
	FindAllWithKeluargaContext(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal2, error)
	FindAllHistoryWithTimeFilter(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error)
	FindAllSetoran(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal2, error)
	FindByGroup(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.UpdatedThFinal, error)
	Delete(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error)
	Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.CreatedTh, error)
	AddBatch(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.CreatedTh, error)
}

type DataKeluargaRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) (entity.DataKeluargaFinal, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) ([]entity.DataKeluargaFinal, error)
	GetTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalKeluarga, error)
	AddKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaRaw, error)
	UpdateDataKeluarga(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) (entity.UpdatedDataKeluarga, error)
	DeleteDataKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DeletedDataKeluarga, error)
	CountKeluargaWithParam(ctx *fiber.Ctx, tx *sql.Tx, from string, id int32) (entity.TotalInt, error)
}

type DataAnggotaRepository interface {
	GetTotalAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalAnggota, error)
	AddAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggota, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataAnggotaComplete, error)
	FindAllWithIdKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataAnggotaComplete, error)
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaComplete, error)
	UpdateAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithStatus, error)
	UpdateKeteranganAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithKeteranganOnly, error)
	DeleteOneAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdDataAnggota, error)
	DeleteBulkAnggota(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.IdDataAnggota, error)
}

type DataAnggotaKeluargaRelRepository interface {
	FindKeluargaAnggotaRel(id int32,  db *sql.DB) ([]entity.DataAnggotaWithKeluargaRel, error)
}

type DataLingkunganRepository interface {
	FindOneById(id int32, tx *sql.Tx) (entity.DataLingkungan, error)
	FindOneByIdSeparateTx(id int32, db *sql.DB) (entity.DataLingkungan, error)
	FindOneWithParam(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkungan, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataLingkungan, error)
	FindAllWithTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataLingkunganWithTotalKeluarga, error)
	Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkunganWithIdWilayah, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkunganWithIdWilayah, error)
	DeleteOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdDataLingkungan, error)
	GetTotalLingkungan(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalInt, error)
	CountLingkunganWithIdWilayah(ctx *fiber.Ctx, tx *sql.Tx, idWilayah int32) (entity.TotalInt, error)
}

type DataWilayahRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error)
	FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataWilayah, error)
	Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error)
	Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error)
	DeleteOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error)
	GetTotalWilayah(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalInt, error)
}

type Repositories struct {
	DataLingkunganRepository         DataLingkunganRepository
	DataWilayahRepository            DataWilayahRepository
	DataKeluargaRepository           DataKeluargaRepository
	DataAnggotaRepository            DataAnggotaRepository
	DataAnggotaKeluargaRelRepository DataAnggotaKeluargaRelRepository
}
