package routes

import (
	"project/group2/factory"
	_middleware "project/group2/features/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/users", presenter.UserPresenter.GetAll)
	e.POST("/users", presenter.UserPresenter.PostUser)
	e.POST("/login", presenter.UserPresenter.LoginAuth)
	e.GET("/users/:id", presenter.UserPresenter.GetById, _middleware.JWTMiddleware())
	e.PUT("/users/:id", presenter.UserPresenter.PutUser, _middleware.JWTMiddleware())
	e.DELETE("/users/:id", presenter.UserPresenter.DeleteById, _middleware.JWTMiddleware())

	e.GET("/products", presenter.ProductPresenter.GetAll)
	e.POST("/products", presenter.ProductPresenter.PostProduct, _middleware.JWTMiddleware())
	e.GET("/products/:id", presenter.ProductPresenter.GetById)
	e.PUT("/products/:id", presenter.ProductPresenter.PutProduct, _middleware.JWTMiddleware())
	e.GET("/products/me", presenter.ProductPresenter.GetByMe, _middleware.JWTMiddleware())
	e.DELETE("/products/:id", presenter.ProductPresenter.DeleteById, _middleware.JWTMiddleware())

	e.POST("/carts", presenter.CartPresenter.PostCart, _middleware.JWTMiddleware())
	e.GET("/carts", presenter.CartPresenter.GetAll, _middleware.JWTMiddleware())
	e.PUT("/carts/:id", presenter.CartPresenter.UpdateCart, _middleware.JWTMiddleware())
	e.DELETE("/carts/:id", presenter.CartPresenter.DeleteCart, _middleware.JWTMiddleware())

	e.POST("/orders", presenter.OrderPresenter.PostOrder, _middleware.JWTMiddleware())
	e.PUT("/orders/:id/confirm", presenter.OrderPresenter.Confirmed, _middleware.JWTMiddleware())
	e.PUT("/orders/:id/cancel", presenter.OrderPresenter.Cancelled, _middleware.JWTMiddleware())
	e.GET("/orders/history", presenter.OrderPresenter.History, _middleware.JWTMiddleware())
	e.GET("/orderdetail/:id", presenter.OrderPresenter.OrderDetail)
	e.GET("/orders/me", presenter.OrderPresenter.GetMyOrder, _middleware.JWTMiddleware())
	return e

}
