package dto

type CartDto struct {
	Id      IdObject   `json:"id"`
	Product ProductDto `json:"product"`
	UserId  IdObject   `json:"userId"`
}
