package ports

import (
	"acortlink/core/domain/models"
	"context"
	"time"

	"github.com/google/uuid"
)

type ShortenApp interface {
	CreateShortURL(ctx context.Context, url models.URLCreate) (string, error)
	SearchUrl(ctx context.Context, path string) (string, error)
}

type ShortenRepoPostgres interface {
	Save(ctx context.Context, url models.URLCreate) error
	SearchUrl(ctx context.Context, path string) (models.URLResponse, error)
	AddContToQuerysUrl(ctx context.Context, id uuid.UUID) error
}

type ShortenRepoRedis interface {
	Save(ctx context.Context, key string, url models.URLResponse, time time.Duration) error
	SearchUrl(ctx context.Context, path string) (models.URLResponse, error)
}
