package validation_test

import (
	"fmt"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("MessageOnly", func(t *testing.T) {
		msg := "foo"
		err := validation.NewError(msg)

		test.Assert(t, err.Error() == msg)
	})

	t.Run("MessageWithArgs", func(t *testing.T) {
		msg := "foo %s %s"
		arg1 := "bar"
		arg2 := "baz"

		err := validation.NewError(msg, arg1, arg2)

		test.Assert(t, err.Error() == fmt.Sprintf(msg, arg1, arg2))
	})
}
