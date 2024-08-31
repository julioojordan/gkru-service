package entity

type DataLingkunganRaw struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	IdWilayah      int32
	KodeWilayah    string
	NamaWilayah    string
}

type DataLingkungan struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	Wilayah        DataWilayah
}
