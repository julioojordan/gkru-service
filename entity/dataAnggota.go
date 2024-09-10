package entity

import "time"

type DataAnggota struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
}

type IdDataAnggota struct {
	Id int32
}

type DataAnggotaWithStatus struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
	Status        string
}

type DataAnggotaWithKeteranganOnly struct {
	Id         int32
	Keterangan string
}

type TotalAnggota struct {
	Total int32
}

type DataAnggotaResponse struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
}

type DataAnggotaComplete struct {
	Id             int32
	NamaLengkap    string
	TanggalLahir   time.Time
	TanggalBaptis  time.Time
	Keterangan     string
	Status         string
	IdKeluarga     int32
	IdWilayah      int32
	IdLingkungan   int32
	KodeLingkungan string
	NamaLingkungan string
	KodeWilayah    string
	NamaWilayah    string
}
