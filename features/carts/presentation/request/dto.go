package request

type Cart struct {
	Qty int `json:"qty" form:"qty"`
}
