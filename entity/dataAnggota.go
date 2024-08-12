package entity

import "time"

type DataAnggota struct {
	Id            int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
}
