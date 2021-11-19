package traversing

import (
	"manifold/internal/steps"
	"manifold/internal/validation"
)

type Context interface {
	validation.Context

	CurrentFile() string
	GetStepProvider() steps.StepProvider
}
