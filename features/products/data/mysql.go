package data

import (
	"errors"
	"fmt"
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
	result := repo.DB.Limit(limit).Offset(offset).Find(&dataProduct)
	if result.Error != nil {
		return []products.Core{}, result.Error
	}
	fmt.Println("dataProductbefore: ", dataProduct)
	fmt.Println("dataProductafter: ", toCoreList(dataProduct))
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

func (repo *mysqlProductRepository) UpdateDataDB(data map[string]interface{}, idProd, idFromToken int) (row int, err error) {
	dataProduct := Product{}
	idCheck := repo.DB.First(&dataProduct, idProd)
	if idCheck.Error != nil {
		return 0, idCheck.Error
	}
	if dataProduct.UserID != uint(idFromToken) {
		return -1, errors.New("you don't have access")
	}
	result := repo.DB.Model(&Product{}).Where("id = ?", idProd).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}

	return int(result.RowsAffected), nil
}

func (repo *mysqlProductRepository) GetDataByMeDB(idUser int) (data []products.Core, err error) {
	dataProduct := []Product{}
	result := repo.DB.Find(&dataProduct).Where("userid = ?", idUser)
	if result.Error != nil {
		return []products.Core{}, result.Error
	}
	return toCoreList(dataProduct), nil
}

func (repo *mysqlProductRepository) DeleteDataByIdDB(idProd, idFromToken int) (row int, err error) {
	dataProduct := Product{}
	idCheck := repo.DB.First(&dataProduct, idProd)
	if idCheck.Error != nil {
		return 0, idCheck.Error

	}
	fmt.Println("idUser: ", dataProduct.UserID)
	fmt.Println("idFromToken: ", idFromToken)
	if idFromToken != int(dataProduct.UserID) {
		return -1, errors.New("you don't have access")
	}

	result := repo.DB.Delete(&User{}, idProd)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected != 1 {
		return 0, errors.New("failed to delete data")
	}
	return int(result.RowsAffected), nil
}
