package handler_test

import (
	"acortlink/cmd/api/handler"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var healthJson = `{"status":200,"message":"Active!"}`

type ControllerCase struct {
	Req *http.Request
	Res *httptest.ResponseRecorder
	Ctx echo.Context
}

func SetupControllerCase(url, method string, body io.Reader) ControllerCase {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)

	return ControllerCase{req, res, ctx}
}

func TestHealthCheck(t *testing.T) {
	controller := SetupControllerCase("/api/health", http.MethodGet, nil)

	if assert.NoError(t, handler.HealthCheck(controller.Ctx)) {
		assert.Equal(t, http.StatusOK, controller.Res.Code)
		assert.Equal(t, healthJson, strings.TrimSpace(controller.Res.Body.String()))
	}
}
