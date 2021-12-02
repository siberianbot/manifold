package builder

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/mock"
	"testing"
)

func TestEmptyDefinition(t *testing.T) {
	provider := new(mock.StepProvider)

	definition := config.ProjectStep{}

	builder := NewBuilder(provider)

	assert.NotEmpty(t, builder)

	result, err := builder.FromConfig(definition)

	assert.Nil(t, result)
	assert.EqualError(t, err, emptyStep)
}

func TestUnknownDefinition(t *testing.T) {
	provider := new(mock.StepProvider)
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
	stepMock := new(mock.Step)

	factory := new(mock.StepFactory)
	factory.On("CreateFrom", stepValue).Return(stepMock, nil)

	provider := new(mock.StepProvider)
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

	fooFactory := new(mock.StepFactory)
	barFactory := new(mock.StepFactory)

	provider := new(mock.StepProvider)
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
	fooFactory := new(mock.StepFactory)
	fooFactory.On("CreateFrom", stepValue).Return(nil, fooFactoryErr)

	provider := new(mock.StepProvider)
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
