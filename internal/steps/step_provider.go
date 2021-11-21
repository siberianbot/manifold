package steps

import (
	"manifold/internal/config"
	"manifold/internal/validation"
)

type StepProvider interface {
	CreateFor(definition config.Step, ctx validation.Context) Step
}
