package command_step

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/steps"
	"testing"
)

func TestPopulateOptions(t *testing.T) {
	options := steps.NewProviderOptions()

	PopulateOptions(options)

	assert.NotEmptyf(t, options.Executors[Name], "no executor with name %s", Name)
	assert.NotEmptyf(t, options.Factories[Name], "no factory with name %s", Name)
}

func TestFactory(t *testing.T) {
	t.Run("InvalidInput", func(t *testing.T) {
		test := func(t *testing.T, definition interface{}) {
			step, err := newStep(definition)

			assert.Empty(t, step)
			assert.EqualError(t, err, StepIsInvalid)
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

		assert.NotEmpty(t, step)
		assert.Equal(t, Name, step.Name())
		assert.IsType(t, &commandStep{cmd: cmd}, step)
		assert.NoError(t, err)
	})
}
