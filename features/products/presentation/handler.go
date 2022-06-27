package presentation

import (
	"net/http"
	_middleware "project/group2/features/middlewares"
	"project/group2/features/products"
	_requestProduct "project/group2/features/products/presentation/request"
	_responseProduct "project/group2/features/products/presentation/response"
	"strconv"

	_helper "project/group2/helper"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productBusiness products.Business
}

func NewProductHandler(productBusiness products.Business) *ProductHandler {
	return &ProductHandler{
		productBusiness: productBusiness,
	}
}

func (h *ProductHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.productBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all products"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseProduct.FromCoreList(result)))
}

func (h *ProductHandler) PostProduct(c echo.Context) error {
	prodReq := _requestProduct.Product{}
	id, _ := _middleware.ExtractToken(c)
	err := c.Bind(&prodReq)
	prodReq.UserID = id
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataProduct := _requestProduct.ToCore(prodReq)
	row, errCreate := h.productBusiness.CreateData(dataProduct)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}

	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to create product, check your input"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *ProductHandler) GetById(c echo.Context) error {
	id := c.Param("id")
	idUser, _ := strconv.Atoi(id)
	result, errGet := h.productBusiness.GetProductById(idUser)
	if errGet != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get data product"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseProduct.FromCore(result)))
}

func (h *ProductHandler) PutProduct(c echo.Context) error {
	id := c.Param("id")
	idProd, _ := strconv.Atoi(id)
	idFromToken, _ := _middleware.ExtractToken(c)
	prodReq := _requestProduct.Product{}
	err := c.Bind(&prodReq)
	if prodReq.UserID != idFromToken {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("you dont have access"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to bind data, check your input"))
	}
	dataProduct := _requestProduct.ToCore(prodReq)
	row, errUpd := h.productBusiness.UpdateData(dataProduct, idProd)
	if errUpd != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to update data product"))
	}
	if row == 0 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to update data product"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}
