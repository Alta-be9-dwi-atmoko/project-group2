package factory

import (
	_userBusiness "project/group2/features/users/business"
	_userData "project/group2/features/users/data"
	_userPresentation "project/group2/features/users/presentation"

	_productBusiness "project/group2/features/products/business"
	_productData "project/group2/features/products/data"
	_productPresentation "project/group2/features/products/presentation"

	_cartBusiness "project/group2/features/carts/business"
	_cartData "project/group2/features/carts/data"
	_cartPresentation "project/group2/features/carts/presentation"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    *_userPresentation.UserHandler
	ProductPresenter *_productPresentation.ProductHandler
	CartPresenter    *_cartPresentation.CartHandler
}

func InitFactory(dbConn *gorm.DB) Presenter {
	// dbConn := config.InitDB()
	userData := _userData.NewUserRepository(dbConn)
	userBusiness := _userBusiness.NewUserBusiness(userData)
	userPresentation := _userPresentation.NewUserHandler(userBusiness)

	productData := _productData.NewProductRepository(dbConn)
	productBusiness := _productBusiness.NewProductBusiness(productData)
	productPresentation := _productPresentation.NewProductHandler(productBusiness)

	cartData := _cartData.NewCartRepository(dbConn)
	cartBusiness := _cartBusiness.NewCartBusiness(cartData)
	cartPresentation := _cartPresentation.NewCartHandler(cartBusiness)

	return Presenter{
		UserPresenter:    userPresentation,
		ProductPresenter: productPresentation,
		CartPresenter:    cartPresentation,
	}
}
