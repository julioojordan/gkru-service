package entity

import (
	"gkru-service/customType"
)

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
	TanggalLahir  customType.CustomTime
	TanggalBaptis customType.CustomTime
	Keterangan    string
	Status        string
	JenisKelamin  string
}
