package steps

import (
	dd "manifold/internal/document_definition"
	"manifold/internal/traversing"
)

func FromStepDefinition(stepDefinition *dd.StepDefinition, ctx traversing.TraverseContext) Step {
	// TODO: handle different definitions here
	return fromCmdStepDefinition(stepDefinition, ctx)
}

func fromCmdStepDefinition(stepDefinition *dd.StepDefinition, ctx traversing.TraverseContext) Step {
	if stepDefinition.Command == "" {
		ctx.AddError("step cmd does not have any command to execute")
		return nil
	}

	step := CommandStep{
		Command: stepDefinition.Command,
	}

	return step
}
