package data

import (
	"errors"
	"fmt"
	"project/group2/features/middlewares"
	"project/group2/features/users"

	_bcrypt "golang.org/x/crypto/bcrypt"
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

func (repo *mysqlUserRepository) InsertData(input users.Core) (row int, err error) {
	user := FromCore(input)
	//email duplicates check
	// resultemail := repo.DB.Model(&User{}).Where("email = ?", user.Email)
	// if resultemail.Error != nil {
	// 	return 0, errors.New("email is already registered")
	// }
	passwordHashed, errorHash := _bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if errorHash != nil {
		fmt.Println("Error hash", errorHash.Error())
	}
	user.Password = string(passwordHashed)
	resultcreate := repo.DB.Create(&user)
	if resultcreate.Error == nil {
		return 0, resultcreate.Error
	}
	if resultcreate.RowsAffected != 1 {
		return 0, errors.New("failed to insert data, your email is already registered")
	}
	return int(resultcreate.RowsAffected), nil
}

func (repo *mysqlUserRepository) LoginUserDB(authData users.AuthRequestData) (token string, name string, err error) {
	userData := User{}
	result := repo.DB.Where("email = ?", authData.Email).First(&userData)
	if result.Error != nil {
		return "", "", result.Error
	}

	if result.RowsAffected != 1 {
		return "", "", errors.New("failed to login")
	}

	errCrypt := _bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(authData.Password))
	if errCrypt != nil {
		return "", "", errors.New("password incorrect")
	}

	token, errToken := middlewares.CreateToken(int(userData.ID))
	if errToken != nil {
		return "", "", errToken
	}

	return token, userData.Name, nil
}

func (repo *mysqlUserRepository) SelectDataById(idUser int) (data users.Core, err error) {
	dataUser := User{}
	result := repo.DB.Find(&dataUser, idUser)
	if result.Error != nil {
		return users.Core{}, result.Error
	}
	return dataUser.toCore(), nil
}

func (repo *mysqlUserRepository) UpdateDataDB(data map[string]interface{}, idUser int) (row int, err error) {
	result := repo.DB.Model(&User{}).Where("id = ?", idUser).Updates(data)
	if result.Error != nil {
		return 0, result.Error
	}

	if result.RowsAffected != 1 {
		return 0, errors.New("failed to update data")
	}

	return int(result.RowsAffected), nil
}
