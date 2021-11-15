package traversing

import "manifold/internal/validation"

type Context interface {
	validation.Context

	CurrentFile() string
}
