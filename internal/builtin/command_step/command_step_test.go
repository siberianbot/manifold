package command_step

import (
	"manifold/internal/steps"
	"testing"
)

func TestPopulateOptions(t *testing.T) {
	options := steps.NewProviderOptions()

	PopulateOptions(options)

	if options.Executors[Name] == nil {
		t.Errorf("no executor with name %s", Name)
	}

	if options.Factories[Name] == nil {
		t.Errorf("no factory with name %s", Name)
	}
}

func TestFactory(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		test := func(t *testing.T, definition interface{}) {
			step, err := newStep(definition)

			if step != nil {
				t.Error("step is not nil")
			}

			if err == nil {
				t.Error("no error")
			} else if err.Error() != StepIsInvalid {
				t.Errorf("error is %s, not %s", err.Error(), StepIsInvalid)
			}
		}

		t.Run("Number", func(t *testing.T) { test(t, 42) })
		t.Run("Boolean", func(t *testing.T) { test(t, true) })
		t.Run("Map", func(t *testing.T) { test(t, make(map[interface{}]interface{})) })
		t.Run("Slice", func(t *testing.T) { test(t, make([]interface{}, 3)) })
		t.Run("Struct", func(t *testing.T) { test(t, struct{}{}) })
		t.Run("EmptyString", func(t *testing.T) { test(t, "") })
	})

	t.Run("ValidInput", func(t *testing.T) {
		cmd := "foo"

		step, err := newStep(cmd)

		if step == nil {
			t.Error("step is nil")
		} else if step.Name() != Name {
			t.Errorf("step's name is %s, not %s", step.Name(), Name)
		} else {
			cmdStep, ok := step.(*commandStep)

			if !ok {
				t.Error("step is not commandStep ptr")
			} else if cmdStep.cmd != cmd {
				t.Errorf("step.cmd is %s, not %s", cmdStep.cmd, cmd)
			}
		}

		if err != nil {
			t.Errorf("error is not nil: %v", err)
		}
	})
}
