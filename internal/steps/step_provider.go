package steps

import (
	"manifold/internal/document_definition"
	"manifold/internal/validation"
)

type StepProvider interface {
	CreateFor(definition document_definition.StepDefinition, ctx validation.Context) Step
}
