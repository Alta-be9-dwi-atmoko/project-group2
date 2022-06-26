package business

import (
	"errors"
	"project/group2/features/users"
)

type userUseCase struct {
	userData users.Data
}

func NewUserBusiness(dataUser users.Data) users.Business {
	return &userUseCase{
		userData: dataUser,
	}
}

func (uc *userUseCase) GetAllData(limit, offset int) (resp []users.Core, err error) {
	var param string
	resp, err = uc.userData.SelectData(param)
	return resp, err
}

func (uc *userUseCase) CreateData(input users.Core) (row int, err error) {
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.userData.InsertData(input)
	return row, err
}

func (uc *userUseCase) LoginUser(authData users.AuthRequestData) (token, name string, err error) {
	token, name, err = uc.userData.LoginUserDB(authData)
	return token, name, err
}

func (uc *userUseCase) GetUserById(idUser int) (data users.Core, err error) {
	data, err = uc.userData.SelectDataById(idUser)
	return data, err
}

func (uc *userUseCase) UpdateData(input users.Core, idUser int) (row int, err error) {
	userReq := map[string]interface{}{}
	if input.Name != "" {
		userReq["name"] = input.Name
	}
	if input.Email != "" {
		userReq["email"] = input.Email
	}
	if input.Password != "" {
		userReq["password"] = input.Password
	}
	row, err = uc.userData.UpdateDataDB(userReq, idUser)
	return row, err
}
