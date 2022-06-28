package data

import (
	"errors"
	_cart "project/group2/features/carts"

	"gorm.io/gorm"
)

type mysqlCartRepository struct {
	DB *gorm.DB
}

func NewCartRepository(db *gorm.DB) _cart.Data {
	return &mysqlCartRepository{
		DB: db,
	}
}

func (repo *mysqlCartRepository) SelectData(limit, offset, idFromToken int) (data []_cart.Core, err error) {
	dataCart := []Cart{}
	result := repo.DB.Find(&dataCart).Where("user_id=?", idFromToken)
	if result.Error != nil {
		return []_cart.Core{}, result.Error
	}
	return toCoreList(dataCart), nil
}

func (repo *mysqlCartRepository) InsertData(data _cart.Core) (row int, err error) {
	cart := FromCore(data)
	result := repo.DB.Create(&cart)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to create data cart")
	}
	return int(result.RowsAffected), nil
}
