package app

import (
	"context"
	"fmt"
	"net/http"

	"acort.link/acort.link/config"
	"acort.link/acort.link/core/adapters/postgres/repo"
	"acort.link/acort.link/core/domain/models"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ShortenApp interface {
	CreateShortenURL(ctx context.Context, url models.URL) (string, error)
	SearchUrl(ctx context.Context, path models.Path) (models.URL, error)
}

type shortenApp struct {
	config config.Config
	repo   repo.ShortenRepo
}

func NewShortenApp(config config.Config, repo repo.ShortenRepo) ShortenApp {
	return &shortenApp{config, repo}
}

func (s *shortenApp) CreateShortenURL(ctx context.Context, url models.URL) (string, error) {

	if url.Path == "" {

		url.Path = uuid.New().String()[:6]

	}

	urlBD, err := s.repo.SearchUrl(ctx, url.Path)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	fmt.Println(urlBD)

	if urlBD.URL != "" {
		return "", echo.NewHTTPError(http.StatusConflict, "path already exists")
	}

	if err := s.repo.CreateShorten(ctx, url); err != nil {
		println(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	return s.config.Domain.Link + url.Path, nil
}

func (s *shortenApp) SearchUrl(ctx context.Context, path models.Path) (models.URL, error) {

	url, err := s.repo.SearchUrl(ctx, path.Path)
	if err != nil {
		return models.URL{}, echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if url.URL == "" {
		return models.URL{}, echo.NewHTTPError(http.StatusNotFound, "url not found")
	}

	return url, nil
}
