package steps_test

import (
	"manifold/internal/steps/command_step"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestCommandStepFactory(t *testing.T) {
	t.Run("InstantiationCheck", func(t *testing.T) {
		factory := command_step.NewStepFactory()

		test.Assert(t, factory != nil)
		test.Assert(t, factory.Name() == "cmd")
	})

	t.Run("EmptyString", func(t *testing.T) {
		cmd := ""
		factory := command_step.NewStepFactory()

		step, err := factory.Construct(cmd)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == validation.CmdStepIsInvalid)
	})

	t.Run("InvalidStruct", func(t *testing.T) {
		cmd := struct{}{}
		factory := command_step.NewStepFactory()

		step, err := factory.Construct(cmd)

		test.Assert(t, step == nil)
		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == validation.CmdStepIsInvalid)
	})

	t.Run("ValidString", func(t *testing.T) {
		cmd := "foo"
		factory := command_step.NewStepFactory()

		step, err := factory.Construct(cmd)

		test.Assert(t, err == nil)
		test.Assert(t, step != nil)
	})
}
