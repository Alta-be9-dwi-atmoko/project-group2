package business

import (
	"errors"
	"project/group2/features/products"
)

type productUseCase struct {
	productData products.Data
}

func NewProductBusiness(dataProduct products.Data) products.Business {
	return &productUseCase{
		productData: dataProduct,
	}
}

func (uc *productUseCase) GetAllData(limit, offset int) (resp []products.Core, err error) {
	resp, err = uc.productData.SelectData(limit, offset)
	return resp, err
}

func (uc *productUseCase) CreateData(input products.Core) (row int, err error) {
	if input.Name == "" || input.Price == 0 || input.Qty == 0 || input.Image == "" || input.Description == "" {
		return -1, errors.New("please make sure all fields are filled in correctly")
	}
	row, err = uc.productData.InsertData(input)
	return row, err
}

func (uc *productUseCase) GetProductById(idProd int) (data products.Core, err error) {
	data, err = uc.productData.SelectDataById(idProd)
	return data, err
}

func (uc *productUseCase) UpdateData(input products.Core, idProd int) (row int, err error) {
	prodReq := map[string]interface{}{}
	if input.Name != "" {
		prodReq["name"] = input.Name
	}
	if input.Price != 0 {
		prodReq["price"] = input.Price
	}
	if input.Qty != 0 {
		prodReq["qty"] = input.Qty
	}
	if input.Image != "" {
		prodReq["image"] = input.Image
	}
	if input.Description != "" {
		prodReq["description"] = input.Description
	}
	row, err = uc.productData.UpdateDataDB(prodReq, idProd)
	return row, err
}
