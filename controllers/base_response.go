package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = data

	return c.JSON(http.StatusOK, response)
}