package presentation

import (
	"net/http"
	"project/group2/features/users"
	"strconv"

	_requestUser "project/group2/features/users/presentation/request"
	_responseUser "project/group2/features/users/presentation/response"
	_helper "project/group2/helper"

	"github.com/labstack/echo/v4"
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

func (h *UserHandler) PostUser(c echo.Context) error {
	userReq := _requestUser.User{}
	err := c.Bind(&userReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to bind data, check your input"))
	}

	dataUser := _requestUser.ToCore(userReq)
	row, errCreate := h.userBusiness.CreateData(dataUser)
	if row == -1 {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("please make sure all fields are filled in correctly"))
	}
	if errCreate != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to create data users"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseOkNoData("success"))
}

func (h *UserHandler) LoginAuth(c echo.Context) error {
	authData := users.AuthRequestData{}
	c.Bind(&authData)
	token, name, e := h.userBusiness.LoginUser(authData)
	if e != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("email or password incorrect"))
	}

	data := map[string]interface{}{
		"token": token,
		"name":  name,
	}
	return c.JSON(http.StatusOK, _helper.ResponseOkWithData("login success", data))
}
