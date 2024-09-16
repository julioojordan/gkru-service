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

type IdDataLingkungan struct {
	Id             int32
}

type TotalLingkungan struct {
	Total int32
}
type DataLingkunganWithIdWilayah struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	Wilayah        int32
}
