package routes

import (
	"project/group2/factory"

	"github.com/labstack/echo/v4"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	e.GET("/users", presenter.UserPresenter.GetAll)
	e.POST("/users", presenter.UserPresenter.PostUser)
	return e
}
