package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gkru-service/entity"
	"gkru-service/helper"
	"net/url"
	"os"
	"path/filepath"
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
	sqlScript := "SELECT SUM(a.nominal) FROM riwayat_transaksi a JOIN data_keluarga b ON a.id_keluarga = b.id WHERE keterangan = 'IN'"
	// kalo mau ada bulanan tinggal kasih ini di sql script
	/* AND MONTH(tanggal) = MONTH(CURRENT_DATE)
	   AND YEAR(tanggal) = YEAR(CURRENT_DATE);*/
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx)
	if err != nil {
		return entity.AmountHistory{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	// Proses hasil query
	AmountHistory := entity.AmountHistory{}
	if result.Next() {
		err := result.Scan(&AmountHistory.Nominal)
		if err != nil {
			return AmountHistory, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
	} else {
		return AmountHistory, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}

	return AmountHistory, nil
}

func (repository *transactionHistoryRepositoryImpl) GetTotalOutcome(ctx *fiber.Ctx, tx *sql.Tx) (entity.AmountHistory, error) {
	// total over all -> tidak ada filter tanggal dll
	idWilayah := ctx.Query("id_wilayah")
	idLingkungan := ctx.Query("id_lingkungan")
	sqlScript := "SELECT SUM(a.nominal) FROM riwayat_transaksi a JOIN data_keluarga b ON a.id_keluarga = b.id WHERE keterangan = 'OUT'"
	// kalo mau ada bulanan tinggal kasih ini di sql script
	/* AND MONTH(tanggal) = MONTH(CURRENT_DATE)
	   AND YEAR(tanggal) = YEAR(CURRENT_DATE);*/
	result, err := helper.AddLingkunganOrWilayahQueryHelper(idWilayah, idLingkungan, sqlScript, tx)
	if err != nil {
		return entity.AmountHistory{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()
	var nominal sql.NullInt64

	if result.Next() {
		err := result.Scan(&nominal)
		if err != nil {
			return entity.AmountHistory{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}
		AmountHistory := entity.AmountHistory{
			Nominal: int32(0),
		}
		if nominal.Valid {
			AmountHistory.Nominal = int32(nominal.Int64)
		}

		return AmountHistory, nil
	} else {
		return entity.AmountHistory{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
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

	File := ""
	if dataThRaw.File.Valid {
		File = dataThRaw.File.String
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
		GroupId:       dataThRaw.GroupId,
		File:          File,
	}
}

// Tambahkan fungsi untuk mapping dari ThRaw ke ThFinal
func mapToThFinal2(dataThRaw entity.ThRaw2) entity.ThFinal2 {
	// Set nilai default jika NULL
	subKeterangan := ""
	if dataThRaw.SubKeterangan.Valid {
		subKeterangan = dataThRaw.SubKeterangan.String
	}

	updatorId := int32(0)
	if dataThRaw.UpdatorId.Valid {
		updatorId = int32(dataThRaw.UpdatorId.Int32)
	}

	return entity.ThFinal2{
		Id:                 dataThRaw.Id,
		Nominal:            dataThRaw.Nominal,
		IdKeluarga:         dataThRaw.IdKeluarga,
		Keterangan:         dataThRaw.Keterangan,
		Creator:            entity.User{Id: dataThRaw.CreatorId, Username: dataThRaw.UserName},
		Wilayah:            entity.DataWilayah{Id: dataThRaw.IdWilayah, KodeWilayah: dataThRaw.KodeWilayah, NamaWilayah: dataThRaw.NamaWilayah},
		Lingkungan:         entity.DataLingkunganWithIdWilayah{Id: dataThRaw.IdLingkungan, KodeLingkungan: dataThRaw.KodeLingkungan, NamaLingkungan: dataThRaw.NamaLingkungan, Wilayah: dataThRaw.IdWilayah},
		UpdatorId:          updatorId,
		SubKeterangan:      subKeterangan,
		CreatedDate:        dataThRaw.CreatedDate,
		UpdatedDate:        dataThRaw.UpdatedDate,
		Bulan:              dataThRaw.Bulan,
		Tahun:              dataThRaw.Tahun,
		GroupId:            dataThRaw.GroupId,
		NamaKepalaKeluarga: dataThRaw.NamaKepalaKeluarga,
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
        c.id_wilayah, 
        c.id_lingkungan, 
        a.updated_by, 
        a.sub_keterangan, 
        a.created_date, 
        a.updated_date, 
        a.bulan,
        a.tahun,
        a.group_id,
        b.username, 
        d.kode_lingkungan, 
        d.nama_lingkungan, 
        e.kode_wilayah, 
        e.nama_wilayah,
		f.file
    FROM 
        riwayat_transaksi a
    JOIN 
        users b ON a.created_by = b.id
	JOIN
		data_keluarga c ON a.id_keluarga = c.id
    JOIN 
        lingkungan d ON c.id_lingkungan = d.id
    JOIN 
        wilayah e ON c.id_wilayah = e.id
	LEFT JOIN 
        grouped_transaksi f ON a.group_id = f.id
    WHERE 
        a.id = ?`

	result, err := tx.Query(sqlScript, idTh)
	if err != nil {
		return entity.ThFinal{}, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	dataThRaw := entity.ThRaw{}

	if result.Next() {
		err := result.Scan(
			&dataThRaw.Id,
			&dataThRaw.Nominal,
			&dataThRaw.IdKeluarga,
			&dataThRaw.Keterangan,
			&dataThRaw.CreatorId,
			&dataThRaw.IdWilayah,
			&dataThRaw.IdLingkungan,
			&dataThRaw.UpdatorId,
			&dataThRaw.SubKeterangan,
			&dataThRaw.CreatedDate,
			&dataThRaw.UpdatedDate,
			&dataThRaw.Bulan,
			&dataThRaw.Tahun,
			&dataThRaw.GroupId, // Tambahan kolom file_bukti
			&dataThRaw.UserName,
			&dataThRaw.KodeLingkungan,
			&dataThRaw.NamaLingkungan,
			&dataThRaw.KodeWilayah,
			&dataThRaw.NamaWilayah,
			&dataThRaw.File,
		)
		if err != nil {
			return entity.ThFinal{}, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

	} else {
		return entity.ThFinal{}, fiber.NewError(fiber.StatusInternalServerError, "Data Tidak Ditemukan")
	}

	response := mapToThFinal(dataThRaw)

	return response, nil
}

func (repository *transactionHistoryRepositoryImpl) FindByGroup(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	idGroup := ctx.Query("idGroup")
	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, c.id_wilayah, c.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   d.kode_lingkungan, d.nama_lingkungan, 
		   e.kode_wilayah, e.nama_wilayah, f.file, a.group_id
	FROM riwayat_transaksi a
	LEFT JOIN users b ON a.created_by = b.id
	LEFT JOIN data_keluarga c ON a.id_keluarga = c.id
	LEFT JOIN lingkungan d ON c.id_lingkungan = d.id
	LEFT JOIN wilayah e ON c.id_wilayah = e.id
	JOIN
		grouped_transaksi f ON a.group_id = f.id
    WHERE 
        a.group_id = ?
	ORDER BY a.created_date ASC`

	// Eksekusi query
	result, err := tx.Query(sqlScript, idGroup)
	if err != nil {
		return nil, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal

	// Iterasi hasil
	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah, &dataThRaw.File, &dataThRaw.GroupId)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Gunakan fungsi mapping
		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	return thFinals, nil
}

// findAll
func (repository *transactionHistoryRepositoryImpl) FindAll(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, c.id_wilayah, c.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   d.kode_lingkungan, d.nama_lingkungan, 
		   e.kode_wilayah, e.nama_wilayah
	FROM riwayat_transaksi a
	LEFT JOIN users b ON a.created_by = b.id
	LEFT JOIN data_keluarga c ON a.id_keluarga = c.id
	LEFT JOIN lingkungan d ON c.id_lingkungan = d.id
	LEFT JOIN wilayah e ON c.id_wilayah = e.id
	ORDER BY a.created_date ASC`

	// Eksekusi query
	result, err := tx.Query(sqlScript)
	if err != nil {
		return nil, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal

	// Iterasi hasil
	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Gunakan fungsi mapping
		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	// Jika tidak ada data
	// if len(thFinals) == 0 {
	// 	return nil, fiber.NewError(fiber.StatusNotFound, "Data Tidak Ditemukan")
	// }

	return thFinals, nil
}

func (repository *transactionHistoryRepositoryImpl) FindAllWithKeluargaContext(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal2, error) {
	tahun := ctx.Query("tahun")
	sqlScript := `
		SELECT 
		a.id, 
		a.nominal, 
		a.id_keluarga, 
		a.keterangan, 
		a.created_by, 
		c.id_wilayah, 
		c.id_lingkungan, 
		a.updated_by, 
		a.sub_keterangan, 
		a.created_date, 
		a.updated_date, 
		a.bulan, 
		a.tahun,
		b.username, 
		d.kode_lingkungan, 
		d.nama_lingkungan, 
		e.kode_wilayah, 
		e.nama_wilayah,
		f.nama_lengkap AS nama_kepala_keluarga
	FROM 
		riwayat_transaksi a
	LEFT JOIN users b 
		ON a.created_by = b.id
	LEFT JOIN data_keluarga c 
		ON a.id_keluarga = c.id
	LEFT JOIN lingkungan d 
		ON c.id_lingkungan = d.id
	LEFT JOIN wilayah e 
		ON c.id_wilayah = e.id
	LEFT JOIN data_anggota f 
		ON c.id_kepala_keluarga = f.id
	WHERE
		YEAR(a.created_date) = ?
	ORDER BY 
		a.created_date ASC;
	`

	// Eksekusi query
	result, err := tx.Query(sqlScript, tahun)
	if err != nil {
		return nil, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal2

	// Iterasi hasil
	for result.Next() {
		dataThRaw := entity.ThRaw2{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah, &dataThRaw.NamaKepalaKeluarga)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Gunakan fungsi mapping
		thFinal := mapToThFinal2(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	return thFinals, nil
}

// TO DO WIP
func (repository *transactionHistoryRepositoryImpl) FindAllSetoran(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal2, error) {
	tahun := ctx.Query("tahun")
	bulan := ctx.Query("bulan")
	sqlScript := `
		SELECT 
		a.id, 
		a.nominal, 
		a.id_keluarga, 
		a.keterangan, 
		a.created_by, 
		c.id_wilayah, 
		c.id_lingkungan, 
		a.updated_by, 
		a.sub_keterangan, 
		a.created_date, 
		a.updated_date, 
		a.bulan, 
		a.tahun,
		b.username, 
		d.kode_lingkungan, 
		d.nama_lingkungan, 
		e.kode_wilayah, 
		e.nama_wilayah,
		f.nama_lengkap AS nama_kepala_keluarga
	FROM 
		riwayat_transaksi a
	LEFT JOIN users b 
		ON a.created_by = b.id
	LEFT JOIN data_keluarga c 
		ON a.id_keluarga = c.id
	LEFT JOIN lingkungan d 
		ON c.id_lingkungan = d.id
	LEFT JOIN wilayah e 
		ON c.id_wilayah = e.id
	LEFT JOIN data_anggota f 
		ON c.id_kepala_keluarga = f.id
	WHERE
		tahun = ? AND bulan = ?
	ORDER BY 
		a.created_date ASC;
	`

	// Eksekusi query
	result, err := tx.Query(sqlScript, tahun, bulan)
	if err != nil {
		return nil, helper.CreateErrorMessage("Gagal mengeksekusi query", err)
	}
	defer result.Close()

	var thFinals []entity.ThFinal2

	// Iterasi hasil
	for result.Next() {
		dataThRaw := entity.ThRaw2{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah, &dataThRaw.NamaKepalaKeluarga)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		// Gunakan fungsi mapping
		thFinal := mapToThFinal2(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	return thFinals, nil
}

// findAllWithIdKeluarga -> untuk mengecek history pembayaran keluarga
func (repository *transactionHistoryRepositoryImpl) FindAllWithIdKeluarga(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	idKeluargaStr := ctx.Query("idKeluarga")
	tahunStr := ctx.Query("tahun")

	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, c.id_wilayah, c.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   d.kode_lingkungan, d.nama_lingkungan, 
		   e.kode_wilayah, e.nama_wilayah
	FROM riwayat_transaksi a
	JOIN users b ON a.created_by = b.id
	JOIN data_keluarga c ON a.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id
	WHERE a.id_keluarga = ?`

	args := []interface{}{idKeluargaStr}

	if tahunStr != "" {
		sqlScript += " AND a.tahun = ?"
		args = append(args, tahunStr)
	}

	sqlScript += " ORDER BY a.created_date ASC"

	result, err := tx.Query(sqlScript, args...)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var thFinals []entity.ThFinal

	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
	}

	// if len(thFinals) == 0 {
	// 	return thFinals, nil
	// }

	return thFinals, nil
}

func (repository *transactionHistoryRepositoryImpl) FindAllHistoryWithTimeFilter(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.ThFinal, error) {
	tahun := ctx.Query("tahun")
	bulan := ctx.Query("bulan")
	idLingkungan := ctx.Query("lingkungan")
	idWilayah := ctx.Query("wilayah")

	sqlScript := `
	SELECT a.id, a.nominal, a.id_keluarga, a.keterangan, a.created_by, c.id_wilayah, c.id_lingkungan, a.updated_by, a.sub_keterangan, a.created_date, a.updated_date, a.bulan, a.tahun,
		   b.username, 
		   d.kode_lingkungan, d.nama_lingkungan, 
		   e.kode_wilayah, e.nama_wilayah
	FROM riwayat_transaksi a
	JOIN users b ON a.created_by = b.id
	JOIN data_keluarga c ON a.id_keluarga = c.id
	JOIN lingkungan d ON c.id_lingkungan = d.id
	JOIN wilayah e ON c.id_wilayah = e.id
	WHERE a.tahun = ? AND a.bulan = ?`

	args := []interface{}{tahun, bulan}

	if idLingkungan != "" {
		sqlScript += " AND c.id_lingkungan = ?"
		args = append(args, idLingkungan)
	}

	if idWilayah != "" {
		sqlScript += " AND c.id_wilayah = ?"
		args = append(args, idWilayah)
	}

	sqlScript += " ORDER BY a.created_date ASC"

	result, err := tx.Query(sqlScript, args...)
	if err != nil {
		return nil, err
	}

	defer result.Close()

	var thFinals []entity.ThFinal

	for result.Next() {
		dataThRaw := entity.ThRaw{}
		err := result.Scan(&dataThRaw.Id, &dataThRaw.Nominal, &dataThRaw.IdKeluarga, &dataThRaw.Keterangan, &dataThRaw.CreatorId, &dataThRaw.IdWilayah, &dataThRaw.IdLingkungan, &dataThRaw.UpdatorId, &dataThRaw.SubKeterangan, &dataThRaw.CreatedDate, &dataThRaw.UpdatedDate, &dataThRaw.Bulan, &dataThRaw.Tahun, &dataThRaw.UserName, &dataThRaw.KodeLingkungan, &dataThRaw.NamaLingkungan, &dataThRaw.KodeWilayah, &dataThRaw.NamaWilayah)
		if err != nil {
			return nil, helper.CreateErrorMessage("Gagal untuk scan result", err)
		}

		thFinal := mapToThFinal(dataThRaw)
		thFinals = append(thFinals, thFinal)
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
		return entity.UpdatedThFinal{}, helper.CreateErrorMessage("Gagal untuk update data riwayat transaksi", err)
	}

	response := entity.UpdatedThFinal{
		Id:            int32(idTh),
		IdKeluarga:    request.IdKeluarga,
		Keterangan:    request.Keterangan,
		SubKeterangan: request.SubKeterangan,
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
		return entity.IdInt{}, helper.CreateErrorMessage("Gagal untuk delete data riwayat transaksi", err)
	}

	response := entity.IdInt{
		Id: int32(idTh),
	}

	return response, nil
}

// Add Santunan
func (repository *transactionHistoryRepositoryImpl) Add(ctx *fiber.Ctx, tx *sql.Tx) (entity.CreatedTh, error) {
	// untuk add santunan
	sqlScript := "INSERT INTO riwayat_transaksi(nominal, id_keluarga, keterangan, created_by, sub_keterangan, created_date, bulan, tahun, group_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"

	// Parse form data
	form, err := ctx.MultipartForm()
	if err != nil {
		errorMessage := fmt.Sprintf("%s: %v", "Invalid request body", err.Error())
		return entity.CreatedTh{}, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	// Extract JSON fields from form
	nominal, _ := strconv.Atoi(form.Value["Nominal"][0])
	idKeluarga, _ := strconv.Atoi(form.Value["IdKeluarga"][0])
	keterangan := form.Value["Keterangan"][0]
	createdBy, _ := strconv.Atoi(form.Value["CreatedBy"][0])
	subKeterangan := form.Value["SubKeterangan"][0]
	bulan, _ := strconv.Atoi(form.Value["Bulan"][0])
	tahun, _ := strconv.Atoi(form.Value["Tahun"][0])
	currentTime := time.Now()

	// Handle file upload
	fileHeader := form.File["FileBukti"]
	var filePath string
	var lastInsertIdAddFile int64
	if len(fileHeader) > 0 {
		file := fileHeader[0]

		// Get root directory
		rootPath, err := os.Getwd()
		if err != nil {
			return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, "Gagal untuk get working directory")
		}

		// Ensure uploads folder exists
		uploadDir := filepath.Join(rootPath, "uploads")
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, "Gagal untuk create upload directory")
			}
		}

		// Define destination path
		destination := filepath.Join(uploadDir, file.Filename)

		// Save the file
		if err := ctx.SaveFile(file, destination); err != nil {
			return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, "Gagal untuk save file")
		}

		// Encode file name supaya menghindari jika ada spasi di nama file
		safeFileName := url.QueryEscape(file.Filename)
		filePath = "/uploads/" + safeFileName

		//insert to grouped transaction
		resultAddFile, err := tx.Exec("INSERT INTO grouped_transaksi(file) VALUES(?)", filePath)
		if err != nil {
			errorMessage := fmt.Sprintf("%s: %v", "Failed to insert data", err.Error())
			return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
		}

		lastInsertIdAddFile, err = resultAddFile.LastInsertId()
		if err != nil {
			return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to get last inserted ID")
		}
	}

	// Execute database query
	result, err := tx.Exec(sqlScript, nominal, idKeluarga, keterangan, createdBy, subKeterangan, currentTime, bulan, tahun, lastInsertIdAddFile)
	if err != nil {
		errorMessage := fmt.Sprintf("%s: %v", "Gagal memasukan data", err.Error())
		return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.CreatedTh{}, fiber.NewError(fiber.StatusInternalServerError, "Gagal mendapatkan ID data yang terakhir dimasukan")
	}

	response := entity.CreatedTh{
		Id:            int32(lastInsertId),
		Nominal:       int32(nominal),
		IdKeluarga:    int32(idKeluarga),
		Keterangan:    keterangan,
		CreatorId:     int32(createdBy),
		SubKeterangan: subKeterangan,
		CreatedDate:   currentTime,
		Group:         lastInsertIdAddFile,
	}

	return response, nil
}

// Add Iuran
func (repository *transactionHistoryRepositoryImpl) AddBatch(ctx *fiber.Ctx, tx *sql.Tx) ([]entity.CreatedTh, error) {
	// untuk add iuran yang baru
	form, err := ctx.MultipartForm()
	if err != nil {
		errorMessage := fmt.Sprintf("%s: %v", "Invalid request body", err.Error())
		return nil, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	// Extract JSON array from form
	historyData := form.Value["History"]
	if len(historyData) == 0 {
		return nil, fiber.NewError(fiber.StatusBadRequest, "History data is missing")
	}

	// Decode history array
	var histories []map[string]interface{}
	if err := json.Unmarshal([]byte(historyData[0]), &histories); err != nil {
		errorMessage := fmt.Sprintf("%s: %v", "Failed to parse history data", err.Error())
		return nil, fiber.NewError(fiber.StatusBadRequest, errorMessage)
	}

	// Handle file upload (optional)
	fileHeader := form.File["FileBukti"]
	var filePath string
	var lastInsertIdAddFile int64
	if len(fileHeader) > 0 {
		file := fileHeader[0]

		// Get root directory
		rootPath, err := os.Getwd()
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get working directory")
		}

		// Ensure uploads folder exists
		uploadDir := filepath.Join(rootPath, "uploads")
		if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
			if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
				return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to create upload directory")
			}
		}

		// Define destination path
		destination := filepath.Join(uploadDir, file.Filename)

		// Save the file
		if err := ctx.SaveFile(file, destination); err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save file")
		}

		// Encode file name to avoid issues with spaces
		safeFileName := url.QueryEscape(file.Filename)
		filePath = "/uploads/" + safeFileName

		//insert to grouped transaction
		resultAddFile, err := tx.Exec("INSERT INTO grouped_transaksi(file) VALUES(?)", filePath)
		if err != nil {
			errorMessage := fmt.Sprintf("%s: %v", "Failed to insert data", err.Error())
			return nil, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
		}

		lastInsertIdAddFile, err = resultAddFile.LastInsertId()
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get last inserted ID")
		}
	}

	// Prepare for batch insert
	sqlScript := "INSERT INTO riwayat_transaksi(nominal, id_keluarga, keterangan, created_by, sub_keterangan, created_date, bulan, tahun, group_id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var responses []entity.CreatedTh

	for _, history := range histories {
		nominal := int(history["Nominal"].(float64)) // Cast from float64 to int
		idKeluarga := int(history["IdKeluarga"].(float64))
		keterangan := history["Keterangan"].(string)
		createdBy := int(history["CreatedBy"].(float64))
		subKeterangan := history["SubKeterangan"].(string)
		bulan := int(history["Bulan"].(float64))
		tahun := int(history["Tahun"].(float64))
		createdDate := time.Now()

		// Execute query
		result, err := tx.Exec(sqlScript, nominal, idKeluarga, keterangan, createdBy, subKeterangan, createdDate, bulan, tahun, lastInsertIdAddFile)
		if err != nil {
			errorMessage := fmt.Sprintf("%s: %v", "Failed to insert data", err.Error())
			return nil, fiber.NewError(fiber.StatusInternalServerError, errorMessage)
		}

		lastInsertId, err := result.LastInsertId()
		if err != nil {
			return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to get last inserted ID")
		}

		// Append response
		responses = append(responses, entity.CreatedTh{
			Id:            int32(lastInsertId),
			Nominal:       int32(nominal),
			IdKeluarga:    int32(idKeluarga),
			Keterangan:    keterangan,
			CreatorId:     int32(createdBy),
			SubKeterangan: subKeterangan,
			CreatedDate:   createdDate,
			Group:         lastInsertIdAddFile,
		})
	}

	return responses, nil
}
