package helper

import (
	"gkru-service/customType"
	"gkru-service/entity"
	"time"
)

type FindOneRequest struct {
	Id int32 `json:"id"`
}

type AddAnggotaRequest struct {
	NamaLengkap   string `json:"namaLengkap"`
	TanggalLahir  string `json:"tanggalLahir"`
	TanggalBaptis string `json:"tanggalBaptis"`
	Keterangan    string `json:"keterangan"`
	Status        string `json:"status"`
	Hubungan      string `json:"hubungan"`
	IdKeluarga    int32  `json:"idKeluarga"`
	JenisKelamin  string `json:"jenisKelamin"`
	NoTelp        string `json:"noTelp"`
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

type UserRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	KetuaLingkungan int32  `json:"ketuaLingkungan"`
	KetuaWilayah    int32  `json:"ketuaWilayah"`
	UpdatedBy       int32  `json:"updatedBy"`
}

type AddUserRequest struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	KetuaLingkungan int32  `json:"ketuaLingkungan"`
	KetuaWilayah    int32  `json:"ketuaWilayah"`
	UpdatedBy       int32  `json:"updatedBy"`
	CreatedBy       int32  `json:"createdBy"`
}
type DeleteAnggotaRequest struct {
	SelectedAnggota []entity.DataAnggotaComplete `json:"selectedAnggota"`
}

type AddKeluargaRequest struct {
	IdWilayah     int32  `json:"idWilayah"`
	IdLingkungan  int32  `json:"idLingkungan"`
	Nomor         string `json:"nomor"`
	Alamat        string `json:"alamat"`
	NomorKKGereja string `json:"nomorKKGereja"`
}

type UpdateKeluargaRequest struct {
	IdWilayah           int32  `json:"idWilayah"`
	IdLingkungan        int32  `json:"idLingkungan"`
	Nomor               string `json:"nomor"`
	Alamat              string `json:"alamat"`
	Status              string `json:"status"`
	IdKepalaKeluarga    int32  `json:"idKepalaKeluarga"`
	OldIdKepalaKeluarga int32  `json:"OldIdKepalaKeluarga"`
	Keterangan          string `json:"keterangan"`
	NomorKKGereja       string `json:"nomorKKGereja"`
}

type UpdateTHRequest struct {
	Nominal       int32  `json:"nominal"`
	Keterangan    string `json:"keterangan"`
	IdKeluarga    int32  `json:"idKeluarga"`
	SubKeterangan string `json:"subKeterangan"`
	UpdatedBy     int32  `json:"UpdatedBy"`
}

type AddTHRequest struct {
	Nominal       int32     `json:"nominal"`
	IdKeluarga    int32     `json:"idKeluarga"`
	Keterangan    string    `json:"keterangan"`
	CreatedBy     int32     `json:"CreatedBy"`
	IdWilayah     int32     `json:"idWilayah"`
	IdLingkungan  int32     `json:"idLingkungan"`
	SubKeterangan string    `json:"subKeterangan"`
	CreatedDate   time.Time `json:"createdDate"`
	Bulan         int32     `json:"bulan"`
	Tahun         int32     `json:"tahun"`
}

type UpdateAnggotaRequest struct {
	Id               int32                 `json:"id"`
	NamaLengkap      string                `json:"namaLengkap"`
	TanggalLahir     customType.CustomTime `json:"tanggalLahir"`
	TanggalBaptis    customType.CustomTime `json:"tanggalBaptis"`
	Keterangan       string                `json:"keterangan"`
	Status           string                `json:"status"`
	Hubungan         string                `json:"hubungan"`
	IdKeluarga       int32                 `json:"idKeluarga"`
	JenisKelamin     string                `json:"jenisKelamin"`
	IsKepalaKeluarga bool                  `json:"isKepalaKeluarga"`
	NoTelp           string                `json:"noTelp"`
}
type UpdateKeteranganAnggotaRequest struct {
	Id         int32  `json:"idKepalaKeluarga"`
	Keterangan string `json:"keterangan"`
	OldId      int32  `json:"oldIdKepalaKeluarga"`
}
