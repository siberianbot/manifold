package config

import "manifold/internal/validation"

type Validatable interface {
	Validate(ctx validation.Context) error
}
