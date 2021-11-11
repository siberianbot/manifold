package building_steps

import (
	"manifold/internal/building_steps"
	dd "manifold/internal/document_definition"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("EmptyCmd", func(t *testing.T) {
		stepDefinition := dd.StepDefinition{
			Command: "",
		}

		step, err := building_steps.FromStepDefinition(&stepDefinition)

		test.AssertNil(t, step)
		test.AssertNotNil(t, err)
	})

	t.Run("NotEmptyCmd", func(t *testing.T) {
		command := "foo bar baz"
		stepDefinition := dd.StepDefinition{
			Command: command,
		}

		step, err := building_steps.FromStepDefinition(&stepDefinition)

		test.AssertNil(t, err)
		test.Assert(t, step != nil)
		test.Assert(t, step.Kind() == building_steps.CommandStepKind)
		test.Assert(t, step.(building_steps.CommandStep).Command == command)
	})
}
