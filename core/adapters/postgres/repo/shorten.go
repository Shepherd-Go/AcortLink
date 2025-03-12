package repo

import (
	"context"

	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"

	"gorm.io/gorm"
)

type shortenRepo struct {
	db *gorm.DB
}

func NewShortenRepo(db *gorm.DB) ports.ShortenRepoPostgres {
	return &shortenRepo{db}
}

func (r *shortenRepo) Save(ctx context.Context, url models.URL) error {

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
		Where("path = ?", path).
		Find(&url).
		Error; err != nil {
		return models.URL{}, err
	}

	return url, nil
}
