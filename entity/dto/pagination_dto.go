package dto

type GetAllParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}