package handler

import (
	"acortlink/core/app"
	"acortlink/core/domain/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ShortenerRequest interface {
	CreateShortURL(c echo.Context) error
	RedirectURL(c echo.Context) error
}

type shortenerRequest struct {
	shorten app.ShortenApp
}

func NewShortener(shorten app.ShortenApp) ShortenerRequest {
	return &shortenerRequest{shorten}
}

func (r *shortenerRequest) CreateShortURL(c echo.Context) error {

	ctx := c.Request().Context()

	url := models.URL{}

	if err := c.Bind(&url); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	if err := url.Validate(); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	link, err := r.shorten.CreateShortenURL(ctx, url)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, link)
}

func (r *shortenerRequest) RedirectURL(c echo.Context) error {

	ctx := c.Request().Context()
	path := models.Path{}

	if err := c.Bind(&path); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	if err := path.Validate(); err != nil {
		return echo.NewHTTPError(400, err.Error())
	}

	url, err := r.shorten.SearchUrl(ctx, path.Path)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, url)
}
