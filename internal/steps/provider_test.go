package steps

import (
	"errors"
	"fmt"
	"manifold/internal/config"
	"manifold/internal/validation"
	"testing"
)

func TestNewProvider(t *testing.T) {
	options := NewProviderOptions()
	provider := NewProvider(options)

	if provider == nil {
		t.Error("provider is nil")
	} else if provider.options == nil {
		t.Error("provider.options is nil")
	}
}

func TestProvider(t *testing.T) {
	t.Run("CreateFrom", func(t *testing.T) {
		t.Run("EmptyStep", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			configStep := make(config.Step)

			step, err := provider.CreateFrom(configStep)

			if step != nil {
				t.Error("step is not nil")
			}

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != validation.EmptyStep {
				t.Errorf("error is %s, not %s", err.Error(), validation.EmptyStep)
			}
		})

		t.Run("NilStep", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			step, err := provider.CreateFrom(nil)

			if step != nil {
				t.Error("step is not nil")
			}

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != validation.EmptyStep {
				t.Errorf("error is %s, not %s", err.Error(), validation.EmptyStep)
			}
		})

		t.Run("NoFactories", func(t *testing.T) {
			options := NewProviderOptions()
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			if step != nil {
				t.Error("step is not nil")
			}

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != validation.StepNotMatchedAnyToolchain {
				t.Errorf("error is %s, not %s", err.Error(), validation.StepNotMatchedAnyToolchain)
			}
		})

		t.Run("WithNotMatchingFactory", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["bar"] = func(_ interface{}) (Step, error) { return nil, nil }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			if step != nil {
				t.Error("step is not nil")
			}

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != validation.StepNotMatchedAnyToolchain {
				t.Errorf("error is %s, not %s", err.Error(), validation.StepNotMatchedAnyToolchain)
			}
		})

		t.Run("WithMatchingFactoryButFails", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["foo"] = func(_ interface{}) (Step, error) { return nil, errors.New("error") }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			if step != nil {
				t.Error("step is not nil")
			}

			expected := fmt.Sprintf(validation.StepFailed, "foo", "error")

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != expected {
				t.Errorf("error is %s, not %s", err.Error(), expected)
			}
		})

		t.Run("WithMatchingFactoryAndNotFails", func(t *testing.T) {
			options := NewProviderOptions()
			options.Factories["foo"] = func(_ interface{}) (Step, error) { return newTestStep("foo"), nil }
			provider := NewProvider(options)

			configStep := make(config.Step)
			configStep["foo"] = "foo"

			step, err := provider.CreateFrom(configStep)

			if step == nil {
				t.Error("step is nil")
			} else {
				_, ok := step.(*testStep)

				if !ok {
					t.Error("step.name is not testStep")
				} else if step.Name() != "foo" {
					t.Errorf("step.name is %s, not %s", step.Name(), "foo")
				}
			}

			if err != nil {
				t.Errorf("error is %s, not nil", err.Error())
			}
		})
	})
}
