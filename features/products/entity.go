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

type Business interface{}

type Data interface{}
