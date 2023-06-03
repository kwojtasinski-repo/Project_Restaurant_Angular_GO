package dto

type CartDto struct {
	Id      int64      `json:"id"`
	Product ProductDto `json:"product"`
	UserId  int64      `json:"userId"`
}
