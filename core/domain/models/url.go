package models

import (
	"context"

	"github.com/go-playground/mold/modifiers"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	conform  = modifiers.New()
)

type URL struct {
	URL  string `json:"url" validate:"required,url" mold:"trim"`
	Path string `json:"path" mold:"trim"`
}

func (u *URL) Validate() error {
	_ = conform.Struct(context.Background(), u)
	return validate.Struct(u)
}
