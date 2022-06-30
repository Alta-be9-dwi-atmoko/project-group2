package presentation

import (
	"fmt"
	"net/http"
	_middleware "project/group2/features/middlewares"
	_order "project/group2/features/orders"
	_requestOrder "project/group2/features/orders/presentation/request"
	_responseOrder "project/group2/features/orders/presentation/response"
	_helper "project/group2/helper"
	"strconv"

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

func (h *OrderHandler) Confirmed(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	idFromToken, _ := _middleware.ExtractToken(c)
	row, errCon := h.orderBusiness.ConfirmStatus(idOrder, idFromToken)
	if errCon != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *OrderHandler) Cancelled(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	idFromToken, _ := _middleware.ExtractToken(c)
	row, errCon := h.orderBusiness.CancelStatus(idOrder, idFromToken)
	if errCon != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *OrderHandler) History(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken, _ := _middleware.ExtractToken(c)
	result, err := h.orderBusiness.HistoryAll(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseOrder.FromCoreList(result)))
}

func (h *OrderHandler) OrderDetail(c echo.Context) error {
	id := c.Param("id")
	idOrder, _ := strconv.Atoi(id)
	result, err := h.orderBusiness.OrderDetails(idOrder)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseOrder.OdFromCoreList(result)))
}

func (h *OrderHandler) GetMyOrder(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken, _ := _middleware.ExtractToken(c)
	result, err := h.orderBusiness.GetMyDataOrder(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseOrder.FromCoreList(result)))
}
