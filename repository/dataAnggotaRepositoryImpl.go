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
		return entity.TotalAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to execute query")
	}
	defer result.Close()

	totalAnggota := entity.TotalAnggota{}
	if result.Next() {
		err := result.Scan(&totalAnggota.Total)
		if err != nil {
			return entity.TotalAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}
		return totalAnggota, nil
	} else {
		return entity.TotalAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "No data found")
	}
}

func (repository *dataAnggotaRepositoryImpl) AddAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggota, error) {

	fmt.Println("masuk AddAnggota")

	sqlScript := "INSERT INTO data_anggota(nama_lengkap, tanggal_lahir, tanggal_baptis, keterangan, status) VALUES(?, ?, ?, ?, ?)"
	body := ctx.Body()
	fmt.Println(body)
	request := new(helper.AddAnggotaRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		fmt.Println(err)
		return entity.DataAnggota{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// to do kalo misal request.Keterangan = Kepala keluarga harusnya ada yang nambahin ke data Keluarga juga (?) -> ditambahin di view list keluarga aja ntar
	result, err := tx.Exec(sqlScript, request.NamaLengkap, request.TanggalLahir, request.TanggalBabtis, request.Keterangan, request.Status)
	if err != nil {
		fmt.Println(err)
		return entity.DataAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to insert data anggota")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve last insert ID")
	}

	sqlScriptRelData := "INSERT INTO keluarga_anggota_rel (id_keluarga, id_anggota, hubungan) VALUES(?, ?, ?)"
	_, err = tx.Exec(sqlScriptRelData, request.IdKeluarga, lastInsertId, request.Hubungan)
	if err != nil {
		return entity.DataAnggota{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to insert keluarga anggota relation data")
	}

	newDataAnggota := entity.DataAnggota{
		Id:            int32(lastInsertId),
		NamaLengkap:   request.NamaLengkap,
		TanggalLahir:  request.TanggalLahir,
		TanggalBaptis: request.TanggalBabtis,
		Keterangan:    request.Keterangan,
	}

	return newDataAnggota, nil
}

func (repository *dataAnggotaRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaComplete, error) {
	query := `SELECT a.id, a.nama_lengkap, a.tanggal_lahir, a.tanggal_babtis, a.keterangan, a.status, b.id_keluarga, c.id_wilayah, c.id_lingkungan, d.kode_lingkungan, d.nama_lingkungan, e.kode_wilayah, e.nama_wilayah FROM data_anggota a 
	JOIN keluarga_anggota_rel b ON a.id = b.id_anggota 
	JOIN data_keluarga c ON b.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id
	WHERE a.id = ?`
	idAnggota := ctx.Params("idAnggota")

	result, err := tx.Query(query, idAnggota)
	if err != nil {
		return entity.DataAnggotaComplete{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to execute query")
	}
	defer result.Close()

	dataAnggotaComplete := entity.DataAnggotaComplete{}
	for result.Next() {
		err := result.Scan(&dataAnggotaComplete.Id, &dataAnggotaComplete.NamaLengkap, &dataAnggotaComplete.TanggalLahir, &dataAnggotaComplete.TanggalBaptis, &dataAnggotaComplete.Keterangan, &dataAnggotaComplete.Status, &dataAnggotaComplete.IdKeluarga,  &dataAnggotaComplete.IdWilayah,  &dataAnggotaComplete.IdLingkungan,  &dataAnggotaComplete.KodeLingkungan,  &dataAnggotaComplete.NamaLingkungan,  &dataAnggotaComplete.KodeWilayah,  &dataAnggotaComplete.NamaWilayah)
		if err != nil {
			return entity.DataAnggotaComplete{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}
	}

	return dataAnggotaComplete, nil
}

func (repository *dataAnggotaRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataAnggotaComplete, error) {
	// note -> id lingkungan & wilayah kemungkinan untuk filter di data table nanti
	query := `SELECT a.id, a.nama_lengkap, a.tanggal_lahir, a.tanggal_babtis, a.keterangan, a.status, b.id_keluarga, c.id_wilayah, c.id_lingkungan, d.kode_lingkungan, d.nama_lingkungan, e.kode_wilayah, e.nama_wilayah FROM data_anggota a 
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
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to execute query")
	}
	defer result.Close()

	var dataAnggotaList []entity.DataAnggotaComplete

	for result.Next() {
		dataAnggotaComplete := entity.DataAnggotaComplete{}
		err := result.Scan(&dataAnggotaComplete.Id, &dataAnggotaComplete.NamaLengkap, &dataAnggotaComplete.TanggalLahir, &dataAnggotaComplete.TanggalBaptis, &dataAnggotaComplete.Keterangan, &dataAnggotaComplete.Status, &dataAnggotaComplete.IdKeluarga,  &dataAnggotaComplete.IdWilayah,  &dataAnggotaComplete.IdLingkungan,  &dataAnggotaComplete.KodeLingkungan,  &dataAnggotaComplete.NamaLingkungan,  &dataAnggotaComplete.KodeWilayah,  &dataAnggotaComplete.NamaWilayah)
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}

		// Add to list
		dataAnggotaList = append(dataAnggotaList, dataAnggotaComplete)
	}

	// If no rows were found, return an empty list
	if len(dataAnggotaList) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "No Data Anggota found")
	}

	return dataAnggotaList, nil
}

// method akan dipanggil jika current kepala keluarga status nya di edit dari hidup ke mati
func (repository *dataAnggotaRepositoryImpl) UpdateKepalaKeluarga(ctx *fiber.Ctx, tx *sql.Tx, idAnggota int32, idKeluarga int32) error {
	getIstriScript := "SELECT id FROM data_anggota WHERE id_keluarga = ? AND keterangan LIKE ?"
	updateDataAnggotasqlScript := `
		UPDATE data_anggota 
		SET keterangan = 'Kepala Keluarga' 
		WHERE id = ?
	`
	likeCondition := "%Istri%"

	result, err := tx.Query(getIstriScript, idKeluarga, likeCondition)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to execute select query")
	}
	defer result.Close()
	idAnggotaResult := entity.IdDataAnggota{}
	for result.Next() {
		err := result.Scan(&idAnggotaResult.Id)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}
	}

	_, err = tx.Exec(updateDataAnggotasqlScript, idAnggotaResult.Id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update Kepala Keluarga")
	}

	return nil
}

// kalo mau update hubungan jadi kepala keluarga tidak bisa ya harus dari edit KK -> mending di fe dibuat pilihan aja keternagan dan status itu
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
		fmt.Println(marshalError)
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
		params = append(params, request.TanggalLahir)
	}
	if !request.TanggalBabtis.IsZero() {
		setClauses = append(setClauses, "tanggal_baptis = ?")
		params = append(params, request.TanggalBabtis)
	}
	if request.Keterangan != "" {
		setClauses = append(setClauses, "keterangan = ?")
		params = append(params, request.Keterangan)
	}
	if request.Status != "" {
		setClauses = append(setClauses, "status = ?")
		params = append(params, request.Status)
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
		fmt.Println(err)
		return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data anggota")
	}

	// kalau request.keterangan ada -> request.hubungan dibuat keisi otomatis ya nanti di FE
	if request.Hubungan != "" {
		sqlScriptRelData := "UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?"
		_, err = tx.Exec(sqlScriptRelData, request.Hubungan, idAnggota)
		if err != nil {
			return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data keluarga_anggota_rel")
		}
	}

	if request.Status == "MENINGGAL" {
		errUpdateKepalaKeluarga := repository.UpdateKepalaKeluarga(ctx, tx, int32(idAnggota), request.IdKeluarga)
		if errUpdateKepalaKeluarga != nil {
			return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data kepala keluarga")
		}
	}

	newDataAnggota := entity.DataAnggotaWithStatus{
		Id:            int32(idAnggota),
		NamaLengkap:   request.NamaLengkap,
		TanggalLahir:  request.TanggalLahir,
		TanggalBaptis: request.TanggalBabtis,
		Keterangan:    request.Keterangan,
		Status:        request.Status,
	}

	return newDataAnggota, nil
}

func (repository *dataAnggotaRepositoryImpl) UpdateKeteranganAnggota(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataAnggotaWithKeteranganOnly, error) {
	sqlScript := "UPDATE data_anggota SET keterangan = ? where id = ?"
	sqlScriptRelData := "UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?"
	body := ctx.Body()
	request := new(helper.UpdateKeteranganAnggotaRequest)
	fmt.Println(request)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Update data anggota
	_, err := tx.Exec(sqlScript, request.Keterangan, request.Id)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data anggota")
	}
	_, err = tx.Exec(sqlScriptRelData, "Kepala Keluarga", request.Id)
	if err != nil {
		return entity.DataAnggotaWithKeteranganOnly{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data keluarga_anggota_rel")
	}

	newDataAnggota := entity.DataAnggotaWithKeteranganOnly{
		Id:         request.Id,
		Keterangan: request.Keterangan,
	}

	return newDataAnggota, nil
}
