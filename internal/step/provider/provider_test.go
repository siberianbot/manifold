package provider

import (
	"github.com/stretchr/testify/assert"
	mockStep "manifold/internal/mock/step"
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
	factory := new(mockStep.StepFactory)
	factoryName := "foo"

	options := Options{
		Factories: map[string]step.Factory{
			factoryName: factory,
		},
		Executors: map[string]step.Executor{},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)

	f := provider.FactoryFor(factoryName)
	assert.NotNil(t, f)
	assert.Equal(t, factory, f)
}

func TestExecutorExists(t *testing.T) {
	executor := new(mockStep.StepExecutor)
	executorName := "foo"

	options := Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{
			executorName: executor,
		},
	}

	provider := NewProvider(options)

	assert.NotEmpty(t, provider)

	e := provider.ExecutorFor(executorName)
	assert.NotNil(t, e)
}
