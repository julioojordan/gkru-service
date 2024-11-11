package entity

type LoginResponse struct {
	Auth            string `json:"auth"`
	Id              int32  `json:"id"`
	Username        string `json:"username"`
	KetuaLingkungan int32  `json:"ketuaLingkungan"`
	KetuaWilayah    int32  `json:"ketuaWilayah"`
}
