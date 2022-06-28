package business

import (
	"errors"
	"fmt"
	_cart "project/group2/features/carts"
)

type cartUseCase struct {
	cartData _cart.Data
}

func NewCartBusiness(datacart _cart.Data) _cart.Business {
	return &cartUseCase{
		cartData: datacart,
	}
}

func (uc *cartUseCase) GetAllData(limit, offset, idFromToken int) (data []_cart.Core, err error) {
	data, err = uc.cartData.SelectData(limit, offset, idFromToken)
	fmt.Println("data:", data)
	for k, v := range data {
		data[k].TotalPrice = v.Qty * v.Product.Price
	}
	fmt.Println("data after:", data)
	return data, err
}

func (uc *cartUseCase) CreateData(data _cart.Core) (row int, err error) {
	if data.Qty == 0 || data.Product.ID == 0 {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.cartData.InsertData(data)
	return row, err
}
