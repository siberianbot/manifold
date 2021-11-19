package steps

import (
	dd "manifold/internal/document_definition"
	"manifold/internal/validation"
)

func FromStepDefinition(stepDefinition dd.StepDefinition, ctx validation.Context) Step {
	if len(stepDefinition) > 1 {
		ctx.AddError(validation.StepWithManyToolchains)
		return nil
	}

	// TODO: change a way of definition of step
	switch {
	case stepDefinition["cmd"] != nil:
		return fromCmdStepDefinition(stepDefinition["cmd"], ctx)

	default:
		ctx.AddError(validation.StepNotMatch)
		return nil
	}
}

func fromCmdStepDefinition(cmdDefinition interface{}, ctx validation.Context) Step {
	cmd, ok := cmdDefinition.(string)

	if !ok {
		ctx.AddError(validation.CmdStepIsInvalid)
		return nil
	}

	return CommandStep{
		Command: cmd,
	}
}
