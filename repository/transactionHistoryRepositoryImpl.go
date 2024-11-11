package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gkru-service/entity"
	"gkru-service/helper"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type transactionHistoryRepositoryImpl struct {
}

func NewTransactionHistoryRepository(db *sql.DB) TransactionHistoryRepository {
	return &transactionHistoryRepositoryImpl{}
}

func (repository *transactionHistoryRepositoryImpl) GetTotalIncome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	// total over all -> tidak ada filter tanggal dll
	idWilayah := ctx.Query("id_wilayah")
	idLingkungan := ctx.Query("id_lingkungan")
	sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'IN'"
	// kalo mau ada bulanan tinggal kasih ini di sql script
	/* AND MONTH(tanggal) = MONTH(CURRENT_DATE)
	   AND YEAR(tanggal) = YEAR(CURRENT_DATE);*/
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx)
	if err != nil {
		return entity.AmountHistory{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	// Proses hasil query
	AmountHistory := entity.AmountHistory{}
	if result.Next() {
		err := result.Scan(&AmountHistory.Nominal)
		if err != nil {
			return AmountHistory, helper.CreateErrorMessage("Failed to scan result", err)
		}
	} else {
		return AmountHistory, fiber.NewError(fiber.StatusInternalServerError, "No Data Found")
	}

	return AmountHistory, nil
}

func (repository *transactionHistoryRepositoryImpl) GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	// total over all -> tidak ada filter tanggal dll
	idWilayah := ctx.Query("id_wilayah")
	idLingkungan := ctx.Query("id_lingkungan")
	sqlScript := "SELECT SUM(nominal) FROM riwayat_transaksi WHERE keterangan = 'OUT'"
	// kalo mau ada bulanan tinggal kasih ini di sql script
	/* AND MONTH(tanggal) = MONTH(CURRENT_DATE)
	   AND YEAR(tanggal) = YEAR(CURRENT_DATE);*/
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx)
	if err != nil {
		return entity.AmountHistory{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()
	var nominal sql.NullInt64

	if result.Next() {
		err := result.Scan(&nominal)
		if err != nil {
			return entity.AmountHistory{}, helper.CreateErrorMessage("Failed to scan result", err)
		}
		AmountHistory := entity.AmountHistory{
			Nominal: int32(0),
		}
		if nominal.Valid {
			AmountHistory.Nominal = int32(nominal.Int64)
		}

		return AmountHistory, nil
	} else {
		return entity.AmountHistory{}, fiber.NewError(fiber.StatusInternalServerError, "No Data Found")
	}
}

// Tambahkan fungsi untuk mapping dari ThRaw ke ThFinal
func mapToThFinal(dataThRaw entity.ThRaw) entity.ThFinal {
	// Set nilai default jika NULL
	subKeterangan := ""
	if dataThRaw.SubKeterangan.Valid {
		subKeterangan = dataThRaw.SubKeterangan.String
	}

	updatorId := int32(0)
	if dataThRaw.UpdatorId.Valid {
		updatorId = int32(dataThRaw.UpdatorId.Int32)
	}

	return entity.ThFinal{
		Id:            dataThRaw.Id,
		Nominal:       dataThRaw.Nominal,
		IdKeluarga:    dataThRaw.IdKeluarga,
		Keterangan:    dataThRaw.Keterangan,
		Creator:       entity.User{Id: dataThRaw.CreatorId, Username: dataThRaw.UserName},
		Wilayah:       entity.DataWilayah{Id: dataThRaw.IdWilayah, KodeWilayah: dataThRaw.KodeWilayah, NamaWilayah: dataThRaw.NamaWilayah},
		Lingkungan:    entity.DataLingkunganWithIdWilayah{Id: dataThRaw.IdLingkungan, KodeLingkungan: dataThRaw.KodeLingkungan, NamaLingkungan: dataThRaw.NamaLingkungan, Wilayah: dataThRaw.IdWilayah},
		UpdatorId:     updatorId,
		SubKeterangan: subKeterangan,
		CreatedDate:   dataThRaw.CreatedDate,
		UpdatedDate:   dataThRaw.UpdatedDate,
		Bulan:         dataThRaw.Bulan,
		Tahun:         dataThRaw.Tahun,
	}
}

// findOne
func (repository *transactionHistoryRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.ThFinal, error) {
	idTh, err := strconv.Atoi(ctx.Params("idTh"))
	if err != nil {
		return entity.ThFinal{}, fiber.NewError(fiber.StatusBadRequest, "Invalid id TH, it must be an integer")
	}
	sqlScript := `
    SELECT 
        a.id, 
        a.nominal, 
        a.id_keluarga, 
        a.keterangan, 
        a.created_by, 
        a.id_wilayah, 
        a.id_lingkungan, 
        a.updated_by, 
        a.sub_keterangan, 
        a.created_date, 
        a.updated_date, 
		a.bulan,
		a.tahun,
        b.username, 
        c.kode_lingkungan, 
        c.nama_lingkungan, 
        d.kode_wilayah, 
        d.nama_wilayah
    FROM 
        riwayat_transaksi a
    JOIN 
        users b ON a.created_by = b.id
    JOIN 
        lingkungan c ON a.id_lingkungan = c.id
    JOIN 
        wilayah d ON a.id_wilayah = d.id
    WHERE 
        a.id = ?`

	result, err := tx.Query(sqlScript, idTh)
	if err != nil {
		return entity.ThFinal{}, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	dataThRaw := entity.ThRaw{}
	if result.Next() {
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return entity.ThFinal{}, helper.CreateErrorMessage("Failed to scan result", err)
		}
	} else {
		return entity.ThFinal{}, fiber.NewError(fiber.StatusInternalServerError, "No data found")
	}

	response := mapToThFinal(dataThRaw)

	return response, nil
}

// findAll
func (repository *transactionHistoryRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, a.id_wilayah, a.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   c.kode_lingkungan, c.nama_lingkungan, 
		   d.kode_wilayah, d.nama_wilayah
	FROM riwayat_transaksi a
	LEFT JOIN users b ON a.created_by = b.id
	LEFT JOIN lingkungan c ON a.id_lingkungan = c.id
	LEFT JOIN wilayah d ON a.id_wilayah = d.id
	ORDER BY a.created_date ASC`

	// Eksekusi query
	result, err := tx.Query(sqlScript)
	if err != nil {
		return nil, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal

	// Iterasi hasil
	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Failed to scan result", err)
		}

		// Gunakan fungsi mapping
		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	// Jika tidak ada data
	if len(thFinals) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "No data found")
	}

	return thFinals, nil
}

// [TBD]findWithFilter -> untuk export maybe (?)

// findAllWithIdKeluarga -> untuk mengecek history pembayaran keluarga
func (repository *transactionHistoryRepositoryImpl) FindAllWithIdKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	idKeluargaStr := ctx.Query("idKeluarga")
	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, a.id_wilayah, a.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   c.kode_lingkungan, c.nama_lingkungan, 
		   d.kode_wilayah, d.nama_wilayah
	FROM riwayat_transaksi a
	JOIN users b ON a.created_by = b.id
	JOIN lingkungan c ON a.id_lingkungan = c.id
	JOIN wilayah d ON a.id_wilayah = d.id
	WHERE a.id_keluarga = ?
	ORDER BY a.created_date ASC`

	result, err := tx.Query(sqlScript, idKeluargaStr)
	if err != nil {
		return nil, helper.CreateErrorMessage("Failed to execute query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal

	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Failed to scan result", err)
		}

		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	if len(thFinals) == 0 {
		return nil, fiber.NewError(fiber.StatusNotFound, "No data found")
	}

	return thFinals, nil
}

// update
func (repository *transactionHistoryRepositoryImpl) Update(ctx *fiber.Ctx, tx *sql.Tx) (entity.UpdatedThFinal, error) {
	sqlScript := "UPDATE riwayat_transaksi SET"
	idTh, err := strconv.Atoi(ctx.Params("idTh"))
	if err != nil {
		return entity.UpdatedThFinal{}, fiber.NewError(fiber.StatusBadRequest, "Invalid TH, it must be an integer")
	}
	body := ctx.Body()
	request := new(helper.UpdateTHRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		fmt.Println(marshalError)
		return entity.UpdatedThFinal{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	currentTime := time.Now()
	var params []interface{}
	var setClauses []string

	// Dynamically building the SET clause
	if request.Nominal != 0 {
		setClauses = append(setClauses, "nominal = ?")
		params = append(params, request.Nominal)
	}
	if request.Keterangan != "" {
		setClauses = append(setClauses, "keterangan = ?")
		params = append(params, request.Keterangan)
	}
	if request.SubKeterangan != "" {
		setClauses = append(setClauses, "sub_keterangan = ?")
		params = append(params, request.SubKeterangan)
	}
	if request.IdKeluarga != 0 {
		setClauses = append(setClauses, "id_keluarga = ?")
		params = append(params, request.IdKeluarga)
	}
	if request.IdLingkungan != 0 {
		setClauses = append(setClauses, "id_lingkungan = ?")
		params = append(params, request.IdLingkungan)
	}
	if request.IdWilayah != 0 {
		setClauses = append(setClauses, "id_wilayah = ?")
		params = append(params, request.IdWilayah)
	}

	// Check if there are fields to update
	if len(setClauses) == 0 {
		return entity.UpdatedThFinal{}, fiber.NewError(fiber.StatusBadRequest, "Error No fields to update")
	}

	// Joining all set clauses
	setClauses = append(setClauses, "updated_by = ?")
	params = append(params, request.UpdatedBy)
	setClauses = append(setClauses, "updated_date = ?")
	params = append(params, currentTime)
	sqlScript += " " + strings.Join(setClauses, ", ") + " WHERE id = ?"
	params = append(params, idTh)

	// Executing the update statement
	_, err = tx.Exec(sqlScript, params...)
	if err != nil {
		fmt.Println(err)
		return entity.UpdatedThFinal{}, helper.CreateErrorMessage("Failed to update data riwayat transaksi", err)
	}

	response := entity.UpdatedThFinal{
		Id:            int32(idTh),
		IdKeluarga:    request.IdKeluarga,
		Keterangan:    request.Keterangan,
		SubKeterangan: request.SubKeterangan,
		IdWilayah:     request.IdWilayah,
		IdLingkungan:  request.IdLingkungan,
		UpdatedDate:   currentTime,
		UpdatorId:     request.UpdatedBy,
	}

	return response, nil
}

// delete
func (repository *transactionHistoryRepositoryImpl) Delete(ctx *fiber.Ctx, tx *sql.Tx) (entity.IdInt, error) {
	sqlScript := "DELETE FORM riwayat_transaksi WHERE id = ?"
	idTh, err := strconv.Atoi(ctx.Params("idTh"))
	if err != nil {
		return entity.IdInt{}, fiber.NewError(fiber.StatusBadRequest, "Invalid TH, it must be an integer")
	}

	// Executing the update statement
	_, err = tx.Exec(sqlScript, idTh)
	if err != nil {
		fmt.Println(err)
		return entity.IdInt{}, helper.CreateErrorMessage("Failed to delete data riwayat transaksi", err)
	}

	response := entity.IdInt{
		Id: int32(idTh),
	}

	return response, nil
}

// Add
func (repository *transactionHistoryRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.CreatedTh, error) {
	sqlScript := "INSERT INTO riwayat_transaksi(nominal, id_keluarga, keterangan, created_by, id_wilayah, id_lingkungan, sub_keterangan, created_date, bulan, tahun) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	body := ctx.Body()
	request := new(helper.AddTHRequest)
	marshalError := json.Unmarshal(body, request)
	if marshalError != nil {
		errorMessage := fmt.Sprintf("%s: %v", "Invalid request body", marshalError.Error())
		return entity.CreatedTh{}, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	currentTime := time.Now()

	result, err := tx.Exec(sqlScript, request.Nominal, request.IdKeluarga, request.Keterangan, request.CreatedBy, request.IdWilayah, request.IdLingkungan, request.SubKeterangan, currentTime, request.Bulan, request.Tahun)
	if err != nil {
		fmt.Println(err)
		return entity.CreatedTh{}, helper.CreateErrorMessage("Failed to insert data anggot", err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.CreatedTh{}, helper.CreateErrorMessage("Failed to retrieve last insert ID", err)
	}

	response := entity.CreatedTh{
		Id:            int32(lastInsertId),
		Nominal:       request.Nominal,
		IdKeluarga:    request.IdKeluarga,
		Keterangan:    request.Keterangan,
		CreatorId:     request.CreatedBy,
		IdWilayah:     request.IdWilayah,
		IdLingkungan:  request.IdLingkungan,
		SubKeterangan: request.SubKeterangan,
		CreatedDate:   request.CreatedDate,
	}

	return response, nil
}
