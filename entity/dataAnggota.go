package entity

import "time"

type DataAnggota struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
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
	Id            int32
	Keterangan    string
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
