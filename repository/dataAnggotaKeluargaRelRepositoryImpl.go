package repository

import (
	"database/sql"
	"fmt"

	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type dataAnggotaKeluargaRelRepositoryImpl struct {
}

func NewDataAnggotaKeluargaRelRepository(db *sql.DB) DataAnggotaKeluargaRelRepository {
	return &dataAnggotaKeluargaRelRepositoryImpl{}
}

func (repository *dataAnggotaKeluargaRelRepositoryImpl) FindKeluargaAnggotaRel(id int32, tx *sql.Tx) ([]entity.DataAnggotaWithKeluargaRel, error) {
	sqlScript := "SELECT a.id, a.hubungan, a.id_anggota, b.nama_lengkap, b.tanggal_lahir, b.tanggal_baptis, b.keterangan FROM keluarga_anggota_rel a JOIN data_anggota b ON a.id_anggota = b.id WHERE a.id_keluarga = ?"

	rows, err := tx.Query(sqlScript, id)
	helper.PanicIfError(err)
	defer rows.Close()

	
	fmt.Println("FindKeluargaAnggotaRel 1")

	var results []entity.DataAnggotaWithKeluargaRel
	for rows.Next() {
		data := entity.DataAnggotaWithKeluargaRel{}
		err := rows.Scan(&data.Id, &data.Hubungan, &data.IdAnggota, &data.NamaLengkap, &data.TanggalLahir, &data.TanggalBaptis, &data.Keterangan)
		fmt.Println("data", data)
		helper.PanicIfError(err)
		results = append(results, data)
		fmt.Println("results", results)
	}
	
	fmt.Println("FindKeluargaAnggotaRel 2", results)

	if len(results) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Data is not found")
	}

	return results, nil
}
