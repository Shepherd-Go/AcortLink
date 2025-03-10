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
)

type shortenApp struct {
	config config.Config
	postgr ports.ShortenRepoPostgres
	redis  ports.ShortenRepoRedis
}

func NewShortenApp(config config.Config, repo ports.ShortenRepoPostgres, redis ports.ShortenRepoRedis) ports.ShortenApp {
	return &shortenApp{config, repo, redis}
}

func (s *shortenApp) CreateShortURL(ctx context.Context, url models.URL) (string, error) {

	if url.Path == "" {
		url.Path = uuid.New().String()[:6]
	}

	urlBD, err := s.postgr.SearchUrl(ctx, url.Path)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if urlBD.URL != "" {
		return "", echo.NewHTTPError(http.StatusConflict, "path already exists")
	}

	if err := s.postgr.CreateShorten(ctx, url); err != nil {
		fmt.Println(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")

	}

	return s.config.Domain.Link + url.Path, nil
}

func (s *shortenApp) SearchUrl(ctx context.Context, path string) (string, error) {

	val, err := s.redis.SearchUrl(ctx, path)
	if err != nil {
		fmt.Println("Error al obtener valor de Redis:", err)
	}

	if val != "" {
		return val, nil
	}

	url, err := s.postgr.SearchUrl(ctx, path)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if url.URL == "" {
		return "", echo.NewHTTPError(http.StatusNotFound, "url not found")
	}

	err = s.redis.CreateShorten(ctx, path, url.URL, 0)
	if err != nil {
		fmt.Println(err.Error())
	}

	return url.URL, nil
}
