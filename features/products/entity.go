package products

import "time"

type Core struct {
	ID          int
	Name        string
	Price       int
	Qty         int
	Image       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User
}

type User struct {
	ID   int
	Name string
}

type Business interface {
	GetAllData(limit, offset int) (data []Core, err error)
	CreateData(data Core) (row int, err error)
	GetProductById(idProd int) (product Core, err error)
}

type Data interface {
	SelectData(limit, offset int) (data []Core, err error)
	InsertData(data Core) (row int, err error)
	SelectDataById(idProd int) (product Core, err error)
}
