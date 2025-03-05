package provider

import (
	"acortlink/cmd/api/handler"
	"acortlink/cmd/api/router"
	"acortlink/cmd/api/router/groups"
	"acortlink/config"
	"acortlink/core/adapters/postgres"
	"acortlink/core/adapters/postgres/repo"
	"acortlink/core/adapters/redis"
	"acortlink/core/app"

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
