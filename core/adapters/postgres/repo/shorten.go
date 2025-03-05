package repo

import (
	"context"
	"time"

	"acortlink/core/domain/models"

	"gorm.io/gorm"
)

type ShortenRepo interface {
	CreateShorten(ctx context.Context, url models.URL) error
	SearchUrl(ctx context.Context, path string) (models.URL, error)
}

type shortenRepo struct {
	db *gorm.DB
}

func NewShortenRepo(db *gorm.DB) ShortenRepo {
	return &shortenRepo{db}
}

func (r *shortenRepo) CreateShorten(ctx context.Context, url models.URL) error {

	if err := r.db.WithContext(ctx).
		Table("urls.urls").
		Create(&url).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *shortenRepo) SearchUrl(ctx context.Context, path string) (models.URL, error) {

	var url models.URL

	if err := r.db.WithContext(ctx).
		Table("urls.urls").
		Where("path = ? and (expires_at > ? or expires_at = ?)", path, time.Now().UTC(), "0001-01-01 00:00:00.000").
		Find(&url).
		Error; err != nil {
		return models.URL{}, err
	}

	return url, nil
}
