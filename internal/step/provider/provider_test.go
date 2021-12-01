package provider

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/step"
	"testing"
)

func TestEmptyProvider(t *testing.T) {
	options := Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)
}

func TestFactoryNotExists(t *testing.T) {
	factoryName := "foo"
	options := Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)
	assert.Nil(t, provider.FactoryFor(factoryName))
}

func TestExecutorNotExists(t *testing.T) {
	executorName := "foo"
	options := Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)
	assert.Nil(t, provider.ExecutorFor(executorName))
}

func TestFactoryExists(t *testing.T) {
	factoryFn := func(_ interface{}) (step.Step, error) {
		return nil, nil
	}
	factoryName := "foo"

	options := Options{
		Factories: map[string]step.Factory{
			factoryName: factoryFn,
		},
		Executors: map[string]step.Executor{},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)

	factory := provider.FactoryFor(factoryName)
	assert.NotNil(t, factory)
}

func TestExecutorExists(t *testing.T) {
	executorFn := func(_ step.Step, _ *step.ExecutorContext) error {
		return nil
	}
	executorName := "foo"

	options := Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{
			executorName: executorFn,
		},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)

	executor := provider.ExecutorFor(executorName)
	assert.NotNil(t, executor)
}
