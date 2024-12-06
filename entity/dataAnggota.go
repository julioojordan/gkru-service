package entity

import (
	"gkru-service/customType"
	"time"
)

type DataAnggota struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
	JenisKelamin  string
}

type IdDataAnggota struct {
	Id int32
}

type DataAnggotaWithStatus struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  customType.CustomTime
	TanggalBaptis customType.CustomTime
	Keterangan    string
	Status        string
	JenisKelamin  string
}

type DataAnggotaWithKeteranganOnly struct {
	Id         int32
	Keterangan string
}

type TotalAnggota struct {
	Total int32
}

type DataAnggotaComplete struct {
	Id             int32
	NamaLengkap    string
	TanggalLahir   time.Time
	TanggalBaptis  time.Time
	Keterangan     string
	Status         string
	JenisKelamin   string
	IdKeluarga     int32
	Hubungan       string
	IdWilayah      int32
	IdLingkungan   int32
	KodeLingkungan string
	NamaLingkungan string
	KodeWilayah    string
	NamaWilayah    string
}
