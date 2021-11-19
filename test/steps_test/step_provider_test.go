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
		ctx := test.NewFakeContext()
		definition := make(config.StepDefinition)

		step := provider.CreateFor(definition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) == 1)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("NoFactories", func(t *testing.T) {
		provider := steps.NewDefaultStepProvider()
		ctx := test.NewFakeContext()

		name := "foo"
		definition := make(config.StepDefinition)
		definition[name] = name
		expected := fmt.Sprintf(validation.StepResolveFailed, fmt.Sprintf(validation.StepNotMatchedAnyToolchain, name))

		step := provider.CreateFor(definition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) == 1)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("FactoryWithDifferentName", func(t *testing.T) {
		name := "foo"

		barFactory := test.NewFakeStepFactory("bar", false, false)
		provider := steps.NewDefaultStepProvider(barFactory)
		ctx := test.NewFakeContext()

		fooDefinition := make(config.StepDefinition)
		fooDefinition[name] = name
		expected := fmt.Sprintf(validation.StepResolveFailed, fmt.Sprintf(validation.StepNotMatchedAnyToolchain, name))

		step := provider.CreateFor(fooDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) == 1)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("FactoriesWithSameName", func(t *testing.T) {
		name := "foo"

		fooFactory := test.NewFakeStepFactory(name, false, false)
		provider := steps.NewDefaultStepProvider(fooFactory, fooFactory)
		ctx := test.NewFakeContext()

		fooDefinition := make(config.StepDefinition)
		fooDefinition[name] = name

		expected := fmt.Sprintf(validation.StepResolveFailed, validation.StepMatchesManyToolchains)

		step := provider.CreateFor(fooDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) == 1)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("FactoryWithSameName", func(t *testing.T) {
		name := "foo"

		fooFactory := test.NewFakeStepFactory(name, false, false)
		provider := steps.NewDefaultStepProvider(fooFactory)
		ctx := test.NewFakeContext()

		fooDefinition := make(config.StepDefinition)
		fooDefinition[name] = name

		step := provider.CreateFor(fooDefinition, &ctx)

		test.Assert(t, step != nil)
		test.Assert(t, len(ctx.Errors) == 0)
	})

	t.Run("FactoryWithSameName_StepConstructionFails", func(t *testing.T) {
		name := "foo"
		expected := fmt.Sprintf(validation.StepFailed, name, test.Error)

		fooFactory := test.NewFakeStepFactory(name, false, true)
		provider := steps.NewDefaultStepProvider(fooFactory)
		ctx := test.NewFakeContext()

		fooDefinition := make(config.StepDefinition)
		fooDefinition[name] = name

		step := provider.CreateFor(fooDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) == 1)
		test.Assert(t, ctx.Errors[0] == expected)
	})
}
