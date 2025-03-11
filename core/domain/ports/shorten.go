package ports

import (
	"acortlink/core/domain/models"
	"context"
	"time"
)

type ShortenApp interface {
	CreateShortURL(ctx context.Context, url models.URL) (string, error)
	SearchUrl(ctx context.Context, path string) (string, error)
}

type ShortenRepoPostgres interface {
	Save(ctx context.Context, url models.URL) error
	SearchUrl(ctx context.Context, path string) (models.URL, error)
}

type ShortenRepoRedis interface {
	Save(ctx context.Context, key string, value interface{}, time time.Duration) error
	SearchUrl(ctx context.Context, path string) (string, error)
}
