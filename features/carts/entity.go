package cart

import "time"

type Core struct {
	ID         int
	Qty        int
	TotalPrice int
	Status     string
	UserID     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Product    Product
}

type Product struct {
	ID          int
	Name        string
	Price       int
	Qty         int
	Description string
}

type Business struct{}

type Data struct{}
