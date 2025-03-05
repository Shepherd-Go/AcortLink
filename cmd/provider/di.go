package provider

import (
	"acort.link/acort.link/cmd/api/handler"
	"acort.link/acort.link/cmd/api/router"
	"acort.link/acort.link/cmd/api/router/groups"
	"acort.link/acort.link/config"
	"acort.link/acort.link/core/adapters/postgres"
	"acort.link/acort.link/core/adapters/postgres/repo"
	"acort.link/acort.link/core/adapters/redis"
	"acort.link/acort.link/core/app"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

var Container *dig.Container

func BuildContainer() *dig.Container {
	Container = dig.New()

	_ = Container.Provide(func() config.Config {
		config.Environments()
		return *config.Get()
	})

	_ = Container.Provide(func() *echo.Echo {
		return echo.New()
	})

	_ = Container.Provide(postgres.NewPostgresConnection)
	_ = Container.Provide(redis.NewRedisConnection)

	_ = Container.Provide(router.New)

	_ = Container.Provide(groups.NewGroupShortener)

	_ = Container.Provide(handler.NewShortener)

	_ = Container.Provide(app.NewShortenApp)

	_ = Container.Provide(repo.NewShortenRepo)

	return Container

}
