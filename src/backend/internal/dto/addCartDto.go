package dto

type AddCart struct {
	ProductId IdObject `json:"productId"`
	UserId    IdObject `json:"userId"`
}
