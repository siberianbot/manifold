package command_step

import (
	"errors"
	"manifold/internal/steps"
	"manifold/internal/validation"
)

type commandStepFactory struct {
	executor *commandStepExecutor
}

func (commandStepFactory) Name() string {
	return "cmd"
}

func (factory commandStepFactory) Construct(definition interface{}) (steps.Step, error) {
	cmd, ok := definition.(string)

	if !ok || cmd == "" {
		return nil, errors.New(validation.CmdStepIsInvalid)
	}

	return newCommandStep(cmd, factory.executor), nil
}

func NewStepFactory() steps.StepFactory {
	factory := new(commandStepFactory)
	factory.executor = newCommandStepExecutor()

	return factory
}
