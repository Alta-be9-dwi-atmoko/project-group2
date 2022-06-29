package business

import "project/group2/features/orders"

type orderUseCase struct {
	orderData orders.Data
}

func NewOrderBusiness(dataorder orders.Data) orders.Business {
	return &orderUseCase{
		orderData: dataorder,
	}
}

func (uc *orderUseCase) CreateData(data orders.Core, idCart []int) (row int, err error) {
	addressCore := ToAddressCore(data)
	idAddress, _ := uc.orderData.InsertAddress(addressCore)

}
