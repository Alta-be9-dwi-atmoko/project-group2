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
			Price:       resultCart[i].Qty * resultCart[i].Product.Price,
			ProductID:   resultCart[i].ProductID,
			ProductName: resultCart[i].Product.Name,
			Qty:         resultCart[i].Qty,
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

func (repo *mysqlOrderRepository) ConfirmStatusData(orderID, idFromToken int) (row int, err error) {
	dataOrder := Order{}
	idCheck := repo.DB.First(&dataOrder, orderID)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataOrder.UserID != idFromToken {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Model(&Order{}).Where("id = ?", orderID).Update("status", "Confirmed")
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update order status")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlOrderRepository) CancelStatusData(orderID, idFromToken int) (row int, err error) {
	dataOrder := Order{}
	idCheck := repo.DB.First(&dataOrder, orderID)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataOrder.UserID != idFromToken {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Model(&Order{}).Where("id = ?", orderID).Update("status", "Cancelled")
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update order status")
	}
	return int(result.RowsAffected), nil
}

func (repo *mysqlOrderRepository) HistoryAllData(limitint, offsetint, idFromToken int) (data []orders.Core, err error) {
	dataOrder := []Order{}
	result := repo.DB.Limit(limitint).Offset(offsetint).Preload("OrderDetail").Preload("Address").Preload("Payment").Where("user_id = ? AND status = ? OR status = ?", idFromToken, "Confirmed", "Cancelled").Find(&dataOrder)
	if result.Error != nil {
		return []orders.Core{}, result.Error
	}
	return ToCoreList(dataOrder), nil
}

func (repo *mysqlOrderRepository) OrderDetailDB(orderID int) (data []orders.OrderDetail, err error) {
	dataOrderdetail := []OrderDetail{}
	result := repo.DB.Where("order_id = ?", orderID).Find(&dataOrderdetail)
	if result.Error != nil {
		return []orders.OrderDetail{}, result.Error
	}
	return ToOrderDetailCoreList(dataOrderdetail), nil
}
