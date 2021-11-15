package steps

import (
	dd "manifold/internal/document_definition"
	"manifold/internal/validation"
)

func FromStepDefinition(stepDefinition dd.StepDefinition, ctx validation.Context) Step {
	// TODO: change a way of definition of step
	switch {
	case stepDefinition.Command != "":
		return fromCmdStepDefinition(stepDefinition, ctx)

	default:
		ctx.AddError(validation.StepNotMatch)
		return nil
	}
}

func fromCmdStepDefinition(stepDefinition dd.StepDefinition, _ validation.Context) Step {
	return CommandStep{
		Command: stepDefinition.Command,
	}
}
