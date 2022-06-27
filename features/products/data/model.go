package data

import (
	"project/group2/features/products"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Price       int
	Qty         int
	Image       string
	Description string
	UserID      uint
	User        User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Product  []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// DTO

func (data *Product) toCore() products.Core {
	return products.Core{
		ID:          int(data.ID),
		Name:        data.Name,
		Price:       data.Price,
		Qty:         data.Qty,
		Image:       data.Image,
		Description: data.Description,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
		User: products.User{
			ID:   int(data.User.ID),
			Name: data.User.Name,
		},
	}
}

func toCoreList(data []Product) []products.Core {
	result := []products.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

func FromCore(core products.Core) Product {
	return Product{
		Name:        core.Name,
		Price:       core.Price,
		Qty:         core.Qty,
		Image:       core.Image,
		Description: core.Description,
		UserID:      uint(core.User.ID),
	}
}
