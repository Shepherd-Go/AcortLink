package groups

import (
	"acortlink/cmd/api/handler"

	"github.com/labstack/echo/v4"
)

type ShortenerRequest interface {
	Resource(g *echo.Group)
}

type shortenerRequest struct {
	handShortener handler.ShortenerRequest
}

func NewGroupShortener(handShortener handler.ShortenerRequest) ShortenerRequest {
	return &shortenerRequest{handShortener}
}

func (r *shortenerRequest) Resource(g *echo.Group) {
	g.POST("/create", r.handShortener.CreateShortURL)
	g.GET("/search", r.handShortener.RedirectURL)
}
