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

type DataKeluargaRepository interface {
	FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaFinal, error)
}

type DataAnggotaRepository interface {
	// FindKeluargaAnggotaRel(relId int32, tx *sql.Tx) ([]entity.DataAnggota, error)
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
