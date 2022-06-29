package data

import (
	"fmt"
	_mProduct "project/group2/features/carts/data"
	"project/group2/features/orders"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PaymentID int
	UserID    int
	Price     int
	AddressID int
	Status    string
	Address   Address
	Payment   Payment
}

type OrderDetail struct {
	ID          int `gorm:"autoIncrement"`
	OrderID     int
	ProductName string
	Price       int
	Qty         int
}

type Cart struct {
	gorm.Model
	Qty       int
	Status    string
	ProductID int
	UserID    int
	Product   _mProduct.Product
}

type Address struct {
	ID         int `gorm:"autoIncrement"`
	City       string
	Province   string
	PostalCode string
	Street     string
}

type Payment struct {
	ID          int `gorm:"autoIncrement"`
	PaymentName string
	CardNumber  string
	PaymentCode string
}

func (data Order) toCore() orders.Core {
	return orders.Core{
		ID:     int(data.ID),
		Price:  data.Price,
		UserID: data.UserID,
	}
}

func ToCoreList(data []Order) []orders.Core {
	result := []orders.Core{}
	for key := range data {
		result = append(result, data[key].toCore())
	}
	return result
}

// func (dataOrderDetail) toOrderDetailCore() orders.OrderDetail {
// 	return orders.OrderDetail{
// 		ID:      int(data.ID),
// 		OrderID: data.OrderID,
// 		Price:   data.Price,
// 		Qty:     data.Qty,
// 	}
// }
// func ToOrderDetailCoreList(data []OrderDetail) []orders.OrderDetail {
// 	result := []orders.OrderDetail{}
// 	for key := range data {
// 		result = append(result, data[key].toOrderDetailCore())
// 	}
// 	return result
// }

func (data *Cart) toOrderDetailFromCart() orders.OrderDetail {
	return orders.OrderDetail{
		ID:          int(data.ID),
		Qty:         data.Qty,
		ProductName: data.Product.Name,
		Price:       data.Product.Price,
	}
}

func ToOrderDetailCoreListFromCart(data []Cart) []orders.OrderDetail {
	result := []orders.OrderDetail{}
	for key := range data {
		result = append(result, data[key].toOrderDetailFromCart())
	}
	return result
}

func fromCore(core orders.Core) Order {
	return Order{
		PaymentID: core.PaymentID,
		AddressID: core.AddressID,
		Price:     core.Price,
		UserID:    core.UserID,
		Status:    core.Status,
	}
}

func fromOrderDetailCore(orderDetailCore orders.OrderDetail) OrderDetail {
	return OrderDetail{
		ProductName: orderDetailCore.ProductName,
		Price:       orderDetailCore.Price,
		OrderID:     orderDetailCore.OrderID,
		Qty:         orderDetailCore.Qty,
	}
}

func fromOrderDetailCoreList(data []orders.OrderDetail) []OrderDetail {
	result := []OrderDetail{}
	for key := range data {
		result = append(result, fromOrderDetailCore(data[key]))
	}
	fmt.Println(result)
	fmt.Println(data)
	return result
}

func fromAddressCore(addressCore orders.AddressCore) Address {
	return Address{
		City:       addressCore.City,
		Province:   addressCore.Province,
		PostalCode: addressCore.PostalCode,
		Street:     addressCore.Street,
	}
}

func fromPaymentCore(paymentCore orders.PaymentCore) Payment {
	return Payment{
		CardNumber:  paymentCore.CardNumber,
		PaymentCode: paymentCore.PaymentCode,
		PaymentName: paymentCore.PaymentName,
	}
}
