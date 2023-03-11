package entities

type Cart struct {
	Id        int64
	UserId    int64
	User      User
	ProductId int64
	Product   Product
}
