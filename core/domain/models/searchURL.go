package models

import "context"

type Path struct {
	Path string `json:"path" query:"path" validate:"required"`
}

func (p *Path) Validate() error {
	_ = conform.Struct(context.Background(), p)
	return validate.Struct(p)
}
