package http

import (
	"github.com/TeplyyMaksim/bookstore_oauth-api/domain/access_token"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
	"github.com/labstack/echo"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(echo.Context) error
	Create(ctx echo.Context) error
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c echo.Context) error {
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c echo.Context) error {
	var request access_token.AccessTokenRequest

	if err := c.Bind(&request); err != nil {
		restError := errors_utils.NewBadRequestError("invalid json body")
		return c.JSON(restError.Status, restError)
	}

	if err := handler.service.Create(&request); err != nil {
		return c.JSON(err.Status, err)
	}

	return c.JSON(http.StatusCreated, request)
}
