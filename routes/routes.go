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
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/users", presenter.UserPresenter.GetAll)
	e.POST("/users", presenter.UserPresenter.PostUser)
	e.POST("/login", presenter.UserPresenter.LoginAuth)
	e.GET("/users/:id", presenter.UserPresenter.GetById, _middleware.JWTMiddleware())
	return e
}
