package steps

import (
	"manifold/internal/building/steps"
	dd "manifold/internal/document_definition"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyCmd", func(t *testing.T) {
		stepDefinition := dd.StepDefinition{
			Command: "",
		}

		step, err := steps.FromStepDefinition(&stepDefinition)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
	})

	t.Run("NotEmptyCmd", func(t *testing.T) {
		command := "foo bar baz"
		stepDefinition := dd.StepDefinition{
			Command: command,
		}

		step, err := steps.FromStepDefinition(&stepDefinition)

		test.Assert(t, err == nil)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == steps.CommandStepKind)
		test.Assert(t, step.(steps.CommandStep).Command == command)
	})
}
