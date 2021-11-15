package steps_test

import (
	dd "manifold/internal/document_definition"
	steps2 "manifold/internal/steps"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyDefinition", func(t *testing.T) {
		ctx := test.NewFakeContext()
		stepDefinition := dd.StepDefinition{}

		step := steps2.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == validation.StepNotMatch)
	})

	t.Run("NotEmptyCmd", func(t *testing.T) {
		command := "foo bar baz"

		ctx := test.NewFakeContext()
		stepDefinition := dd.StepDefinition{
			Command: command,
		}

		step := steps2.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, len(ctx.Errors) == 0)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == steps2.CommandStepKind)
		test.Assert(t, step.(steps2.CommandStep).Command == command)
	})
}
