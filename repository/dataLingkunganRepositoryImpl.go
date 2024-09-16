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

func (repository *dataLingkunganRepositoryImpl) FindOneWithParam(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkungan, error) {
	idLingkungan, err := strconv.Atoi(ctx.Params("idLingkungan"))
	if err != nil {
		return entity.DataLingkungan{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id lingkungan, it must be an integer")
	}
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id WHERE l.id = ?"
	result, err := tx.Query(sqlScript, idLingkungan)
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

func (repository *dataLingkunganRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataLingkungan, error) {
	sqlScript := "SELECT l.id, l.kode_lingkungan, l.nama_lingkungan, w.id, w.kode_wilayah, w.nama_wilayah FROM lingkungan l JOIN wilayah w ON l.id_wilayah = w.id ORDER BY w.id ASC"
	result, err := tx.Query(sqlScript)
	helper.PanicIfError(err)
	defer result.Close()

	raw := entity.DataLingkunganRaw{}
	data := []entity.DataLingkungan{}
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
		data = append(data, lingkungan)
		return data, nil
	} else {
		return []entity.DataLingkungan{}, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}
}

func (repository *dataLingkunganRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataLingkunganWithIdWilayah, error) {
	sqlScript := "INSERT INTO data_lingkungan(kode_lingkungan, nama_lingkungan, id_wilayah) VALUES(?, ?, ?,)"
	body := ctx.Body()
	request := new(helper.LingkunganRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := tx.Exec(sqlScript, request.KodeLingkungan, request.NamaLingkungan, request.Wilayah)
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to insert data lingkungan")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve last insert ID")
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
	sqlScript := "UPDATE data_lingkungan SET"
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
		return entity.DataLingkunganWithIdWilayah{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data lingkungan")
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
	sqlScript := "DELETE data_lingkungan WHERE id = ?"
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
		return entity.IdDataLingkungan{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete data lingkungan karena data lingkungan masih digunakan oleh KK")
	}

	// Executing the update statement
	_, err = tx.Exec(sqlScript, idLingkungan)
	if err != nil {
		return entity.IdDataLingkungan{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to delete data lingkungan")
	}

	response := entity.IdDataLingkungan{
		Id: int32(idLingkungan),
	}

	return response, nil
}
