package steps_test

import (
	"fmt"
	"manifold/internal/config"
	"manifold/internal/steps"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestStepProvider(t *testing.T) {
	t.Run("InstantiationCheck", func(t *testing.T) {
		provider := steps.NewDefaultStepProvider()

		test.Assert(t, provider != nil)
	})

	t.Run("EmptyDefinition", func(t *testing.T) {
		expected := fmt.Sprintf(validation.StepResolveFailed, validation.EmptyStep)

		provider := steps.NewDefaultStepProvider()
		configStep := make(config.Step)

		step, err := provider.CreateFrom(configStep)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("NoFactories", func(t *testing.T) {
		name := "foo"
		expected := fmt.Sprintf(validation.StepResolveFailed, fmt.Sprintf(validation.StepNotMatchedAnyToolchain, name))

		provider := steps.NewDefaultStepProvider()
		configStep := make(config.Step)
		configStep[name] = name

		step, err := provider.CreateFrom(configStep)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("FactoryWithDifferentName", func(t *testing.T) {
		name := "foo"
		expected := fmt.Sprintf(validation.StepResolveFailed, fmt.Sprintf(validation.StepNotMatchedAnyToolchain, name))

		barFactory := test.NewFakeStepFactory("bar", false, false)
		provider := steps.NewDefaultStepProvider(barFactory)
		fooConfigStep := make(config.Step)
		fooConfigStep[name] = name

		step, err := provider.CreateFrom(fooConfigStep)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("FactoriesWithSameName", func(t *testing.T) {
		name := "foo"
		expected := fmt.Sprintf(validation.StepResolveFailed, validation.StepMatchesManyToolchains)

		fooFactory := test.NewFakeStepFactory(name, false, false)
		provider := steps.NewDefaultStepProvider(fooFactory, fooFactory)
		fooConfigStep := make(config.Step)
		fooConfigStep[name] = name

		step, err := provider.CreateFrom(fooConfigStep)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("FactoryWithSameName", func(t *testing.T) {
		name := "foo"

		fooFactory := test.NewFakeStepFactory(name, false, false)
		provider := steps.NewDefaultStepProvider(fooFactory)
		fooConfigStep := make(config.Step)
		fooConfigStep[name] = name

		step, err := provider.CreateFrom(fooConfigStep)

		test.Assert(t, step != nil)
		test.Assert(t, err == nil)
	})

	t.Run("FactoryWithSameName_StepConstructionFails", func(t *testing.T) {
		name := "foo"
		expected := fmt.Sprintf(validation.StepFailed, name, test.Error)

		fooFactory := test.NewFakeStepFactory(name, false, true)
		provider := steps.NewDefaultStepProvider(fooFactory)
		fooConfigStep := make(config.Step)
		fooConfigStep[name] = name

		step, err := provider.CreateFrom(fooConfigStep)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})
}
