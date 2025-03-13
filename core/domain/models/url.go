package models

import (
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type URLCreate struct {
	URL    string `json:"url" gorm:"column:original_url" validate:"required,url" mold:"trim"`
	Domain string `json:"domain" gorm:"column:domain" validate:"required,url" mold:"trim"`
	Path   string `json:"path" gorm:"column:path" validate:"max=6" mold:"trim"`
}

type URLResponse struct {
	ID                uuid.UUID `json:"id"`
	URL               string    `json:"url" gorm:"column:original_url"`
	Domain            string    `json:"domain"`
	Path              string    `json:"path"`
	Number_Of_Queries int       `json:"number_of_queries"`
}

func (u *URLCreate) Validate() error {
	_ = conform.Struct(context.Background(), u)
	return validate.Struct(u)
}
