package handler

import (
	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShortenerRequest interface {
	CreateShortURL(c echo.Context) error
	SearchOriginalUrl(c echo.Context) error
}

type shortenerRequest struct {
	shorten ports.ShortenApp
}

func NewShortener(shorten ports.ShortenApp) ShortenerRequest {
	return &shortenerRequest{shorten}
}

func (r *shortenerRequest) CreateShortURL(c echo.Context) error {

	ctx := c.Request().Context()

	url := models.URL{}

	if err := c.Bind(&url); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := url.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	link, err := r.shorten.CreateShortURL(ctx, url)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, link)
}

func (r *shortenerRequest) SearchOriginalUrl(c echo.Context) error {

	ctx := c.Request().Context()
	path := models.Path{}

	if err := c.Bind(&path); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := path.Validate(); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error())
	}

	url, err := r.shorten.SearchUrl(ctx, path.Path)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, url)
}
