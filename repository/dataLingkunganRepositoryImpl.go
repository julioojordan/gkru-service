package repository

import (
	"database/sql"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type dataLingkunganRepositoryImpl struct {
}

func NewDataLingkunganRepository(db *sql.DB) DataLingkunganRepository {
	return &dataLingkunganRepositoryImpl{}
}

func (repository *dataLingkunganRepositoryImpl) FindOneById(id int32, tx *sql.Tx) (entity.DataLingkungan, error) {
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id WHERE l.id = ?"
	result, err :=tx.Query(sqlScript, id)
	helper.PanicIfError(err);
	defer result.Close()
	
	raw := entity.DataLingkunganRaw{}
	if result.Next(){
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah)
		helper.PanicIfError(err)
		wilayah := entity.DataWilayah{
			Id: raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkungan{
			Id: raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah: wilayah,
		}
		return lingkungan, nil
	} else{
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}
}
