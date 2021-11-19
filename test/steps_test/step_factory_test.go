package steps_test

import (
	"manifold/internal/document_definition"
	"manifold/internal/steps"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyDefinition", func(t *testing.T) {
		ctx := test.NewFakeContext()
		stepDefinition := document_definition.StepDefinition{}

		step := steps.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == validation.StepNotMatch)
	})

	t.Run("ManyDefinitions", func(t *testing.T) {
		ctx := test.NewFakeContext()
		stepDefinition := document_definition.StepDefinition{
			"foo": "foo",
			"bar": "bar",
		}

		step := steps.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == validation.StepWithManyToolchains)
	})

	t.Run("Cmd", testCmd)
}

func testCmd(t *testing.T) {
	t.Run("ValidCmd", func(t *testing.T) {
		command := "foo bar baz"

		ctx := test.NewFakeContext()
		stepDefinition := document_definition.StepDefinition{
			"cmd": command,
		}

		step := steps.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, len(ctx.Errors) == 0)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == steps.CommandStepKind)
		test.Assert(t, step.(steps.CommandStep).Command == command)
	})

	t.Run("InvalidCmd", func(t *testing.T) {
		command := map[string]string{
			"incorrect": "cmd",
		}

		ctx := test.NewFakeContext()
		stepDefinition := document_definition.StepDefinition{
			"cmd": command,
		}

		step := steps.FromStepDefinition(stepDefinition, &ctx)

		test.Assert(t, step == nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == validation.CmdStepIsInvalid)
	})
}
