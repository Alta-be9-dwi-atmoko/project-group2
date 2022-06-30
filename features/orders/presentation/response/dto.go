package response

import "time"

type Order struct {
	ID         int       `json:"id"`
	Qty        int       `json:"qty"`
	TotalPrice int       `json:"total_price"`
	UserID     int       `json:"user_id"`
	Product    Product   `json:"product"`
	CreatedAt  time.Time `json:"created_at"`
}

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}
