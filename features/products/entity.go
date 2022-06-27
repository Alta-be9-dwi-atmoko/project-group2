package products

type Core struct {
	ID          int
	Name        string
	Price       int
	Qty         int
	Image       string
	Description string
	User        User
}

type User struct {
	ID   int
	Name string
}

type Business interface{}

type Data interface{}
