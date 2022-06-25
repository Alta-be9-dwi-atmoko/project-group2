package presentation

import (
	"net/http"
	"project/group2/features/users"
	"strconv"

	_responseUser "project/group2/features/users/presentation/response"
	_helper "project/group2/helper"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userBusiness users.Business
}

func NewUserHandler(userBusiness users.Business) *UserHandler {
	return &UserHandler{
		userBusiness: userBusiness,
	}
}

func (h *UserHandler) GetAll(c echo.Context) error {
	limit := c.QueryParam("limit")
	offset := c.QueryParam("offset")
	limitint, _ := strconv.Atoi(limit)
	offsetint, _ := strconv.Atoi(offset)
	result, err := h.userBusiness.GetAllData(limitint, offsetint)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data users"))
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("success", _responseUser.FromCoreList(result)))
}
