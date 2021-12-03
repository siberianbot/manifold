package builder

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/mock/step"
	"testing"
)

func TestEmptyDefinition(t *testing.T) {
	provider := new(step.StepProvider)

	definition := config.ProjectStep{}

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	result, err := builder.FromConfig(definition)

	assert.Nil(t, result)
	assert.EqualError(t, err, emptyStep)
}

func TestUnknownDefinition(t *testing.T) {
	provider := new(step.StepProvider)
	provider.On("FactoryFor", "foo").Return(nil)

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	definition := config.ProjectStep{
		"foo": struct{}{},
	}

	result, err := builder.FromConfig(definition)

	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf(unknownStep, "foo"))

	provider.AssertExpectations(t)
}

func TestValidDefinition(t *testing.T) {
	stepValue := struct{}{}
	stepMock := new(step.Step)

	factory := new(step.StepFactory)
	factory.On("CreateFrom", stepValue).Return(stepMock, nil)

	provider := new(step.StepProvider)
	provider.On("FactoryFor", "foo").Return(factory)

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	definition := config.ProjectStep{
		"foo": stepValue,
	}

	result, err := builder.FromConfig(definition)

	assert.NotNil(t, result)
	assert.Equal(t, stepMock, result)
	assert.NoError(t, err)

	factory.AssertExpectations(t)
	provider.AssertExpectations(t)
}

func TestAmbiguousDefinition(t *testing.T) {
	stepValue := struct{}{}

	fooFactory := new(step.StepFactory)
	barFactory := new(step.StepFactory)

	provider := new(step.StepProvider)
	provider.On("FactoryFor", "foo").Return(fooFactory)
	provider.On("FactoryFor", "bar").Return(barFactory)

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	definition := config.ProjectStep{
		"foo": stepValue,
		"bar": stepValue,
	}

	result, err := builder.FromConfig(definition)

	assert.Nil(t, result)
	assert.EqualError(t, err, ambiguousStep)

	provider.AssertExpectations(t)
}

func TestFailedToCreate(t *testing.T) {
	stepValue := struct{}{}

	fooFactoryErr := errors.New("error")
	fooFactory := new(step.StepFactory)
	fooFactory.On("CreateFrom", stepValue).Return(nil, fooFactoryErr)

	provider := new(step.StepProvider)
	provider.On("FactoryFor", "foo").Return(fooFactory)

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	definition := config.ProjectStep{
		"foo": stepValue,
	}

	result, err := builder.FromConfig(definition)

	assert.Nil(t, result)
	assert.EqualError(t, err, fmt.Sprintf(stepFactoryFailed, "foo", fooFactoryErr))

	fooFactory.AssertExpectations(t)
	provider.AssertExpectations(t)
}
