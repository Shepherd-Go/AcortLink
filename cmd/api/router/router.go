package router

import (
	"net/http"

	"acortlink/cmd/api/handler"
	"acortlink/cmd/api/router/groups"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Router struct {
	server    *echo.Echo
	shortener groups.ShortenerRequest
}

func New(server *echo.Echo, shortener groups.ShortenerRequest) *Router {
	return &Router{
		server,
		shortener,
	}
}

func (r *Router) Init() {

	r.server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency=${latency_human}\n\n",
	}))

	r.server.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:                             []string{"*"},
		AllowMethods:                             []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:                             []string{echo.HeaderContentType},
		AllowCredentials:                         true,
		UnsafeWildcardOriginWithAllowCredentials: true,
	}))

	r.server.Use(middleware.Recover())

	basePath := r.server.Group("/api")
	basePath.GET("/health", handler.HealthCheck)

	r.shortener.Resource(basePath)

}
