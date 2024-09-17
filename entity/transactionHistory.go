package entity

import (
	"time"
)

type AmountHistory struct {
	Nominal int32
}

type CreatedTh struct {
	Id             int32
	Nominal        int32
	IdKeluarga     int32
	Keterangan     string
	CreatorId      int32
	IdWilayah      int32
	IdLingkungan   int32
	SubKeterangan  string
	CreatedDate    time.Time
}

type ThRaw struct {
	Id             int32
	Nominal        int32
	IdKeluarga     int32
	Keterangan     string
	CreatorId      int32
	IdWilayah      int32
	IdLingkungan   int32
	UpdatorId      int32
	SubKeterangan  string
	CreatedDate    time.Time
	UpdatedDate    time.Time
	UserName       string
	KodeLingkungan string
	NamaLingkungan string
	KodeWilayah    string
	NamaWilayah    string
}

type ThFinal struct {
	Id            int32
	Nominal       int32
	IdKeluarga    int32
	Keterangan    string
	Creator       User
	Wilayah       DataWilayah
	Lingkungan    DataLingkunganWithIdWilayah
	UpdatorId     int32
	SubKeterangan string
	CreatedDate   time.Time
	UpdatedDate   time.Time
}

type UpdatedThFinal struct {
	Id            int32
	Nominal       int32
	IdKeluarga    int32
	Keterangan    string
	IdWilayah     int32
	IdLingkungan  int32
	UpdatorId     int32
	SubKeterangan string
	UpdatedDate   time.Time
}
