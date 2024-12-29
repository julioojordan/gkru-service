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

type dataLingkunganRepositoryImpl struct {
}

func NewDataLingkunganRepository(db *sql.DB) DataLingkunganRepository {
	return &dataLingkunganRepositoryImpl{}
}

func (repository *dataLingkunganRepositoryImpl) FindOneById(id int32, tx *sql.Tx) (entity.DataLingkungan, error) {
	// ini digunakan untuk data keluarga
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id WHERE l.id = ?"
	result, err := tx.Query(sqlScript, id)
	helper.PanicIfError(err)
	defer result.Close()

	raw := entity.DataLingkunganRaw{}
	if result.Next() {
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah)
		helper.PanicIfError(err)
		wilayah := entity.DataWilayah{
			Id:          raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkungan{
			Id:             raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah:        wilayah,
		}
		return lingkungan, nil
	} else {
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}
}

func (repository *dataLingkunganRepositoryImpl) FindOneByIdSeparateTx(id int32, db *sql.DB) (entity.DataLingkungan, error) {
	// ini digunakan untuk data keluarga
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id WHERE l.id = ?"
	result, err := db.Query(sqlScript, id)
	helper.PanicIfError(err)
	defer result.Close()

	raw := entity.DataLingkunganRaw{}
	if result.Next() {
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah)
		helper.PanicIfError(err)
		wilayah := entity.DataWilayah{
			Id:          raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkungan{
			Id:             raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah:        wilayah,
		}
		return lingkungan, nil
	} else {
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}
}

func (repository *dataLingkunganRepositoryImpl) FindOneWithParam(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkungan, error) {
	idLingkungan, err := strconv.Atoi(ctx.Params("idLingkungan"))
	if err != nil {
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id lingkungan, it must be an integer")
	}
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id WHERE l.id = ?"
	result, err := tx.Query(sqlScript, idLingkungan)
	if err != nil {
		return entity.DataLingkungan{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	raw := entity.DataLingkunganRaw{}
	if result.Next() {
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah)
		if err != nil {
			return entity.DataLingkungan{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		wilayah := entity.DataWilayah{
			Id:          raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkungan{
			Id:             raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah:        wilayah,
		}
		return lingkungan, nil
	} else {
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}
}

func (repository *dataLingkunganRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataLingkungan, error) {
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id ORDER BY w.id ASC"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return []entity.DataLingkungan{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	data := []entity.DataLingkungan{}
	for result.Next() {
		raw := entity.DataLingkunganRaw{}
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah)
		if err != nil {
			return []entity.DataLingkungan{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		wilayah := entity.DataWilayah{
			Id:          raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkungan{
			Id:             raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah:        wilayah,
		}
		data = append(data, lingkungan)
	}

	// if len(data) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	// }

	return data, nil
}

func (repository *dataLingkunganRepositoryImpl) FindAllWithTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataLingkunganWithTotalKeluarga, error) {
	sqlScript := `SELECT 
	l.id, 
	l.kode_lingkungan, 
	l.nama_lingkungan, 
	w.id, 
	w.kode_wilayah, 
	w.nama_wilayah, 
	COUNT(k.nomor) AS total_keluarga 
	FROM lingkungan l 
	JOIN wilayah w ON l.id_wilayah = w.id 
	LEFT JOIN data_keluarga k ON l.id = k.id_lingkungan
	GROUP BY l.id, l.nama_lingkungan, l.kode_lingkungan
	ORDER BY w.id ASC`
	result, err := tx.Query(sqlScript)
	if err != nil {
		return []entity.DataLingkunganWithTotalKeluarga{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	data := []entity.DataLingkunganWithTotalKeluarga{}
	for result.Next() {
		raw := entity.DataLingkunganRawWithTotalKeluarga{}
		err := result.Scan(&raw.Id, &raw.KodeLingkungan, &raw.NamaLingkungan, &raw.IdWilayah, &raw.KodeWilayah, &raw.NamaWilayah, &raw.TotalKeluarga)
		if err != nil {
			return []entity.DataLingkunganWithTotalKeluarga{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		wilayah := entity.DataWilayah{
			Id:          raw.IdWilayah,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		lingkungan := entity.DataLingkunganWithTotalKeluarga{
			Id:             raw.Id,
			KodeLingkungan: raw.KodeLingkungan,
			NamaLingkungan: raw.NamaLingkungan,
			Wilayah:        wilayah,
			TotalKeluarga:  raw.TotalKeluarga,
		}
		data = append(data, lingkungan)
	}

	// if len(data) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	// }

	return data, nil
}

func (repository *dataLingkunganRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkunganWithIdWilayah, error) {
	sqlScript := "INSERT INTO lingkungan(kode_lingkungan, nama_lingkungan, id_wilayah) VALUES(?, ?, ?)"
	body := ctx.Body()
	request := new(helper.LingkunganRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := tx.Exec(sqlScript, request.KodeLingkungan, request.NamaLingkungan, request.Wilayah)
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, helper.CreateErrorMessage("Gagal memasukan data lingkungan", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, helper.CreateErrorMessage("Gagal untuk retrieve last inserted ID", err)
	}

	response := entity.DataLingkunganWithIdWilayah{
		Id:             int32(lastInsertId),
		KodeLingkungan: request.KodeLingkungan,
		NamaLingkungan: request.NamaLingkungan,
		Wilayah:        request.Wilayah,
	}

	return response, nil
}

func (repository *dataLingkunganRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkunganWithIdWilayah, error) {
	sqlScript := "UPDATE lingkungan SET"
	idLingkungan, err := strconv.Atoi(ctx.Params("idLingkungan"))
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id lingkungan, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.LingkunganRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	var params []interface{}
	var setClauses []string

	// Dynamically building the SET clause
	if request.KodeLingkungan != "" {
		setClauses = append(setClauses, "kode_lingkungan = ?")
		params = append(params, request.KodeLingkungan)
	}
	if request.NamaLingkungan != "" {
		setClauses = append(setClauses, "nama_lingkungan = ?")
		params = append(params, request.NamaLingkungan)
	}
	if request.Wilayah != 0 {
		setClauses = append(setClauses, "id_wilayah = ?")
		params = append(params, request.Wilayah)
	}

	// Check if there are fields to update
	if len(setClauses) == 0 {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	// Joining all set clauses
	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idLingkungan)

	// Executing the update statement
	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, helper.CreateErrorMessage("Gagal untuk update data lingkungan", err)
	}

	response := entity.DataLingkunganWithIdWilayah{
		Id:             int32(idLingkungan),
		KodeLingkungan: request.KodeLingkungan,
		NamaLingkungan: request.NamaLingkungan,
		Wilayah:        request.Wilayah,
	}

	return response, nil
}

func (repository *dataLingkunganRepositoryImpl) DeleteOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdDataLingkungan, error) {
	repositories := ctx.Locals("repositories").(Repositories)
	sqlScript := "DELETE FROM lingkungan WHERE id = ?"
	idLingkungan, err := strconv.Atoi(ctx.Params("idLingkungan"))
	if err != nil {
		return entity.IdDataLingkungan{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id Lingkungan, it must be an integer")
	}

	//check if kk masih ada yang pakai lingkungan yang ingin di delete -> jika ada throw error
	totalKeluarga, errTotalKeluarga := repositories.DataKeluargaRepository.CountKeluargaWithParam(ctx, tx, "lingkungan", int32(idLingkungan))
	if errTotalKeluarga != nil {
		return entity.IdDataLingkungan{}, errTotalKeluarga
	}

	if totalKeluarga.Total != 0 {
		return entity.IdDataLingkungan{}, fiber.NewError(fiber.StatusInternalServerError, "Gagal untuk delete data lingkungan karena data lingkungan masih digunakan oleh KK")
	}

	// Executing the update statement
	_, err = tx.Exec(sqlScript, idLingkungan)
	if err != nil {
		return entity.IdDataLingkungan{}, helper.CreateErrorMessage("Gagal untuk delete data lingkungan", err)
	}

	response := entity.IdDataLingkungan{
		Id: int32(idLingkungan),
	}

	return response, nil
}

func (repository *dataLingkunganRepositoryImpl) GetTotalLingkungan(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalInt, error) {
	sqlScript := "SELECT COUNT(*) FROM lingkungan"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return entity.TotalInt{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	totalInt := entity.TotalInt{}
	if result.Next() {
		err := result.Scan(&totalInt.Total)
		if err != nil {
			return entity.TotalInt{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return totalInt, nil
	} else {
		return entity.TotalInt{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}
}

func (repository *dataLingkunganRepositoryImpl) CountLingkunganWithIdWilayah(ctx *fiber.Ctx, tx *sql.Tx, idWilayah int32) (entity.TotalInt, error) {
	sqlScript := "SELECT COUNT(*) FROM lingkungan WHERE id_wilayah = ?"
	result, err := tx.Query(sqlScript, idWilayah)
	if err != nil {
		return entity.TotalInt{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	totalInt := entity.TotalInt{}
	if result.Next() {
		err := result.Scan(&totalInt.Total)
		if err != nil {
			return entity.TotalInt{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		return totalInt, nil
	} else {
		return entity.TotalInt{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}
}
