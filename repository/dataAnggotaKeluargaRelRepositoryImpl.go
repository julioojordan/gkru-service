package repository

import (
	"database/sql"

	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type dataAnggotaKeluargaRelRepositoryImpl struct {
}

func NewDataAnggotaKeluargaRelRepository(db *sql.DB) DataAnggotaKeluargaRelRepository {
	return &dataAnggotaKeluargaRelRepositoryImpl{}
}

func (repository *dataAnggotaKeluargaRelRepositoryImpl) FindKeluargaAnggotaRel(id int32, db *sql.DB) ([]entity.DataAnggotaWithKeluargaRel, error) {
	sqlScript := "SELECT a.id, a.hubungan, a.id_anggota, b.nama_lengkap, b.tanggal_lahir, b.tanggal_baptis, b.keterangan, b.status, b.jenis_kelamin FROM keluarga_anggota_rel a JOIN data_anggota b ON a.id_anggota = b.id WHERE a.id_keluarga = ?"

	rows, err := db.Query(sqlScript, id)
	helper.PanicIfError(err)
	defer rows.Close()

	var results []entity.DataAnggotaWithKeluargaRel
	for rows.Next() {
		data := entity.DataAnggotaWithKeluargaRel{}
		err := rows.Scan(&data.Id, &data.Hubungan, &data.IdAnggota, &data.NamaLengkap, &data.TanggalLahir, &data.TanggalBaptis, &data.Keterangan, &data.Status, &data.JenisKelamin)
		helper.PanicIfError(err)
		results = append(results, data)
	}

	if len(results) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "Data Tidak Ditemukan")
	}

	return results, nil
}
