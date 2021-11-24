package steps

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"testing"
)

func TestNewProvider(t *testing.T) {
	options := NewProviderOptions()
	provider := NewProvider(options)

	assert.NotEmpty(t, provider)
	assert.NotEmpty(t, provider.options)
}

func TestProvider(t *testing.T) {
	t.Run("CreateFrom", func(t *testing.T) {
		t.Run("EmptyStep", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			configStep := make(config.Step)

			step, err := provider.CreateFrom(configStep)

			assert.Empty(t, step)
			assert.EqualError(t, err, EmptyStep)
		})

		t.Run("NilStep", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			step, err := provider.CreateFrom(nil)

			assert.Empty(t, step)
			assert.EqualError(t, err, EmptyStep)
		})

		t.Run("NoFactories", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			assert.Empty(t, step)
			assert.EqualError(t, err, StepNotMatched)
		})

		t.Run("WithNotMatchingFactory", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["bar"] = func(_ interface{}) (Step, error) { return nil, nil }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			assert.Empty(t, step)
			assert.EqualError(t, err, StepNotMatched)
		})

		t.Run("WithMatchingFactoryButFails", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["foo"] = func(_ interface{}) (Step, error) { return nil, errors.New("error") }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			assert.Empty(t, step)
			assert.EqualError(t, err, fmt.Sprintf(StepFailed, "foo", "error"))
		})

		t.Run("WithMatchingFactoryAndNotFails", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["foo"] = func(_ interface{}) (Step, error) { return newTestStep("foo"), nil }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			assert.NotEmpty(t, step)
			assert.NoError(t, err)
			assert.IsType(t, newTestStep("foo"), step)
			assert.Equal(t, "foo", step.Name())
		})
	})
}
