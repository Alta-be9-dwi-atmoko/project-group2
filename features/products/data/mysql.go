package data

import (
	"errors"
	"project/group2/features/products"

	"gorm.io/gorm"
)

type mysqlProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) products.Data {
	return &mysqlProductRepository{
		DB: db,
	}
}

func (repo *mysqlProductRepository) SelectData(limit, offset int) (response []products.Core, err error) {
	dataProduct := []Product{}
	result := repo.DB.Limit(5).Offset(5).Find(&dataProduct)
	if result.Error != nil {
		return []products.Core{}, result.Error
	}
	return toCoreList(dataProduct), nil
}

func (repo *mysqlProductRepository) InsertData(input products.Core) (row int, err error) {
	product := FromCore(input)
	resultcreate := repo.DB.Create(&product)
	if resultcreate.Error == nil {
		return 0, resultcreate.Error
	}
	if resultcreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data")
	}
	return int(resultcreate.RowsAffected), nil
}

func (repo *mysqlProductRepository) SelectDataById(idProd int) (data products.Core, err error) {
	dataProduct := Product{}
	result := repo.DB.Find(&dataProduct, idProd)
	if result.Error != nil {
		return products.Core{}, result.Error
	}
	return dataProduct.toCore(), nil
}
