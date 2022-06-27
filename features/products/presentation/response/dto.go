package response

import (
	"project/group2/features/products"
)

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Image       string `json:"image"`
	Description string `json:"description"`
	User        User
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func FromCore(data products.Core) Product {
	return Product{
		ID:          data.ID,
		Name:        data.Name,
		Price:       data.Price,
		Qty:         data.Qty,
		Image:       data.Image,
		Description: data.Description,
		User: User{
			ID:   data.User.ID,
			Name: data.User.Name,
		},
	}
}
