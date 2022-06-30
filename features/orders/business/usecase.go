package business

import (
	"errors"
	"fmt"
	"project/group2/features/orders"
)

type orderUseCase struct {
	orderData orders.Data
}

func NewOrderBusiness(dataorder orders.Data) orders.Business {
	return &orderUseCase{
		orderData: dataorder,
	}
}

func (uc *orderUseCase) CreateData(data orders.Core, idCart []int) (row int, err error) {

	if data.Address.City == "" || data.Address.Street == "" || data.Address.PostalCode == "" || data.Address.Province == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	idAddress, errAddress := uc.orderData.InsertAddress(data.Address)
	if errAddress != nil {
		return 0, errAddress
	}

	if data.Payment.PaymentName == "" || data.Payment.CardNumber == "" || data.Payment.PaymentCode == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	idPayment, errPayment := uc.orderData.InsertPayment(data.Payment)
	if errPayment != nil {
		return 0, errPayment
	}
	data.PaymentID = idPayment
	data.AddressID = idAddress
	data.Status = "Pending"
	totalPrice, errTotalPrice := uc.orderData.TotalPrice(idCart)
	if errTotalPrice != nil {
		return 0, errTotalPrice
	}
	data.TotalPrice = totalPrice
	idOrder, errOrder := uc.orderData.InsertOrder(data)
	if errOrder != nil {
		return 0, errOrder
	}
	resultCart, errCart := uc.orderData.SelectCart(idCart, data.UserID)
	if errCart != nil {
		return 0, errCart
	}
	for i := 0; i < len(resultCart); i++ {
		fmt.Println("result: ", resultCart[i])
		row, err = uc.orderData.InsertOrderDetail(resultCart[i], idOrder)
	}
	return row, err
}

func (uc *orderUseCase) ConfirmStatus(orderID, idFromToken int) (row int, err error) {
	row, err = uc.orderData.ConfirmStatusData(orderID, idFromToken)
	return row, err
}
