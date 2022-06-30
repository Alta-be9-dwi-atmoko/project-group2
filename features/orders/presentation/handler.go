package presentation

import (
	"fmt"
	"net/http"
	_middleware "project/group2/features/middlewares"
	_order "project/group2/features/orders"
	_requestOrder "project/group2/features/orders/presentation/request"
	_helper "project/group2/helper"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderBusiness _order.Business
}

func NewOrderHandler(orderBusiness _order.Business) *OrderHandler {
	return &OrderHandler{
		orderBusiness: orderBusiness,
	}
}

func (h *OrderHandler) PostOrder(c echo.Context) error {
	orderReq := _requestOrder.Order{}
	idFromToken, _ := _middleware.ExtractToken(c)
	err := c.Bind(&orderReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	dataOrder := _requestOrder.ToCore(orderReq)
	dataOrder.UserID = idFromToken
	row, errCreate := h.orderBusiness.CreateData(dataOrder, orderReq.CartID)
	fmt.Println("row: ", row)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to create product, check your input"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
