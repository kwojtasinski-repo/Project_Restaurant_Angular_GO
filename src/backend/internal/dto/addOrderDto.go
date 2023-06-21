package dto

type AddOrderDto struct {
	ProductIds []IdObject `json:"productIds"`
	UserId     IdObject   `json:"userId"`
}
