package entity

import "time"

type DataAnggotaKeluargaRel struct {
	Id         int32
	IdKeluarga int32
	IdAnggota  int32
	Hubungan   string
}

type DataAnggotaWithKeluargaRel struct {
	Id            int32
	Hubungan      string
	IdAnggota     int32
	NamaLengkap   string
	TanggalLahir  time.Time
	TanggalBaptis time.Time
	Keterangan    string
}
