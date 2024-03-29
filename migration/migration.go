package migration

import (
	_mCarts "project/group2/features/carts/data"
	_mOrders "project/group2/features/orders/data"
	_mProducts "project/group2/features/products/data"
	_mUsers "project/group2/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mProducts.Product{})
	db.AutoMigrate(&_mCarts.Cart{})
	db.AutoMigrate(&_mOrders.Order{})
	db.AutoMigrate(&_mOrders.OrderDetail{})
	db.AutoMigrate(&_mOrders.Address{})
	db.AutoMigrate(&_mOrders.Payment{})
}
