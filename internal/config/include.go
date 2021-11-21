package config

import (
	"manifold/internal/utils"
	"manifold/internal/validation"
)

type Include interface {
	Validatable

	Path() string
}

type include struct {
	path string
}

func (i *include) Validate(ctx validation.Context) error {
	return validation.ValidatePath(utils.BuildPath(ctx.Dir(), i.path))
}

func (i *include) Path() string {
	return i.path
}

func newInclude(path string) Include {
	inc := new(include)
	inc.path = path

	return inc
}
