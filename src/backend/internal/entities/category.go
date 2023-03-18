package entities

type Category struct {
	Id      int64
	Name    string
	Deleted bool // soft delete
}
