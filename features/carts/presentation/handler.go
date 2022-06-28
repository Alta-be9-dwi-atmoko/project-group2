package presentation

import (
	"net/http"
	_cart "project/group2/features/carts"
	_requestCart "project/group2/features/carts/presentation/request"
	_responseCart "project/group2/features/carts/presentation/response"
	_middleware "project/group2/features/middlewares"
	_helper "project/group2/helper"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartBusiness _cart.Business
}

func NewCartHandler(cartBusiness _cart.Business) *CartHandler {
	return &CartHandler{
		cartBusiness: cartBusiness,
	}
}

func (h *CartHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	idFromToken, _ := _middleware.ExtractToken(c)
	result, err := h.cartBusiness.GetAllData(limitint, offsetint, idFromToken)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseCart.FromCoreList(result)))
}

func (h *CartHandler) PostCart(c echo.Context) error {
	idFromToken, _ := _middleware.ExtractToken(c)
	cartReq := _requestCart.Cart{}
	err := c.Bind(&cartReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	dataCart := _cart.Core{}
	dataCart.Product.ID = cartReq.IdProduct
	dataCart.Qty = cartReq.Qty
	dataCart.Status = "Pending"
	dataCart.UserID = idFromToken
	row, errCreate := h.cartBusiness.CreateData(dataCart)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to add to cart"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *CartHandler) UpdateCart(c echo.Context) error {
	id := c.Param("id")
	idCart, _ := strconv.Atoi(id)
	idFromToken, _ := _middleware.ExtractToken(c)
	cartReq := _requestCart.Cart{}
	err := c.Bind(&cartReq)

	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	qty := cartReq.Qty
	row, errUpd := h.cartBusiness.UpdateData(qty, idCart, idFromToken)
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("you dont have access"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
