package repository

import (
	"database/sql"
	"encoding/json"
	"gkru-service/entity"
	"gkru-service/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type dataKeluargaRepositoryImpl struct {
}

func NewDataKeluargaRepository(db *sql.DB) DataKeluargaRepository {
	return &dataKeluargaRepositoryImpl{}
}

func (repository *dataKeluargaRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) (entity.DataKeluargaFinal, error) {
	dataKeluargaRawScript := "SELECT id, id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, alamat, status, nomor_kk_gereja FROM data_keluarga WHERE id = ?"
	idKeluarga := ctx.Params("idKeluarga")

	result, err := tx.Query(dataKeluargaRawScript, idKeluarga)
	if err != nil {
		return entity.DataKeluargaFinal{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}

	dataKeluargaRaw := entity.DataKeluargaRaw{}
	if result.Next() {
		err := result.Scan(&dataKeluargaRaw.Id, &dataKeluargaRaw.Wilayah, &dataKeluargaRaw.Lingkungan, &dataKeluargaRaw.Nomor, &dataKeluargaRaw.KepalaKeluarga, &dataKeluargaRaw.Alamat, &dataKeluargaRaw.Status, &dataKeluargaRaw.NomorKKGereja)
		if err != nil {
			return entity.DataKeluargaFinal{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
	} else {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusNotFound, "Data Keluarga is not found")
	}
	result.Close()

	repositories := ctx.Locals("repositories").(Repositories)

	getLingkungan, err := repositories.DataLingkunganRepository.FindOneById(dataKeluargaRaw.Lingkungan, tx)
	if err != nil {
		return entity.DataKeluargaFinal{}, helper.CreateErrorMessage("Gagal untuk retrieve lingkungan data", err)
	}
	wilayah := getLingkungan.Wilayah
	lingkungan := entity.DataLingkungan{
		Id:             getLingkungan.Id,
		KodeLingkungan: getLingkungan.KodeLingkungan,
		NamaLingkungan: getLingkungan.NamaLingkungan,
		Wilayah:        wilayah,
	}

	getAnggotaRel, err := repositories.DataAnggotaKeluargaRelRepository.FindKeluargaAnggotaRel(dataKeluargaRaw.Id, db)
	if err != nil {
		return entity.DataKeluargaFinal{}, helper.CreateErrorMessage("Gagal untuk retrieve anggota relationship data", err)
	}

	var kepalaKeluarga entity.DataAnggotaWithStatus
	var anggota []entity.DataAnggotaWithStatus

	for _, anggotaRel := range getAnggotaRel {
		if anggotaRel.Hubungan == "Kepala Keluarga" {
			kepalaKeluarga = entity.DataAnggotaWithStatus{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
				Status:        anggotaRel.Status,
				JenisKelamin:  anggotaRel.JenisKelamin,
				NoTelp:        anggotaRel.NoTelp,
			}
		} else {
			anggota = append(anggota, entity.DataAnggotaWithStatus{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
				Status:        anggotaRel.Status,
				JenisKelamin:  anggotaRel.JenisKelamin,
				NoTelp:        anggotaRel.NoTelp,
			})
		}
	}

	dataKeluargaFinal := entity.DataKeluargaFinal{
		Id:             dataKeluargaRaw.Id,
		Wilayah:        wilayah,
		Lingkungan:     lingkungan,
		Nomor:          dataKeluargaRaw.Nomor,
		KepalaKeluarga: kepalaKeluarga,
		Alamat:         dataKeluargaRaw.Alamat,
		Anggota:        anggota,
		Status:         dataKeluargaRaw.Status,
		NomorKKGereja:  dataKeluargaRaw.NomorKKGereja,
	}

	return dataKeluargaFinal, nil
}

func (repository *dataKeluargaRepositoryImpl) CountKeluargaWithParam(ctx *fiber.Ctx, tx *sql.Tx, param string, id int32) (entity.TotalInt, error) {
	if param == "" {
		param = "lingkungan"
	}

	dataKeluargaRawScript := "SELECT COUNT(*) FROM data_keluarga WHERE id_lingkungan = ?"
	if param == "wilayah" {
		dataKeluargaRawScript = "SELECT COUNT(*) FROM data_keluarga WHERE id_wilayah = ?"
	}

	result, err := tx.Query(dataKeluargaRawScript, id)
	if err != nil {
		return entity.TotalInt{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	total := entity.TotalInt{}
	if result.Next() {
		err := result.Scan(&total.Total)
		if err != nil {
			return entity.TotalInt{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return total, nil
	} else {
		return entity.TotalInt{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}
}

func (repository *dataKeluargaRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) ([]entity.DataKeluargaFinal, error) {
	query := "SELECT id, id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, alamat, status, nomor_kk_gereja FROM data_keluarga"
	var args []interface{}
	var conditions []string

	// Mengambil query parameters
	idLingkunganStr := ctx.Query("idLingkungan")
	idWilayahStr := ctx.Query("idWilayah")
	idLingkunganParams := ctx.Params("idLingkungan")
	idWilayahParams := ctx.Params("idWilayah")

	// Filter berdasarkan path parameter idLingkungan
	if idLingkunganParams != "" {
		idLingkungan, err := strconv.Atoi(idLingkunganParams)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idLingkungan (path), it must be an integer")
		}
		conditions = append(conditions, "id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan path parameter idWilayah
	if idWilayahParams != "" {
		idWilayah, err := strconv.Atoi(idWilayahParams)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (path), it must be an integer")
		}
		conditions = append(conditions, "id_wilayah = ?")
		args = append(args, idWilayah)
	}

	// Filter berdasarkan query parameter idLingkungan
	if idLingkunganStr != "" {
		idLingkungan, err := strconv.Atoi(idLingkunganStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idLingkungan (query), it must be an integer")
		}
		conditions = append(conditions, "id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan query parameter idWilayah
	if idWilayahStr != "" {
		idWilayah, err := strconv.Atoi(idWilayahStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (query), it must be an integer")
		}
		conditions = append(conditions, "id_wilayah = ?")
		args = append(args, idWilayah)
	}
	// Jika ada kondisi, tambahkan ke query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	result, err := tx.Query(query, args...)
	if err != nil {
		return nil, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	var dataKeluargaList []entity.DataKeluargaFinal
	repositories := ctx.Locals("repositories").(Repositories)

	// Loop through all rows
	for result.Next() {
		dataKeluargaRaw := entity.DataKeluargaRaw{}
		err := result.Scan(&dataKeluargaRaw.Id, &dataKeluargaRaw.Wilayah, &dataKeluargaRaw.Lingkungan, &dataKeluargaRaw.Nomor, &dataKeluargaRaw.KepalaKeluarga, &dataKeluargaRaw.Alamat, &dataKeluargaRaw.Status, &dataKeluargaRaw.NomorKKGereja)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan resul", err)
		}

		// Get lingkungan data
		getLingkungan, err := repositories.DataLingkunganRepository.FindOneByIdSeparateTx(dataKeluargaRaw.Lingkungan, db)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk retrieve lingkungan data", err)
		}
		wilayah := getLingkungan.Wilayah
		lingkungan := entity.DataLingkungan{
			Id:             getLingkungan.Id,
			KodeLingkungan: getLingkungan.KodeLingkungan,
			NamaLingkungan: getLingkungan.NamaLingkungan,
			Wilayah:        wilayah,
		}

		// Get anggota relasi
		getAnggotaRel, err := repositories.DataAnggotaKeluargaRelRepository.FindKeluargaAnggotaRel(dataKeluargaRaw.Id, db)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk retrieve anggota relationship data", err)
		}

		var kepalaKeluarga entity.DataAnggotaWithStatus
		var anggota []entity.DataAnggotaWithStatus

		for _, anggotaRel := range getAnggotaRel {
			if anggotaRel.Hubungan == "Kepala Keluarga" {
				kepalaKeluarga = entity.DataAnggotaWithStatus{
					Id:            anggotaRel.IdAnggota,
					NamaLengkap:   anggotaRel.NamaLengkap,
					TanggalLahir:  anggotaRel.TanggalLahir,
					TanggalBaptis: anggotaRel.TanggalBaptis,
					Keterangan:    anggotaRel.Keterangan,
					Status:        anggotaRel.Status,
					JenisKelamin:  anggotaRel.JenisKelamin,
					NoTelp:        anggotaRel.NoTelp,
				}
			} else {
				anggota = append(anggota, entity.DataAnggotaWithStatus{
					Id:            anggotaRel.IdAnggota,
					NamaLengkap:   anggotaRel.NamaLengkap,
					TanggalLahir:  anggotaRel.TanggalLahir,
					TanggalBaptis: anggotaRel.TanggalBaptis,
					Keterangan:    anggotaRel.Keterangan,
					Status:        anggotaRel.Status,
					JenisKelamin:  anggotaRel.JenisKelamin,
					NoTelp:        anggotaRel.NoTelp,
				})
			}
		}

		// Populate final data structure
		dataKeluargaFinal := entity.DataKeluargaFinal{
			Id:             dataKeluargaRaw.Id,
			Wilayah:        wilayah,
			Lingkungan:     lingkungan,
			Nomor:          dataKeluargaRaw.Nomor,
			KepalaKeluarga: kepalaKeluarga,
			Alamat:         dataKeluargaRaw.Alamat,
			Anggota:        anggota,
			Status:         dataKeluargaRaw.Status,
			NomorKKGereja:  dataKeluargaRaw.NomorKKGereja,
		}

		// Add to list
		dataKeluargaList = append(dataKeluargaList, dataKeluargaFinal)
	}

	return dataKeluargaList, nil
}

func (repository *dataKeluargaRepositoryImpl) GetTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalKeluarga, error) {
	sqlScript := "SELECT COUNT(*) FROM data_keluarga WHERE status='aktif'"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return entity.TotalKeluarga{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	totalKeluarga := entity.TotalKeluarga{}
	if result.Next() {
		err := result.Scan(&totalKeluarga.Total)
		if err != nil {
			return entity.TotalKeluarga{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return totalKeluarga, nil
	} else {
		return entity.TotalKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}
}

func (repository *dataKeluargaRepositoryImpl) AddKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaRaw, error) {
	repositories := ctx.Locals("repositories").(Repositories)

	body := ctx.Body()
	request := new(helper.AddKeluargaRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataKeluargaRaw{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	//first add data_anggota first for kepala keluarga
	kepalaKeluarga, err := repositories.DataAnggotaRepository.AddAnggota(ctx, tx)
	if err != nil {
		return entity.DataKeluargaRaw{}, helper.CreateErrorMessage("Gagal memasukan data kepala keluarga", err)
	}

	//then add data_keluarga
	sqlScript := "INSERT INTO data_keluarga(id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, alamat, nomor_kk_gereja) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := tx.Exec(sqlScript, request.IdWilayah, request.IdLingkungan, request.Nomor, kepalaKeluarga.Id, request.Alamat, request.NomorKKGereja)
	if err != nil {
		return entity.DataKeluargaRaw{}, helper.CreateErrorMessage("Gagal memasukan data keluarga", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataKeluargaRaw{}, helper.CreateErrorMessage("Gagal mendapatkan ID data yang terakhir dimasukan", err)
	}

	//case misal abis add kepala keluarga -> di db ntar nilai relation data id_keluarga = 0 jadi perlu di update disini
	sqlUpdateScript := "UPDATE keluarga_anggota_rel set id_keluarga = ? where id_anggota = ?"
	_, err = tx.Exec(sqlUpdateScript, lastInsertId, kepalaKeluarga.Id)
	if err != nil {
		return entity.DataKeluargaRaw{}, helper.CreateErrorMessage("Gagal untuk update data relasi kepala keluarga", err)
	}

	// todo: note harusnya insert lingkungan dulu baru wilayah ya
	newDataKeluarga := entity.DataKeluargaRaw{
		Id:             int32(lastInsertId),
		Wilayah:        request.IdWilayah,
		Lingkungan:     request.IdLingkungan,
		Nomor:          request.Nomor,
		KepalaKeluarga: kepalaKeluarga.Id,
		Alamat:         request.Alamat,
		NomorKKGereja:  request.NomorKKGereja,
	}

	return newDataKeluarga, nil
}

func (repository *dataKeluargaRepositoryImpl) UpdateDataKeluarga(ctx *fiber.Ctx, tx *sql.Tx, db *sql.DB) (entity.UpdatedDataKeluarga, error) {
	sqlScript := "UPDATE data_keluarga SET"
	isKepalaKeluargaUpdated := false
	repositories := ctx.Locals("repositories").(Repositories)
	idKeluarga, errIdKeluarga := strconv.Atoi(ctx.Params("idKeluarga"))
	if errIdKeluarga != nil {
		return entity.UpdatedDataKeluarga{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id Keluarga, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.UpdateKeluargaRequest)
	request.Keterangan = "Kepala Keluarga" + request.Nomor
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.UpdatedDataKeluarga{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	var params []interface{}
	var setClauses []string

	if request.IdWilayah != 0 {
		setClauses = append(setClauses, "id_wilayah = ?")
		params = append(params, request.IdWilayah)
	}
	if request.IdLingkungan != 0 {
		setClauses = append(setClauses, "id_lingkungan = ?")
		params = append(params, request.IdLingkungan)
	}
	if request.Nomor != "" {
		setClauses = append(setClauses, "nomor = ?")
		params = append(params, request.Nomor)
	}
	if request.Alamat != "" {
		setClauses = append(setClauses, "alamat = ?")
		params = append(params, request.Alamat)
	}
	if request.Status != "" {
		setClauses = append(setClauses, "status = ?")
		params = append(params, request.Status)
	}
	if request.NomorKKGereja != "" {
		setClauses = append(setClauses, "nomor_kk_gereja = ?")
		params = append(params, request.NomorKKGereja)
	}
	if request.IdKepalaKeluarga != 0 {
		setClauses = append(setClauses, "id_kepala_keluarga = ?")
		params = append(params, request.IdKepalaKeluarga)
	}

	if len(setClauses) == 0 {
		return entity.UpdatedDataKeluarga{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idKeluarga)

	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		return entity.UpdatedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk update data keluarga", err)
	}

	//cek apakah ada update di idKepala keluarga dan apakah valuenya berbeda dengan yang old atau tidak
	if (request.IdKepalaKeluarga != request.OldIdKepalaKeluarga) && (request.IdKepalaKeluarga != 0) {
		isKepalaKeluargaUpdated = true
	}

	//update relasi kepala keluarga di db misalkan ada body request untuk update kepala keluarga
	// TO DO ada beberapa case yang masih kurang ->
	// case 1 ayah meniggal -> istri auto updated jadi kepala keluarga
	// case 2 ibu meninggal -> anak/anggota tertua auto updated jadi kepala keluarga
	// case 3 tidak ada yanng meninggal tapi data kepala keluarga terupdate -> old kepala keluarga statusnya jadi "anggota" (karena tidak tahu sebelumnya dia istri atau ayah atau bagaimana jadi di default ke "anggota" saja)
	if isKepalaKeluargaUpdated {
		_, err := repositories.DataAnggotaRepository.UpdateKeteranganAnggota(ctx, tx)
		if err != nil {
			return entity.UpdatedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk update keterangan relasi dan hubungan", err)
		}
	}
	// sepertinya perlu di commit dulu disini
	errorCommit := tx.Commit()
	if errorCommit != nil {
		return entity.UpdatedDataKeluarga{}, helper.CreateErrorMessage("Error when commiting transaction", errorCommit)
	}

	getAnggotaRel, err := repositories.DataAnggotaKeluargaRelRepository.FindKeluargaAnggotaRel(int32(idKeluarga), db)
	if err != nil {
		return entity.UpdatedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk retrieve anggota relationship data", err)
	}

	var kepalaKeluarga entity.DataAnggotaWithStatus
	var anggota []entity.DataAnggotaWithStatus

	for _, anggotaRel := range getAnggotaRel {
		if anggotaRel.Hubungan == "Kepala Keluarga" {
			kepalaKeluarga = entity.DataAnggotaWithStatus{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
				Status:        anggotaRel.Status,
				JenisKelamin:  anggotaRel.JenisKelamin,
				NoTelp:        anggotaRel.NoTelp,
			}
		} else {
			anggota = append(anggota, entity.DataAnggotaWithStatus{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
				Status:        anggotaRel.Status,
				JenisKelamin:  anggotaRel.JenisKelamin,
				NoTelp:        anggotaRel.NoTelp,
			})
		}
	}

	newDataKeluarga := entity.UpdatedDataKeluarga{
		Id:             int32(idKeluarga),
		IdWilayah:      request.IdWilayah,
		IdLingkungan:   request.IdLingkungan,
		Nomor:          request.Nomor,
		KepalaKeluarga: kepalaKeluarga,
		Alamat:         request.Alamat,
		Status:         request.Status,
		Anggota:        anggota,
		NomorKKGereja:  request.NomorKKGereja,
	}

	return newDataKeluarga, nil
}

func (repository *dataKeluargaRepositoryImpl) DeleteDataKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DeletedDataKeluarga, error) {
	idKeluarga, errIdKeluarga := strconv.Atoi(ctx.Params("idKeluarga"))
	if errIdKeluarga != nil {
		return entity.DeletedDataKeluarga{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id Keluarga, it must be an integer")
	}

	// Step 1: Ambil ID anggota dari tabel keluarga_anggota_rel
	var deletedAnggotaIds []int32
	rows, err := tx.Query("SELECT id_anggota FROM keluarga_anggota_rel WHERE id_keluarga = ?", idKeluarga)
	if err != nil {
		return entity.DeletedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk fetch anggota before deletion", err)
	}
	defer rows.Close()

	for rows.Next() {
		var idAnggota int32
		if err := rows.Scan(&idAnggota); err != nil {
			return entity.DeletedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk scan anggota data", err)
		}
		deletedAnggotaIds = append(deletedAnggotaIds, idAnggota)
	}

	// Step 2: Hapus data relasi dari tabel keluarga_anggota_rel
	sqlScript := "DELETE keluarga_anggota_rel WHERE id_keluarga = ?"
	_, err = tx.Exec(sqlScript, idKeluarga)
	if err != nil {
		return entity.DeletedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk delete data keluarga anggota rel", err)
	}

	// Step 3: Hapus data keluarga dari tabel data_keluarga
	sqlScript = "DELETE data_keluarga WHERE id = ?"
	_, err = tx.Exec(sqlScript, idKeluarga)
	if err != nil {
		return entity.DeletedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk delete data keluarga", err)
	}

	// Step 4: Hapus data dari tabel data_anggota berdasarkan deletedAnggotaIds
	if len(deletedAnggotaIds) > 0 {
		query := "DELETE FROM data_anggota WHERE id IN (?" + strings.Repeat(", ?", len(deletedAnggotaIds)-1) + ")"
		args := make([]interface{}, len(deletedAnggotaIds))
		for i, id := range deletedAnggotaIds {
			args[i] = id
		}

		_, err = tx.Exec(query, args...)
		if err != nil {
			return entity.DeletedDataKeluarga{}, helper.CreateErrorMessage("Gagal untuk delete data anggota", err)
		}
	}

	res := entity.DeletedDataKeluarga{
		Id:                int32(idKeluarga),
		DeletedAnggotaIds: deletedAnggotaIds,
	}

	return res, nil
}
