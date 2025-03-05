package app

import (
	"context"
	"fmt"
	"net/http"

	"acortlink/config"
	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

type shortenApp struct {
	config config.Config
	repo   ports.ShortenRepo
	redis  *redis.Client
}

func NewShortenApp(config config.Config, repo ports.ShortenRepo, redis *redis.Client) ports.ShortenApp {
	return &shortenApp{config, repo, redis}
}

func (s *shortenApp) CreateShortURL(ctx context.Context, url models.URL) (string, error) {

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
