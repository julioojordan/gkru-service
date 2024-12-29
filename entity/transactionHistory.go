package entity

import (
	"database/sql"
	"time"
)

type AmountHistory struct {
	Nominal int32
}

type CreatedTh struct {
	Id            int32
	Nominal       int32
	IdKeluarga    int32
	Keterangan    string
	CreatorId     int32
	SubKeterangan string
	CreatedDate   time.Time
	Group         int64
}

type ThRaw struct {
	Id             int32
	Nominal        int32
	IdKeluarga     int32
	Keterangan     string
	CreatorId      int32
	IdWilayah      int32
	IdLingkungan   int32
	UpdatorId      sql.NullInt32
	SubKeterangan  sql.NullString
	CreatedDate    time.Time
	UpdatedDate    time.Time
	Bulan          int32
	Tahun          int32
	UserName       string
	KodeLingkungan string
	NamaLingkungan string
	KodeWilayah    string
	NamaWilayah    string
	GroupId        int32
	File           sql.NullString
}

// pakai nama kepala keluarga
type ThRaw2 struct {
	Id                 int32
	Nominal            int32
	IdKeluarga         int32
	Keterangan         string
	CreatorId          int32
	IdWilayah          int32
	IdLingkungan       int32
	UpdatorId          sql.NullInt32
	SubKeterangan      sql.NullString
	CreatedDate        time.Time
	UpdatedDate        time.Time
	Bulan              int32
	Tahun              int32
	UserName           string
	KodeLingkungan     string
	NamaLingkungan     string
	KodeWilayah        string
	NamaWilayah        string
	GroupId            int32
	NamaKepalaKeluarga string
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
	Bulan         int32
	Tahun         int32
	GroupId       int32
	File          string
}

// th final with nama kepala keluarga
type ThFinal2 struct {
	Id                 int32
	Nominal            int32
	IdKeluarga         int32
	Keterangan         string
	Creator            User
	Wilayah            DataWilayah
	Lingkungan         DataLingkunganWithIdWilayah
	UpdatorId          int32
	SubKeterangan      string
	CreatedDate        time.Time
	UpdatedDate        time.Time
	Bulan              int32
	Tahun              int32
	GroupId            int32
	NamaKepalaKeluarga string
}

type UpdatedThFinal struct {
	Id            int32
	Nominal       int32
	IdKeluarga    int32
	Keterangan    string
	UpdatorId     int32
	SubKeterangan string
	UpdatedDate   time.Time
}
