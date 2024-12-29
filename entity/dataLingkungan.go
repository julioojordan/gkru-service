package entity

type DataLingkunganRaw struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	IdWilayah      int32
	KodeWilayah    string
	NamaWilayah    string
}

type DataLingkunganRawWithTotalKeluarga struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	IdWilayah      int32
	KodeWilayah    string
	NamaWilayah    string
	TotalKeluarga  int32
}

type DataLingkungan struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	Wilayah        DataWilayah
}

type DataLingkunganWithTotalKeluarga struct {
	Id             int32
	KodeLingkungan string
	NamaLingkungan string
	Wilayah        DataWilayah
	TotalKeluarga  int32
}

type IdDataLingkungan struct {
	Id int32
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
