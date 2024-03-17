package entity

type AuthWebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Auth   string      `json:"auth"`
}
