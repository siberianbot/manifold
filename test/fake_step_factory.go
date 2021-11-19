package test

import (
	"errors"
	"manifold/internal/steps"
)

type FakeStepFactory struct {
	FactoryName    string
	ConstructFails bool
	StepFails      bool
}

func (factory FakeStepFactory) Name() string {
	return factory.FactoryName
}

func (factory FakeStepFactory) Construct(_ interface{}) (steps.Step, error) {
	if factory.ConstructFails {
		return nil, errors.New(Error)
	} else {
		return NewFakeStep(factory.FactoryName, factory.StepFails), nil
	}
}

func NewFakeStepFactory(name string, stepFails bool, constructFails bool) steps.StepFactory {
	factory := new(FakeStepFactory)
	factory.FactoryName = name
	factory.StepFails = stepFails
	factory.ConstructFails = constructFails

	return factory
}
