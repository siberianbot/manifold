package validation

import (
	"fmt"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("MessageOnly", func(t *testing.T) {
		msg := "foo"
		err := NewError(msg)

		if err.Error() != msg {
			t.Errorf("error is %v, not %s", err, msg)
		}
	})

	t.Run("MessageWithArgs", func(t *testing.T) {
		msg := "foo %s %s"
		arg1 := "bar"
		arg2 := "baz"

		err := NewError(msg, arg1, arg2)

		expected := fmt.Sprintf(msg, arg1, arg2)
		if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})
}
