package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"

	// "encoding/json"
	// "fmt"
	"gkru-service/entity"
	"gkru-service/helper"

	"github.com/gofiber/fiber/v2"
)

type dataKeluargaRepositoryImpl struct {
}

func NewDataKeluargaRepository(db *sql.DB) DataKeluargaRepository {
	return &dataKeluargaRepositoryImpl{}
}

func (repository *dataKeluargaRepositoryImpl) FindOne(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluargaFinal, error) {
	dataKeluargaRawScript := "SELECT id, id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, id_keluarga_anggota_rel, alamat FROM data_keluarga WHERE id = ?"
	idKeluarga := ctx.Params("idKeluarga")

	result, err := tx.Query(dataKeluargaRawScript, idKeluarga)
	if err != nil {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to execute query")
	}

	dataKeluargaRaw := entity.DataKeluargaRaw{}
	if result.Next() {
		err := result.Scan(&dataKeluargaRaw.Id, &dataKeluargaRaw.Wilayah, &dataKeluargaRaw.Lingkungan, &dataKeluargaRaw.Nomor, &dataKeluargaRaw.KepalaKeluarga, &dataKeluargaRaw.KKRelation, &dataKeluargaRaw.Alamat)
		if err != nil {
			return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}
	} else {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusNotFound, "Data Keluarga is not found")
	}
	result.Close()

	repositories := ctx.Locals("repositories").(Repositories)

	getLingkungan, err := repositories.DataLingkunganRepository.FindOneById(dataKeluargaRaw.Lingkungan, tx)
	if err != nil {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve lingkungan data")
	}
	wilayah := getLingkungan.Wilayah
	lingkungan := entity.DataLingkungan{
		Id:             getLingkungan.Id,
		KodeLingkungan: getLingkungan.KodeLingkungan,
		NamaLingkungan: getLingkungan.NamaLingkungan,
		Wilayah:        wilayah,
	}

	getAnggotaRel, err := repositories.DataAnggotaKeluargaRelRepository.FindKeluargaAnggotaRel(dataKeluargaRaw.Id, tx)
	if err != nil {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve anggota relationship data")
	}

	var kepalaKeluarga entity.DataAnggota
	var anggota []entity.DataAnggota

	for _, anggotaRel := range getAnggotaRel {
		if anggotaRel.Hubungan == "Kepala Keluarga" {
			kepalaKeluarga = entity.DataAnggota{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
			}
		} else {
			anggota = append(anggota, entity.DataAnggota{
				Id:            anggotaRel.IdAnggota,
				NamaLengkap:   anggotaRel.NamaLengkap,
				TanggalLahir:  anggotaRel.TanggalLahir,
				TanggalBaptis: anggotaRel.TanggalBaptis,
				Keterangan:    anggotaRel.Keterangan,
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
	}

	return dataKeluargaFinal, nil
}

func (repository *dataKeluargaRepositoryImpl) GetTotalKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.TotalKeluarga, error) {
	sqlScript := "SELECT COUNT(*) FROM data_keluarga"
	result, err := tx.Query(sqlScript)
	if err != nil {
		return entity.TotalKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to execute query")
	}
	defer result.Close()

	totalKeluarga := entity.TotalKeluarga{}
	if result.Next() {
		err := result.Scan(&totalKeluarga.Total)
		if err != nil {
			return entity.TotalKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to scan result")
		}
		return totalKeluarga, nil
	} else {
		return entity.TotalKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "No data found")
	}
}

func (repository *dataKeluargaRepositoryImpl) AddKeluarga(ctx *fiber.Ctx, tx *sql.Tx) (entity.DataKeluarga, error) {
	repositories := ctx.Locals("repositories").(Repositories)

	body := ctx.Body()
	fmt.Println(body)
	request := new(helper.AddKeluargaRequest)
	err := json.Unmarshal(body, request)
	if err != nil {
		return entity.DataKeluarga{}, fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	//first add data_anggota first for kepala keluarga
	kepalaKeluarga, err := repositories.DataAnggotaRepository.AddAnggota(ctx, tx)
	if err != nil {
		return entity.DataKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to insert data kepala keluarga")
	}

	//add data_keluarga
	sqlScript := "INSERT INTO data_keluarga(id_wilayah, id_lingkungan, nomor, id_kepala_keluarga, alamat) VALUES(?, ?, ?, ?, ?)"
	result, err := tx.Exec(sqlScript, request.IdWilayah, request.IdLingkungan, request.Nomor, kepalaKeluarga.Id, request.Alamat)
	if err != nil {
		return entity.DataKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to insert data keluarga")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return entity.DataKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve last insert ID")
	}

	getLingkungan, err := repositories.DataLingkunganRepository.FindOneById(request.IdLingkungan, tx)
	if err != nil {
		return entity.DataKeluarga{}, fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve lingkungan data")
	}

	wilayah := getLingkungan.Wilayah
	lingkungan := entity.DataLingkungan{
		Id:             getLingkungan.Id,
		KodeLingkungan: getLingkungan.KodeLingkungan,
		NamaLingkungan: getLingkungan.NamaLingkungan,
		Wilayah:        wilayah,
	}

	// todo harusnya insert lingkungan dulu baru wilayah ya
	newDataKeluarga := entity.DataKeluarga{
		Id:             int32(lastInsertId),
		Wilayah:        wilayah,
		Lingkungan:     lingkungan,
		Nomor:          request.Nomor,
		KepalaKeluarga: kepalaKeluarga,
		Alamat:         request.Alamat,
	}

	return newDataKeluarga, nil
}

