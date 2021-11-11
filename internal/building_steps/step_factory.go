package building_steps

import (
	"errors"
	dd "manifold/internal/document_definition"
)

func FromStepDefinition(stepDefinition *dd.StepDefinition) (Step, error) {
	// TODO: handle different definitions here
	return fromCmdStepDefinition(stepDefinition)
}

func fromCmdStepDefinition(stepDefinition *dd.StepDefinition) (Step, error) {
	if stepDefinition.Command == "" {
		return nil, errors.New("step cmd doesn't have any command to execute")
	}

	step := CommandStep{
		Command: stepDefinition.Command,
	}

	return step, nil
}
