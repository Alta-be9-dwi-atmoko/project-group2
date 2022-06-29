package data

import (
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
