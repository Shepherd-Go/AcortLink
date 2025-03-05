package ports

import (
	"acortlink/core/domain/models"
	"context"
)

type ShortenApp interface {
	CreateShortURL(ctx context.Context, url models.URL) (string, error)
	SearchUrl(ctx context.Context, path string) (string, error)
}

type ShortenRepo interface {
	CreateShorten(ctx context.Context, url models.URL) error
	SearchUrl(ctx context.Context, path string) (models.URL, error)
}
