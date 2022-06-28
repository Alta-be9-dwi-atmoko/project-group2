package response

import _cart "project/group2/features/carts"

type Cart struct {
	ID         int `json:"id"`
	Qty        int `json:"qty"`
	TotalPrice int `json:"total_price"`
	UserID     int `json:"user_id"`
	Product    Product
}

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Qty         int    `json:"qty"`
	Description string `json:"description"`
}

func FromCore(data _cart.Core) Cart {
	return Cart{
		ID:         data.ID,
		Qty:        data.Qty,
		TotalPrice: data.TotalPrice,
		UserID:     data.UserID,
		Product: Product{
			ID:          data.Product.ID,
			Name:        data.Product.Name,
			Price:       data.Product.Price,
			Qty:         data.Product.Qty,
			Description: data.Product.Description,
		},
	}
}

func FromCoreList(data []_cart.Core) []Cart {
	result := []Cart{}
	for key := range data {
		result = append(result, FromCore(data[key]))
	}
	return result
}
