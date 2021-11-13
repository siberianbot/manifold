package steps

import (
	dd "manifold/internal/document_definition"
	steps2 "manifold/internal/steps"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyCmd", func(t *testing.T) {
		context := test.NewFakeTraverseContext()

		stepDefinition := dd.StepDefinition{
			Command: "",
		}

		step := steps2.FromStepDefinition(&stepDefinition, &context)

		test.Assert(t, len(context.Errors) > 0)
		test.Assert(t, step == nil)
	})

	t.Run("NotEmptyCmd", func(t *testing.T) {
		context := test.NewFakeTraverseContext()

		command := "foo bar baz"
		stepDefinition := dd.StepDefinition{
			Command: command,
		}

		step := steps2.FromStepDefinition(&stepDefinition, &context)

		test.Assert(t, len(context.Errors) == 0)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == steps2.CommandStepKind)
		test.Assert(t, step.(steps2.CommandStep).Command == command)
	})
}
