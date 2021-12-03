package step

import (
	"github.com/stretchr/testify/mock"
	"manifold/internal/config"
	"manifold/internal/step"
)

type Step struct {
	mock.Mock
}

func (m *Step) Name() string {
	return m.Called().String(0)
}

type StepProvider struct {
	mock.Mock
}

func (m *StepProvider) FactoryFor(name string) step.Factory {
	factory := m.Called(name).Get(0)

	if factory == nil {
		return nil
	}

	return factory.(step.Factory)
}

func (m *StepProvider) ExecutorFor(name string) step.Executor {
	executor := m.Called(name).Get(0)

	if executor == nil {
		return nil
	}

	return executor.(step.Executor)
}

type StepFactory struct {
	mock.Mock
}

func (m *StepFactory) CreateFrom(value interface{}) (step.Step, error) {
	args := m.Called(value)

	var result step.Step

	if args.Get(0) == nil {
		result = nil
	} else {
		result = args.Get(0).(step.Step)
	}

	var err error

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Error(1)
	}

	return result, err
}

type StepExecutor struct {
	mock.Mock
}

func (m *StepExecutor) Execute(context step.ExecutorContext) error {
	return m.Called(context).Error(0)
}

type StepBuilder struct {
	mock.Mock
}

func (m *StepBuilder) FromConfig(definition config.ProjectStep) (step.Step, error) {
	args := m.Called(definition)

	var result step.Step

	if args.Get(0) == nil {
		result = nil
	} else {
		result = args.Get(0).(step.Step)
	}

	var err error

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Error(1)
	}

	return result, err
}
