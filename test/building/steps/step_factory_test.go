package steps

import (
	"manifold/internal/building/steps"
	dd "manifold/internal/document_definition"
	"manifold/test"
	"manifold/test/building"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyCmd", func(t *testing.T) {
		context := building.NewFakeTraverseContext()

		stepDefinition := dd.StepDefinition{
			Command: "",
		}

		step := steps.FromStepDefinition(&stepDefinition, &context)

		test.Assert(t, len(context.Errors) > 0)
		test.Assert(t, step == nil)
	})

	t.Run("NotEmptyCmd", func(t *testing.T) {
		context := building.NewFakeTraverseContext()

		command := "foo bar baz"
		stepDefinition := dd.StepDefinition{
			Command: command,
		}

		step := steps.FromStepDefinition(&stepDefinition, &context)

		test.Assert(t, len(context.Errors) == 0)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == steps.CommandStepKind)
		test.Assert(t, step.(steps.CommandStep).Command == command)
	})
}
