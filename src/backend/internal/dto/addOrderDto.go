package dto

type AddOrderDto struct {
	ProductIds []int64 `json:"productIds"`
	UserId     int64   `json:"userId"`
}
