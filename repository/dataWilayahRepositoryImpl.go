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

type dataWilayahRepositoryImpl struct {
}

func NewDataWilayahRepository(db *sql.DB) DataWilayahRepository {
	return &dataWilayahRepositoryImpl{}
}


func (repository *dataWilayahRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error) {
	IdWilayah, err := strconv.Atoi(ctx.Params("IdWilayah"))
	if err != nil {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id wilayah, it must be an integer")
	}
	sqlScript := "SELECT id, kode_wilayah, nama_wilayah FROM wilayah WHERE id = ?"
	result, err := tx.Query(sqlScript, IdWilayah)
	if err != nil {
		return entity.DataWilayah{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	response := entity.DataWilayah{}
	if result.Next() {
		err := result.Scan(&response.Id, &response.KodeWilayah, &response.NamaWilayah)
		if err != nil {
			return entity.DataWilayah{}, helper.CreateErrorMessage("Failed to scan result", err)
		}
		wilayah := entity.DataWilayah{
			Id:          response.Id,
			KodeWilayah: response.KodeWilayah,
			NamaWilayah: response.NamaWilayah,
		}
		return wilayah, nil
	} else {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusNotFound, "wilayah is not found")
	}
}

func (repository *dataWilayahRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.DataWilayah, error) {
	sqlScript := "SELECT id, kode_wilayah, nama_wilayah FROM wilayah ORDER BY id ASC"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return []entity.DataWilayah{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	dataList := []entity.DataWilayah{}
	for result.Next() {
		raw := entity.DataWilayah{}
		err := result.Scan(&raw.Id, &raw.KodeWilayah, &raw.NamaWilayah)
		if err != nil {
			return []entity.DataWilayah{}, helper.CreateErrorMessage("Failed to scan result", err)
		}
		wilayah := entity.DataWilayah{
			Id:          raw.Id,
			KodeWilayah: raw.KodeWilayah,
			NamaWilayah: raw.NamaWilayah,
		}
		dataList = append(dataList, wilayah)
	}

	if len(dataList) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "lingkungan is not found")
	}

	return dataList, nil
}

func (repository *dataWilayahRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error) {
	sqlScript := "INSERT INTO data_wilayah(kode_wilayah, nama_wilayah) VALUES(?, ?,)"
	body := ctx.Body()
	request := new(helper.WilayahRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	result, err := tx.Exec(sqlScript, request.KodeWilayah, request.NamaWilayah)
	if err != nil {
		return entity.DataWilayah{}, helper.CreateErrorMessage("Failed to insert data wilayah", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataWilayah{}, helper.CreateErrorMessage("Failed to retrieve last inserted ID", err)
	}

	response := entity.DataWilayah{
		Id:          int32(lastInsertId),
		KodeWilayah: request.KodeWilayah,
		NamaWilayah: request.NamaWilayah,
	}

	return response, nil
}

func (repository *dataWilayahRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataWilayah, error) {
	sqlScript := "UPDATE data_wilayah SET"
	idWilayah, err := strconv.Atoi(ctx.Params("idWilayah"))
	if err != nil {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id wilayah, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.WilayahRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	var params []interface{}
	var setClauses []string

	// Dynamically building the SET clause
	if request.KodeWilayah != "" {
		setClauses = append(setClauses, "kode_wilayah = ?")
		params = append(params, request.KodeWilayah)
	}
	if request.NamaWilayah != "" {
		setClauses = append(setClauses, "nama_wilayah = ?")
		params = append(params, request.NamaWilayah)
	}

	// Check if there are fields to update
	if len(setClauses) == 0 {
		return entity.DataWilayah{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	// Joining all set clauses
	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idWilayah)

	// Executing the update statement
	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		return entity.DataWilayah{}, helper.CreateErrorMessage("Failed to update data wilayah", err)
	}

	response := entity.DataWilayah{
		Id:          int32(idWilayah),
		KodeWilayah: request.KodeWilayah,
		NamaWilayah: request.NamaWilayah,
	}

	return response, nil
}

func (repository *dataWilayahRepositoryImpl) DeleteOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error) {
	repositories := ctx.Locals("repositories").(Repositories)
	sqlScript := "DELETE data_wilayah WHERE id = ?"
	idWilayah, err := strconv.Atoi(ctx.Params("idWilayah"))
	if err != nil {
		return entity.IdInt{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id Wilayah, it must be an integer")
	}

	//check if kk masih ada yang pakai wilayah yang ingin di delete -> jika ada throw error
	totalKeluarga, errTotalKeluarga := repositories.DataKeluargaRepository.CountKeluargaWithParam(ctx, tx, "wilayah", int32(idWilayah))
	if errTotalKeluarga != nil {
		return entity.IdInt{}, errTotalKeluarga
	}

	if totalKeluarga.Total != 0 {
		return entity.IdInt{}, helper.CreateErrorMessage("Failed to delete data wilayah karena data wilayah masih digunakan oleh KK", err)
	}

	// Executing the update statement
	_, err = tx.Exec(sqlScript, idWilayah)
	if err != nil {
		return entity.IdInt{}, helper.CreateErrorMessage("Failed to delete data wilayah", err)
	}

	response := entity.IdInt{
		Id: int32(idWilayah),
	}

	return response, nil
}
