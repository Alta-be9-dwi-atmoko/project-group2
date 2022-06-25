package business

import (
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
