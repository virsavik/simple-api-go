package model

type Paginate struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}
