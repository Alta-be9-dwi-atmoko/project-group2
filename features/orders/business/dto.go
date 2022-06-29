package business

import "project/group2/features/orders"

func ToAddressCore(orderCore orders.Core) orders.AddressCore {
	return orders.AddressCore{
		City:       orderCore.Address.City,
		Province:   orderCore.Address.Province,
		PostalCode: orderCore.Address.PostalCode,
		Street:     orderCore.Address.Street,
	}
}

func ToPaymentCore(orderCore orders.Core) orders.PaymentCore {
	return orders.PaymentCore{
		PaymentName: orderCore.Payment.PaymentName,
		CardNumber:  orderCore.Payment.CardNumber,
		PaymentCode: orderCore.Payment.PaymentCode,
	}
}
