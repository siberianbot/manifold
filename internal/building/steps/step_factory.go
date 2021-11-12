package steps

import (
	"manifold/internal/building"
	dd "manifold/internal/document_definition"
)

func FromStepDefinition(stepDefinition *dd.StepDefinition, ctx building.TraverseContext) Step {
	// TODO: handle different definitions here
	return fromCmdStepDefinition(stepDefinition, ctx)
}

func fromCmdStepDefinition(stepDefinition *dd.StepDefinition, ctx building.TraverseContext) Step {
	if stepDefinition.Command == "" {
		ctx.AddError("step cmd does not have any command to execute")
		return nil
	}

	step := CommandStep{
		Command: stepDefinition.Command,
	}

	return step
}
