package test

import (
	"errors"
	"manifold/internal/steps"
)

type FakeStep struct {
	Name  string
	Fails bool
}

func (step FakeStep) Execute() error {
	if step.Fails {
		return errors.New(Error)
	} else {
		return nil
	}
}

func NewFakeStep(name string, fails bool) steps.Step {
	step := new(FakeStep)
	step.Name = name
	step.Fails = fails

	return step
}
