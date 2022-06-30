package data

import (
	"errors"
	"fmt"
	_cart "project/group2/features/carts/data"
	"project/group2/features/orders"

	"gorm.io/gorm"
)

type mysqlOrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) orders.Data {
	return &mysqlOrderRepository{
		DB: db,
	}
}

func (repo *mysqlOrderRepository) InsertAddress(address orders.AddressCore) (idAddress int, err error) {
	addressGorm := fromAddressCore(address)
	result := repo.DB.Create(&addressGorm)
	if result.Error != nil {
		return 0, result.Error
	}
	return addressGorm.ID, nil
}

func (repo *mysqlOrderRepository) InsertPayment(payment orders.PaymentCore) (idPayment int, err error) {
	paymentGorm := fromPaymentCore(payment)
	result := repo.DB.Create(&paymentGorm)
	if result.Error != nil {
		return 0, result.Error
	}
	return paymentGorm.ID, nil
}

func (repo *mysqlOrderRepository) InsertOrder(data orders.Core) (idOrder int, err error) {
	orderGorm := fromCore(data)
	result := repo.DB.Create(&orderGorm)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(orderGorm.ID), nil
}

func (repo *mysqlOrderRepository) TotalPrice(idCart []int) (totalPrice int, err error) {
	totalPriceCart := []_cart.Cart{}
	result := repo.DB.Preload("Product").Find(&totalPriceCart, idCart)
	if result.Error != nil {
		return -1, result.Error
	}
	for _, v := range totalPriceCart {
		totalPrice += (v.Qty * v.Product.Price)
	}
	//[]OrderDetail{}
	// v.Product.ID
	// v.Product.Price
	return totalPrice, nil
}

func (repo *mysqlOrderRepository) SelectCart(idCart []int, userID int) (item []orders.OrderDetail, err error) {
	resultCart := []_cart.Cart{}
	resultQuery := repo.DB.Preload("Product").Find(&resultCart, idCart)

	if resultQuery.Error != nil {
		return []orders.OrderDetail{}, resultQuery.Error
	}

	itemDetail := []orders.OrderDetail{}
	for i := 0; i < len(resultCart); i++ {
		itemDetail = append(itemDetail, orders.OrderDetail{
			Price:     resultCart[i].Qty * resultCart[i].Product.Price,
			ProductID: resultCart[i].ProductID,
			Qty:       resultCart[i].Qty,
		})

	}
	fmt.Println("item: ", itemDetail)
	return itemDetail, nil
}

func (repo *mysqlOrderRepository) InsertOrderDetail(data orders.OrderDetail, orderID int) (row int, err error) {
	dataOrder := fromOrderDetailCore(data)
	dataOrder.OrderID = orderID
	result := repo.DB.Create(&dataOrder)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to insert orderdetail")
	}
	return int(result.RowsAffected), nil
}