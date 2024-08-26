package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	body := ctx.Body()
	request := new(helper.FindOneRequest)
	err := json.Unmarshal(body, request)
	fmt.Println(request)
	helper.PanicIfError(err)

	result, err := tx.Query(dataKeluargaRawScript, request.Id)
	helper.PanicIfError(err)

	dataKeluargaRaw := entity.DataKeluargaRaw{}

	if result.Next() {
		err := result.Scan(&dataKeluargaRaw.Id, &dataKeluargaRaw.Wilayah, &dataKeluargaRaw.Lingkungan, &dataKeluargaRaw.Nomor, &dataKeluargaRaw.KepalaKeluarga, &dataKeluargaRaw.KKRelation, &dataKeluargaRaw.Alamat)
		helper.PanicIfError(err)
	} else {
		return entity.DataKeluargaFinal{}, fiber.NewError(fiber.StatusNotFound, "Data Keluarga is not found")
	}
	fmt.Println("dataKeluargaRaw", dataKeluargaRaw)
	fmt.Println("Masuk Sini 1")
	result.Close()

	repositories := ctx.Locals("repositories").(Repositories)
	
	fmt.Println("Masuk Sini 2")
	getLingkungan, err := repositories.DataLingkunganRepository.FindOneById(dataKeluargaRaw.Lingkungan, tx)
	
	fmt.Println("Masuk Sini 3")
	helper.PanicIfError(err)
	lingkungan := entity.DataLingkungan{
		Id:             getLingkungan.Id,
		KodeLingkungan: getLingkungan.KodeLingkungan,
		NamaLingkungan: getLingkungan.NamaLingkungan,
	}
	
	fmt.Println("lingkungan", lingkungan)
	wilayah := getLingkungan.Wilayah
	fmt.Println("wilayah", wilayah)

	getAnggotaRel, err := repositories.DataAnggotaKeluargaRelRepository.FindKeluargaAnggotaRel(dataKeluargaRaw.KKRelation, tx)
	fmt.Println("getAnggotaRel",getAnggotaRel)
	fmt.Println("getAnggotaRel err", err)
	helper.PanicIfError(err)

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
	result, err :=tx.Query(sqlScript)
	helper.PanicIfError(err);
	defer result.Close()
	
	totalKeluarga := entity.TotalKeluarga{}
	if result.Next(){
		err := result.Scan(&totalKeluarga.Total)
		helper.PanicIfError(err)
		return totalKeluarga, nil
	} else{
		return totalKeluarga, fiber.NewError(fiber.StatusInternalServerError , "Error Internal")
	}
}
