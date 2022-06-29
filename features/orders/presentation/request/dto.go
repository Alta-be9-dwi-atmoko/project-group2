package request

type Order struct {
	CartID  []int
	Address Address
	Payment Payment
}

type Address struct {
	Street     string `json:"street" form:"street"`
	City       string `json:"city" form:"city"`
	Province   string `json:"province" form:"province"`
	PostalCode string `json:"postal_code" form:"postal_code"`
}

type Payment struct {
	PaymentName string `json:"payment_name" form:"payment_name"`
	CardNumber  string `json:"card_number" form:"card_number"`
	PaymentCode string `json:"payment_code" form:"payment_code"`
}
