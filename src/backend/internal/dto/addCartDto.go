package dto

type AddCart struct {
	ProductId int64 `json:"productId"`
	UserId    int64 `json:"userId"`
}
