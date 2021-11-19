package steps

import (
	"manifold/internal/config"
	"manifold/internal/validation"
)

type StepProvider interface {
	CreateFor(definition config.StepDefinition, ctx validation.Context) Step
}
