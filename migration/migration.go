package migration

import (
	_mProducts "project/group2/features/products/data"
	_mUsers "project/group2/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_mUsers.User{})
	db.AutoMigrate(&_mProducts.Product{})
}
