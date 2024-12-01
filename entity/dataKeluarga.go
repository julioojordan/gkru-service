package entity

type DataKeluargaRaw struct {
	Id             int32
	Wilayah        int32
	Lingkungan     int32
	Nomor          string
	KepalaKeluarga int32
	Alamat         string
	Status         string
}

type DataKeluarga struct {
	Id             int32
	Wilayah        DataWilayah
	Lingkungan     DataLingkungan
	Nomor          string
	KepalaKeluarga DataAnggota
	Alamat         string
}

type DeletedDataKeluarga struct {
	Id                int32
	DeletedAnggotaIds []int32
}

type DataKeluargaFinal struct {
	Id             int32
	Wilayah        DataWilayah
	Lingkungan     DataLingkungan
	Nomor          string
	KepalaKeluarga DataAnggotaWithStatus
	Alamat         string
	Anggota        []DataAnggotaWithStatus
	Status         string
}

type UpdatedDataKeluarga struct {
	Id             int32
	IdWilayah      int32
	IdLingkungan   int32
	Nomor          string
	Alamat         string
	KepalaKeluarga DataAnggotaWithStatus
	Anggota        []DataAnggotaWithStatus
	Status         string
}

type TotalKeluarga struct {
	Total int32
}

// cara pakai
/*
var keluarga DataKeluargaFinal

// Inisialisasi keluarga
keluarga := DataKeluargaFinal{
    Id:             1,
    Wilayah:        "Nama Wilayah",
    Lingkungan:     "Nama Lingkungan",
    Nomor:          123,
    KepalaKeluarga: "Nama Kepala Keluarga",
    Alamat:         "Alamat Keluarga",
}

// Menambahkan anggota ke keluarga
anggota1 := DataAnggota{Id: 1, NamaLengkap: "Anggota 1", TanggalLahir: time.Now(), TanggalBaptis: time.Now(), Keterangan: "Keterangan Anggota 1"}
anggota2 := DataAnggota{Id: 2, NamaLengkap: "Anggota 2", TanggalLahir: time.Now(), TanggalBaptis: time.Now(), Keterangan: "Keterangan Anggota 2"}

keluarga.Anggota = append(keluarga.Anggota, anggota1, anggota2)

*/
