package helper

import (
	"gkru-service/entity"
	"time"
)

type FindOneRequest struct {
	Id int32 `json:"id"`
}

type AddAnggotaRequest struct {
	NamaLengkap   string    `json:"namaLengkap"`
	TanggalLahir  time.Time `json:"tanggalLahir"`
	TanggalBabtis time.Time `json:"tanggalBabtis"`
	Keterangan    string    `json:"keterangan"`
	Status        string    `json:"status"`
	Hubungan      string    `json:"hubungan"`
	IdKeluarga    int32     `json:"idKeluarga"`
}
type LingkunganRequest struct {
	KodeLingkungan string `json:"kodeLingkungan"`
	NamaLingkungan string `json:"namaLingkungan"`
	Wilayah        int32  `json:"wilayah"`
}

type WilayahRequest struct {
	KodeWilayah string `json:"kodeWilayah"`
	NamaWilayah string `json:"namaWilayah"`
}
type DeleteAnggotaRequest struct {
	SelectedAnggota []entity.DataAnggotaComplete `json:"selectedAnggota"`
}

type AddKeluargaRequest struct {
	IdWilayah    int32  `json:"idWilayah"`
	IdLingkungan int32  `json:"idLingkungan"`
	Nomor        string `json:"nomor"`
	Alamat       string `json:"alamat"`
}

type UpdateKeluargaRequest struct {
	IdWilayah        int32  `json:"idWilayah"`
	IdLingkungan     int32  `json:"idLingkungan"`
	Nomor            string `json:"nomor"`
	Alamat           string `json:"alamat"`
	IdKepalaKeluarga int32  `json:"idKepalaKeluarga"`
	Keterangan       string `json:"keterangan"`
}

type UpdateAnggotaRequest struct {
	Id            int32     `json:"idAnggota"`
	NamaLengkap   string    `json:"namaLengkap"`
	TanggalLahir  time.Time `json:"tanggalLahir"`
	TanggalBabtis time.Time `json:"tanggalBabtis"`
	Keterangan    string    `json:"keterangan"`
	Status        string    `json:"status"`
	Hubungan      string    `json:"hubungan"`
	IdKeluarga    int32     `json:"idKeluarga"`
}
type UpdateKeteranganAnggotaRequest struct {
	Id         int32  `json:"idKepalaKeluarga"`
	Keterangan string `json:"keterangan"`
}
