package repo

import (
	"context"

	"acortlink/core/domain/models"
	"acortlink/core/domain/ports"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type shortenRepo struct {
	db *gorm.DB
}

func NewShortenRepo(db *gorm.DB) ports.ShortenRepoPostgres {
	return &shortenRepo{db}
}

func (r *shortenRepo) Save(ctx context.Context, url models.URLCreate) error {

	if err := r.db.WithContext(ctx).
		Table("urls.urls").
		Create(&url).
		Error; err != nil {
		return err
	}

	return nil
}

func (r *shortenRepo) SearchUrl(ctx context.Context, path string) (models.URLResponse, error) {

	var url models.URLResponse

	if err := r.db.WithContext(ctx).
		Table("urls.urls").
		Where("path = ?", path).
		Find(&url).
		Error; err != nil {
		return models.URLResponse{}, err
	}

	return url, nil
}

func (r *shortenRepo) AddContToQuerysUrl(ctx context.Context, id uuid.UUID) error {

	var url models.URLResponse

	err := r.db.Transaction(func(tx *gorm.DB) error {

		if err := tx.WithContext(ctx).
			Table("urls.urls").
			Where("id = ?", id).
			Find(&url).
			Error; err != nil {
			return err
		}

		if err := tx.WithContext(ctx).
			Table("urls.urls").
			Where("id = ?", id).
			Update("number_of_queries", url.Number_Of_Queries+1).
			Error; err != nil {
			return err
		}

		return nil
	})

	return err

	// if err := r.db.WithContext(ctx).Table("urls.urls").
	// 	Where("id = ?", id).Update("number_of_queries", +2+3).Error; err != nil {
	// 	return err
	// }
	// return nil
}
