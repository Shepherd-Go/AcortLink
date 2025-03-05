package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type healthResponse struct {
	Code    int    `json:"status"`
	Message string `json:"message"`
}

func HealthCheck(context echo.Context) error {
	response := &healthResponse{
		Code:    http.StatusOK,
		Message: "Active!",
	}

	return context.JSON(http.StatusOK, response)
}
