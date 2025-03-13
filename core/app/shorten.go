package app

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type shortenApp struct {
	postgr ports.ShortenRepoPostgres
	redis  ports.ShortenRepoRedis
}

func NewShortenApp(repo ports.ShortenRepoPostgres, redis ports.ShortenRepoRedis) ports.ShortenApp {
	return &shortenApp{repo, redis}
}

func (s *shortenApp) CreateShortURL(ctx context.Context, url models.URLCreate) (string, error) {

	if url.Path == "" {
		url.Path = uuid.New().String()[:6]
	}

	urlBD, err := s.postgr.SearchUrl(ctx, url.Path)
	if err != nil {
		fmt.Println(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
	}

	if urlBD.ID != uuid.Nil {
		return "", echo.NewHTTPError(http.StatusConflict, "path already exists")
	}

	if err := s.postgr.Save(ctx, url); err != nil {
		fmt.Println(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")

	}

	return url.Domain + url.Path, nil
}

func (s *shortenApp) SearchUrl(ctx context.Context, path string) (string, error) {

	url, err := s.redis.SearchUrl(ctx, path)
	if err != nil {
		fmt.Println("Error al obtener valor de Redis:", err)
	}

	if url.ID == uuid.Nil {

		url, err = s.postgr.SearchUrl(ctx, path)
		if err != nil {
			fmt.Println(err.Error())
			return "", echo.NewHTTPError(http.StatusInternalServerError, "unexpected error")
		}

		if url.ID == uuid.Nil {
			return "", echo.NewHTTPError(http.StatusNotFound, "url not found")
		}

		err = s.redis.Save(ctx, path, url, 24*time.Hour)
		if err != nil {
			fmt.Println(err.Error())
		}

	}

	if err := s.postgr.AddContToQuerysUrl(ctx, url.ID); err != nil {
		fmt.Println(err.Error())
	}

	return url.URL, nil
}
