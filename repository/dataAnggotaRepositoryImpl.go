package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gkru-service/entity"
	"gkru-service/helper"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type dataAnggotaRepositoryImpl struct {
}

func NewDataAnggotaRepository(db *sql.DB) DataAnggotaRepository {
	return &dataAnggotaRepositoryImpl{}
}

func (repository *dataAnggotaRepositoryImpl) GetTotalAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalAnggota, error) {
	sqlScript := "SELECT COUNT(*) FROM data_anggota where status='HIDUP'"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return entity.TotalAnggota{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	totalAnggota := entity.TotalAnggota{}
	if result.Next() {
		err := result.Scan(&totalAnggota.Total)
		if err != nil {
			return entity.TotalAnggota{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return totalAnggota, nil
	} else {
		return entity.TotalAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}
}

func (repository *dataAnggotaRepositoryImpl) AddAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggota, error) {
	sqlScript := "INSERT INTO data_anggota(nama_lengkap, tanggal_lahir, tanggal_baptis, keterangan, status, jenis_kelamin) VALUES(?, ?, ?, ?, ?, ?)"
	body := ctx.Body()
	request := new(helper.AddAnggotaRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataAnggota{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// to do kalo misal request.Keterangan = Kepala keluarga harusnya ada yang nambahin ke data Keluarga juga (?) -> ditambahin di view list keluarga aja ntar
	result, err := tx.Exec(sqlScript, request.NamaLengkap, request.TanggalLahir, request.TanggalBaptis, request.Keterangan, request.Status, request.JenisKelamin)
	if err != nil {
		return entity.DataAnggota{}, helper.CreateErrorMessage("Gagal memasukan data anggota", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataAnggota{}, helper.CreateErrorMessage("Gagal mendapatkan ID data yang terakhir dimasukan", err)
	}

	sqlScriptRelData := "INSERT INTO keluarga_anggota_rel (id_keluarga, id_anggota, hubungan) VALUES(?, ?, ?)"
	_, err = tx.Exec(sqlScriptRelData, request.IdKeluarga, lastInsertId, request.Hubungan)
	if err != nil {
		return entity.DataAnggota{}, helper.CreateErrorMessage("Gagal memasukan keluarga anggota relation data", err)
	}

	newDataAnggota := entity.DataAnggota{
		Id:          int32(lastInsertId),
		NamaLengkap: request.NamaLengkap,
		// TanggalLahir:  request.TanggalLahir,
		// TanggalBaptis: request.TanggalBaptis,
		Keterangan: request.Keterangan,
	}

	return newDataAnggota, nil
}

func (repository *dataAnggotaRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaComplete, error) {
	query := `SELECT a.id, a.nama_lengkap, a.tanggal_lahir, a.tanggal_baptis, a.keterangan, a.status, a.jenis_kelamin, b.id_keluarga, b.hubungan, c.id_wilayah, c.id_lingkungan, d.kode_lingkungan, d.nama_lingkungan, e.kode_wilayah, e.nama_wilayah FROM data_anggota a 
	JOIN keluarga_anggota_rel b ON a.id = b.id_anggota 
	JOIN data_keluarga c ON b.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id
	WHERE a.id = ?`
	idAnggota := ctx.Params("idAnggota")

	result, err := tx.Query(query, idAnggota)
	if err != nil {
		return entity.DataAnggotaComplete{}, helper.CreateErrorMessage("helper.CreateErrorMessage", err)
	}
	defer result.Close()

	dataAnggotaComplete := entity.DataAnggotaComplete{}
	if result.Next() {
		err := result.Scan(&dataAnggotaComplete.Id, &dataAnggotaComplete.NamaLengkap, &dataAnggotaComplete.TanggalLahir, &dataAnggotaComplete.TanggalBaptis, &dataAnggotaComplete.Keterangan, &dataAnggotaComplete.Status, &dataAnggotaComplete.JenisKelamin, &dataAnggotaComplete.IdKeluarga, &dataAnggotaComplete.Hubungan, &dataAnggotaComplete.IdWilayah, &dataAnggotaComplete.IdLingkungan, &dataAnggotaComplete.KodeLingkungan, &dataAnggotaComplete.NamaLingkungan, &dataAnggotaComplete.KodeWilayah, &dataAnggotaComplete.NamaWilayah)
		if err != nil {
			return entity.DataAnggotaComplete{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
	} else {
		return entity.DataAnggotaComplete{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}

	return dataAnggotaComplete, nil
}

func (repository *dataAnggotaRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataAnggotaComplete, error) {
	// note -> id lingkungan & wilayah kemungkinan untuk filter di data table nanti
	query := `SELECT a.id, a.nama_lengkap, a.tanggal_lahir, a.tanggal_baptis, a.keterangan, a.status, a.jenis_kelamin, b.id_keluarga, b.hubungan, c.id_wilayah, c.id_lingkungan, d.kode_lingkungan, d.nama_lingkungan, e.kode_wilayah, e.nama_wilayah FROM data_anggota a 
	JOIN keluarga_anggota_rel b ON a.id = b.id_anggota 
	JOIN data_keluarga c ON b.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id`
	var args []interface{}
	var conditions []string

	// Mengambil query parameters -> ini basic di semua find all karena bakal ada user berdasarkan wilayah/lingkungan
	// tapi untuk filter lebih lengkapnya pake data table jadi yang di handle ini aja
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
		conditions = append(conditions, "c.id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan path parameter idWilayah
	if idWilayahParams != "" {
		idWilayah, err := strconv.Atoi(idWilayahParams)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (path), it must be an integer")
		}
		conditions = append(conditions, "c.id_wilayah = ?")
		args = append(args, idWilayah)
	}

	// Filter berdasarkan query parameter idLingkungan
	if idLingkunganStr != "" {
		idLingkungan, err := strconv.Atoi(idLingkunganStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idLingkungan (query), it must be an integer")
		}
		conditions = append(conditions, "c.id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan query parameter idWilayah
	if idWilayahStr != "" {
		idWilayah, err := strconv.Atoi(idWilayahStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (query), it must be an integer")
		}
		conditions = append(conditions, "c. id_wilayah = ?")
		args = append(args, idWilayah)
	}
	// Jika ada kondisi, tambahkan ke query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	result, err := tx.Query(query, args...)
	if err != nil {
		errorMessage := fmt.Sprintf("Gagal mengeksekusi query: %v", err.Error())
		return nil, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
	}
	defer result.Close()

	var dataAnggotaList []entity.DataAnggotaComplete

	for result.Next() {
		dataAnggotaComplete := entity.DataAnggotaComplete{}
		err := result.Scan(&dataAnggotaComplete.Id, &dataAnggotaComplete.NamaLengkap, &dataAnggotaComplete.TanggalLahir, &dataAnggotaComplete.TanggalBaptis, &dataAnggotaComplete.Keterangan, &dataAnggotaComplete.Status, &dataAnggotaComplete.JenisKelamin, &dataAnggotaComplete.IdKeluarga, &dataAnggotaComplete.Hubungan, &dataAnggotaComplete.IdWilayah, &dataAnggotaComplete.IdLingkungan, &dataAnggotaComplete.KodeLingkungan, &dataAnggotaComplete.NamaLingkungan, &dataAnggotaComplete.KodeWilayah, &dataAnggotaComplete.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Add to list
		dataAnggotaList = append(dataAnggotaList, dataAnggotaComplete)
	}

	// If no rows were found, return an empty list
	// no need to throw error when not found
	// if len(dataAnggotaList) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "No Data Anggota found")
	// }

	return dataAnggotaList, nil
}

func (repository *dataAnggotaRepositoryImpl) FindAllWithIdKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataAnggotaComplete, error) {
	// note -> id lingkungan & wilayah kemungkinan untuk filter di data table nanti
	query := `SELECT a.id, a.nama_lengkap, a.tanggal_lahir, a.tanggal_baptis, a.keterangan, a.status, a.jenis_kelamin, b.id_keluarga, b.hubungan, c.id_wilayah, c.id_lingkungan, d.kode_lingkungan, d.nama_lingkungan, e.kode_wilayah, e.nama_wilayah FROM data_anggota a 
	JOIN keluarga_anggota_rel b ON a.id = b.id_anggota 
	JOIN data_keluarga c ON b.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id`
	var args []interface{}
	var conditions []string

	// Mengambil query parameters -> ini basic di semua find all karena bakal ada user berdasarkan wilayah/lingkungan
	// tapi untuk filter lebih lengkapnya pake data table jadi yang di handle ini aja
	idLingkunganStr := ctx.Query("idLingkungan")
	idWilayahStr := ctx.Query("idWilayah")
	idKeluargaStr := ctx.Query("idKeluarga")
	idLingkunganParams := ctx.Params("idLingkungan")
	idWilayahParams := ctx.Params("idWilayah")

	// Filter berdasarkan path parameter idLingkungan
	if idLingkunganParams != "" {
		idLingkungan, err := strconv.Atoi(idLingkunganParams)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idLingkungan (path), it must be an integer")
		}
		conditions = append(conditions, "c.id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan path parameter idWilayah
	if idWilayahParams != "" {
		idWilayah, err := strconv.Atoi(idWilayahParams)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (path), it must be an integer")
		}
		conditions = append(conditions, "c.id_wilayah = ?")
		args = append(args, idWilayah)
	}

	// Filter berdasarkan query parameter idLingkungan
	if idLingkunganStr != "" {
		idLingkungan, err := strconv.Atoi(idLingkunganStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idLingkungan (query), it must be an integer")
		}
		conditions = append(conditions, "c.id_lingkungan = ?")
		args = append(args, idLingkungan)
	}

	// Filter berdasarkan query parameter idWilayah
	if idWilayahStr != "" {
		idWilayah, err := strconv.Atoi(idWilayahStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idWilayah (query), it must be an integer")
		}
		conditions = append(conditions, "c.id_wilayah = ?")
		args = append(args, idWilayah)
	}

	// Filter berdasarkan query parameter idKeluarga
	if idKeluargaStr != "" {
		idKeluarga, err := strconv.Atoi(idKeluargaStr)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid idKeluarga (query), it must be an integer")
		}
		conditions = append(conditions, "c.id = ?")
		args = append(args, idKeluarga)
	}
	// Jika ada kondisi, tambahkan ke query
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	result, err := tx.Query(query, args...)
	if err != nil {
		errorMessage := fmt.Sprintf("Gagal mengeksekusi query: %v", err.Error())
		return nil, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
	}
	defer result.Close()

	var dataAnggotaList []entity.DataAnggotaComplete

	for result.Next() {
		dataAnggotaComplete := entity.DataAnggotaComplete{}
		err := result.Scan(&dataAnggotaComplete.Id, &dataAnggotaComplete.NamaLengkap, &dataAnggotaComplete.TanggalLahir, &dataAnggotaComplete.TanggalBaptis, &dataAnggotaComplete.Keterangan, &dataAnggotaComplete.Status, &dataAnggotaComplete.JenisKelamin, &dataAnggotaComplete.IdKeluarga, &dataAnggotaComplete.Hubungan, &dataAnggotaComplete.IdWilayah, &dataAnggotaComplete.IdLingkungan, &dataAnggotaComplete.KodeLingkungan, &dataAnggotaComplete.NamaLingkungan, &dataAnggotaComplete.KodeWilayah, &dataAnggotaComplete.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Add to list
		dataAnggotaList = append(dataAnggotaList, dataAnggotaComplete)
	}

	// If no rows were found, return an empty list
	// if len(dataAnggotaList) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "No Data Anggota found")
	// }

	return dataAnggotaList, nil
}

// method akan dipanggil jika current kepala keluarga status nya di edit dari hidup ke mati
func (repository *dataAnggotaRepositoryImpl) UpdateKepalaKeluarga(ctx *fiber.Ctx, tx *sql.Tx, idKeluarga int32, idAnggota *int32) error {
	getIstriScript := `SELECT a.id FROM data_anggota a 
	JOIN keluarga_anggota_rel b
	ON a.id = b.id_anggota
	WHERE b.id_keluarga = ? 
	AND a.keterangan LIKE ?`
	getOldestAnggotaScript := `SELECT a.id FROM data_anggota a 
	JOIN keluarga_anggota_rel b
	ON a.id = b.id_anggota
	WHERE b.id_keluarga = ?
	AND a.keterangan NOT LIKE "%Kepala Keluarga%"
	AND a.status = 'HIDUP'
	ORDER BY tanggal_lahir 
	ASC LIMIT 1`
	updateDataAnggotasqlScript := `
		UPDATE data_anggota 
		SET keterangan = 'Kepala Keluarga' 
		WHERE id = ?
	`
	udpateDataRelationScript := "UPDATE keluarga_anggota_rel SET hubungan = 'Kepala Keluarga' WHERE id_anggota = ?"
	likeCondition := "%Istri%"

	result, err := tx.Query(getIstriScript, idKeluarga, likeCondition)
	if err != nil {
		return helper.CreateErrorMessage("Gagal untuk execute select query for istri", err)
	}
	idAnggotaResult := entity.IdDataAnggota{}
	dataFound := false
	for result.Next() {
		err := result.Scan(&idAnggotaResult.Id)
		if err != nil {
			return helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		dataFound = true
	}
	result.Close()

	// Jika tidak ada anggota dengan keterangan "Istri", pilih anggota dengan tanggal lahir tertua
	if !dataFound {
		oldestResult, err := tx.Query(getOldestAnggotaScript, idKeluarga)
		if err != nil {
			return helper.CreateErrorMessage("Gagal untuk execute select query untuk anggota tertua", err)
		}

		if oldestResult.Next() {
			err := oldestResult.Scan(&idAnggotaResult.Id)
			if err != nil {
				return helper.CreateErrorMessage("Gagal untuk scan result untuk anggota tertua", err)
			}
		} else {
			// kayanya sini perlu dibuat bila tidak ada yang eligible otomatis keluarganya di non aktifkan ? atau sementara biarkan dulu
			return fiber.NewError(fiber.StatusNotFound, "Tidak ada anggota yang bisa menjadi Kepala Keluarga lagi")
		}
		oldestResult.Close()
	}

	//update data kepala keluarga yang baru
	_, err = tx.Exec(updateDataAnggotasqlScript, idAnggotaResult.Id)
	if err != nil {
		return helper.CreateErrorMessage("Gagal untuk update Kepala Keluarga", err)
	}

	//update relation kepala keluarga yang baru 
	_, err = tx.Exec(udpateDataRelationScript, idAnggotaResult.Id)
	if err != nil {
		return helper.CreateErrorMessage("Gagal untuk update data keluarga_anggota_rel", err)
	}

	//update data_keluarga ke id kepala keluarga yang baru 
	_, err = tx.Exec("UPDATE data_keluarga SET id_kepala_keluarga = ? WHERE id = ?", idAnggotaResult.Id, idKeluarga)
	if err != nil {
		return helper.CreateErrorMessage("Gagal untuk update id kepala keluarga di data keluarga", err)
	}

	// jika dari update bukan delete -> misal update status dari hidup -> meniggal, nanti keterangan di data anggota akan menjadi "anggota"
	if idAnggota != nil {
		fmt.Println("masuk sini")
		_, err = tx.Exec("UPDATE data_anggota SET keterangan = 'Anggota' WHERE id = ?", idAnggota)
		if err != nil {
			return helper.CreateErrorMessage("Gagal untuk update keterangan di data anggota", err)
		}

		_, err = tx.Exec("UPDATE keluarga_anggota_rel SET hubungan = 'Anggota' WHERE id_anggota = ?", idAnggota)
		if err != nil {
			return helper.CreateErrorMessage("Gagal untuk update keterangan di relasi", err)
		}
	}

	return nil
}

// kalo mau update hubungan jadi kepala keluarga tidak bisa ya harus dari edit KK -> mending di fe dibuat pilihan aja keterangan dan status itu
func (repository *dataAnggotaRepositoryImpl) UpdateAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithStatus, error) {
	sqlScript := "UPDATE data_anggota SET"
	idAnggota, err := strconv.Atoi(ctx.Params("idAnggota"))
	if err != nil {
		return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusBadRequest, "Invalid idAnggota, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.UpdateAnggotaRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	var params []interface{}
	var setClauses []string

	// Dynamically building the SET clause
	if request.NamaLengkap != "" {
		setClauses = append(setClauses, "nama_lengkap = ?")
		params = append(params, request.NamaLengkap)
	}
	if !request.TanggalLahir.IsZero() {
		setClauses = append(setClauses, "tanggal_lahir = ?")
		params = append(params, request.TanggalLahir.ToTime())
	}
	if !request.TanggalBaptis.IsZero() {
		setClauses = append(setClauses, "tanggal_baptis = ?")
		params = append(params, request.TanggalBaptis.ToTime())
	}
	if request.Keterangan != "" { //apakah keterangan seperti istri, anak, dll bisa diupdate lewat sini ya nanti ? atau cuman dari update kelluarga saja misalkan ada perubahan data kepala keluarga ???? KEMUNGKINAN BESAR TIDAK PERLU
		setClauses = append(setClauses, "keterangan = ?")
		params = append(params, request.Keterangan)
	}
	if request.Status != "" {
		setClauses = append(setClauses, "status = ?")
		params = append(params, request.Status)
	}
	if request.JenisKelamin != "" {
		setClauses = append(setClauses, "jenis_kelamin = ?")
		params = append(params, request.JenisKelamin)
	}

	// Check if there are fields to update
	if len(setClauses) == 0 {
		return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	// Joining all set clauses
	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idAnggota)

	// Executing the update statement
	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		return entity.DataAnggotaWithStatus{}, helper.CreateErrorMessage("Gagal untuk update data anggota", err)
	}

	// kalau request.keterangan ada -> [PENTING] request.hubungan dibuat keisi otomatis ya nanti di FE misalkan awalnya istri diubah ke kepala keluarga (?) -> hubungan diubah jadi kepala kelaurga keterangan di data anggota tetap istri
	// bisa jadi ini tidak perlu karena kemungkinan keterangan itu tidak bisa diupdate lewat UI ya -> dari flow bisanya kaya misal ada kepela keluarga yang meninggal dll
	if request.Hubungan != "" {
		sqlScriptRelData := "UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?"
		_, err = tx.Exec(sqlScriptRelData, request.Hubungan, idAnggota)
		if err != nil {
			return entity.DataAnggotaWithStatus{}, helper.CreateErrorMessage("Gagal untuk update data keluarga_anggota_rel", err)
		}
	}

	if request.Status == "MENINGGAL" && request.IsKepalaKeluarga { 
		errUpdateKepalaKeluarga := repository.UpdateKepalaKeluarga(ctx, tx, request.IdKeluarga, &request.Id)
		if errUpdateKepalaKeluarga != nil {
			return entity.DataAnggotaWithStatus{}, errUpdateKepalaKeluarga
		}
	}

	newDataAnggota := entity.DataAnggotaWithStatus{
		Id:            int32(idAnggota),
		NamaLengkap:   request.NamaLengkap,
		TanggalLahir:  request.TanggalLahir,
		TanggalBaptis: request.TanggalBaptis,
		Keterangan:    request.Keterangan,
		Status:        request.Status,
		JenisKelamin:  request.JenisKelamin,
	}

	return newDataAnggota, nil
}

func (repository *dataAnggotaRepositoryImpl) UpdateKeteranganAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithKeteranganOnly, error) {
	sqlScript := "UPDATE data_anggota SET keterangan = ? where id = ?"
	sqlScriptRelData := "UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?"
	body := ctx.Body()
	request := new(helper.UpdateKeteranganAnggotaRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Update data anggota
	_, err := tx.Exec(sqlScript, "Kepala Keluarga", request.Id)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, helper.CreateErrorMessage("Gagal untuk update data keplaa keluarga di data anggota", err)
	}
	_, err = tx.Exec(sqlScriptRelData, "Kepala Keluarga", request.Id)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, helper.CreateErrorMessage("Gagal untuk update data kepala keluarga lama di  keluarga_anggota_rel", err)
	}

	// update keterangan old kepala keluarga
	_, err = tx.Exec("UPDATE data_anggota SET keterangan = ? where id = ?", "Anggota", request.OldId)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, helper.CreateErrorMessage("Gagal untuk update data keplaa keluarga lama di data anggota", err)
	}
	_, err = tx.Exec("UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?", "Anggota", request.OldId)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, helper.CreateErrorMessage("Gagal untuk update data kepala keluarga lama di keluarga_anggota_rel", err)
	}

	newDataAnggota := entity.DataAnggotaWithKeteranganOnly{
		Id:         request.Id,
		Keterangan: request.Keterangan,
	}

	return newDataAnggota, nil
}

func (repository *dataAnggotaRepositoryImpl) DeleteOneAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdDataAnggota, error) {
	idAnggota, errIdAnggota := strconv.Atoi(ctx.Params("idAnggota"))
	if errIdAnggota != nil {
		return entity.IdDataAnggota{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id anggota, it must be an integer")
	}

	// Step 1: Ambil ID anggota dari tabel keluarga_anggota_rel dan delete
	dataAnggota, errDataAnggota := repository.FindOne(ctx, tx)
	if errDataAnggota != nil {
		return entity.IdDataAnggota{}, errDataAnggota
	}

	sqlScript := "DELETE FROM keluarga_anggota_rel WHERE id_anggota = ?"
	_, err := tx.Exec(sqlScript, idAnggota)
	if err != nil {
		return entity.IdDataAnggota{}, helper.CreateErrorMessage("Gagal untuk delete data keluarga anggota rel", err)
	}

	//step 2: misalkan yang di delete adalah kepala keluarga -> maka istri langsung jadi kepala keluarga baru atau anak yang paling tua
	if dataAnggota.Hubungan == "Kepala Keluarga" { //pakai Hubungan dari data anngota keluarga rel karena bisa jadi anggota dengan status istri atau anak tertua adalah kepala keluarga
		errUpdateKepalaKeluarga := repository.UpdateKepalaKeluarga(ctx, tx, dataAnggota.IdKeluarga, nil)
		if errUpdateKepalaKeluarga != nil {
			return entity.IdDataAnggota{}, errUpdateKepalaKeluarga
		}
	}
	
	// step 3 baru delete data anggota setelah semua kunci tamu diperbarui
	sqlScript = "DELETE FROM data_anggota WHERE id = ?"
	_, err = tx.Exec(sqlScript, idAnggota)
	if err != nil {
		return entity.IdDataAnggota{}, helper.CreateErrorMessage("Gagal untuk delete data anggota", err)
	}

	res := entity.IdDataAnggota{
		Id: int32(idAnggota),
	}

	return res, nil
}

func (repository *dataAnggotaRepositoryImpl) DeleteBulkAnggota(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.IdDataAnggota, error) {
	// paka body aja supaya tidak terlalu banyak proses db nanti sepert iselect dll, jadi di FE misalkan nilai data table di select, dapat ditambahkan ke redux sebagai request body nantinya
	body := ctx.Body()
	request := new(helper.DeleteAnggotaRequest)
	errMarshall := json.Unmarshal(body, request)
	if errMarshall != nil {
		return []entity.IdDataAnggota{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Extract IDs to delete
	var idsToDelete []int32
	for _, anggota := range request.SelectedAnggota {
		idsToDelete = append(idsToDelete, anggota.Id)
	}

	// Prepare placeholders for SQL IN clause
	placeholders := make([]string, len(idsToDelete))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	placeholderString := strings.Join(placeholders, ",")

	// Delete related records from keluarga_anggota_rel
	sqlScript := fmt.Sprintf("DELETE FROM keluarga_anggota_rel WHERE id_anggota IN (%s)", placeholderString)
	_, err := tx.Exec(sqlScript, helper.ConvertToInterfaceSlice(idsToDelete)...)
	if err != nil {
		return []entity.IdDataAnggota{}, helper.CreateErrorMessage("Gagal untuk delete data keluarga anggota rel", err)
	}

	// Delete records from data_anggota
	sqlScript = fmt.Sprintf("DELETE FROM data_anggota WHERE id IN (%s)", placeholderString)
	_, err = tx.Exec(sqlScript, helper.ConvertToInterfaceSlice(idsToDelete)...)
	if err != nil {
		return []entity.IdDataAnggota{}, helper.CreateErrorMessage("Gagal untuk delete data anggota", err)
	}

	// Check if any of the deleted members was the kepala keluarga and update if needed
	for _, anggota := range request.SelectedAnggota {
		if anggota.Hubungan == "Kepala Keluarga" {
			errUpdateKepalaKeluarga := repository.UpdateKepalaKeluarga(ctx, tx, anggota.IdKeluarga, nil)
			if errUpdateKepalaKeluarga != nil {
				return []entity.IdDataAnggota{}, errUpdateKepalaKeluarga
			}
			break
		}
	}

	// Return deleted IDs
	var deletedIds []entity.IdDataAnggota
	for _, id := range idsToDelete {
		deletedIds = append(deletedIds, entity.IdDataAnggota{Id: id})
	}

	return deletedIds, nil
}
