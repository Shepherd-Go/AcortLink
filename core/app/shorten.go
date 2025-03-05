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
	"github.com/redis/go-redis/v9"
)

type ShortenApp interface {
	CreateShortenURL(ctx context.Context, url models.URL) (string, error)
	SearchUrl(ctx context.Context, path string) (string, error)
}

type shortenApp struct {
	config config.Config
	repo   repo.ShortenRepo
	redis  *redis.Client
}

func NewShortenApp(config config.Config, repo repo.ShortenRepo, redis *redis.Client) ShortenApp {
	return &shortenApp{config, repo, redis}
}

func (s *shortenApp) CreateShortenURL(ctx context.Context, url models.URL) (string, error) {

	if url.Path == "" {

		url.Path = uuid.New().String()[:6]

	}

	urlBD, err := s.repo.SearchUrl(ctx, url.Path)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if urlBD.URL != "" {
		return "", echo.NewHTTPError(http.StatusConflict, "path already exists")
	}

	if err := s.repo.CreateShorten(ctx, url); err != nil {
		println(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	err = s.redis.Set(ctx, url.Path, url.URL, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}

	return s.config.Domain.Link + url.Path, nil
}

func (s *shortenApp) SearchUrl(ctx context.Context, path string) (string, error) {

	val, err := s.redis.Get(context.Background(), path).Result()
	if err != nil {
		fmt.Println("Error al obtener valor de Redis:", err)
	}

	if val != "" {
		return val, nil
	}

	url, err := s.repo.SearchUrl(ctx, path)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if url.URL == "" {
		return "", echo.NewHTTPError(http.StatusNotFound, "url not found")
	}

	err = s.redis.Set(ctx, path, url.URL, 0).Err()
	if err != nil {
		fmt.Println(err.Error())
	}

	return url.URL, nil
}
