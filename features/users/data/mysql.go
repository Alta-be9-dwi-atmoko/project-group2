package data

import (
	"project/group2/features/users"

	"gorm.io/gorm"
)

type mysqlUserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.Data {
	return &mysqlUserRepository{
		DB: db,
	}
}

func (repo *mysqlUserRepository) SelectData(data string) (response []users.Core, err error) {
	dataUsers := []User{}
	result := repo.DB.Find(&dataUsers)
	if result.Error != nil {
		return []users.Core{}, result.Error
	}
	return toCoreList(dataUsers), nil
}
