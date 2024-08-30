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

	if request.Hubungan != "" {
		sqlScriptRelData := "UPDATE keluarga_anggota_rel SET hubungan = ? WHERE id_anggota = ?"
		_, err = tx.Exec(sqlScriptRelData, request.Hubungan, idAnggota)
		if err != nil {
			return entity.DataAnggotaWithStatus{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to update data keluarga_anggota_rel")
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

// func (repository *dataAnggotaRepositoryImpl) FindKeluargaAnggotaRel([]ids int32, tx *sql.Tx) ([]entity.DataAnggota, error) {
// 	sqlScript := "SELECT id, username FROM users WHERE username = ? AND password = ?"

// 	result, err :=tx.Query(sqlScript, request.Username, request.Password)
// 	helper.PanicIfError(err);
// 	defer result.Close()

// 	user := entity.User{}
// 	if result.Next(){
// 		err := result.Scan(&user.Id, &user.Username)
// 		helper.PanicIfError(err)
// 		return user, nil
// 	} else{
// 		return user, fiber.NewError(fiber.StatusNotFound, "user is not found")
// 	}
// }
