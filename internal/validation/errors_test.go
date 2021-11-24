package validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Run("MessageOnly", func(t *testing.T) {
		msg := "foo"
		err := NewError(msg)

		assert.EqualError(t, err, msg)
	})

	t.Run("MessageWithArgs", func(t *testing.T) {
		msg := "foo %s %s"
		arg1 := "bar"
		arg2 := "baz"

		err := NewError(msg, arg1, arg2)

		assert.EqualError(t, err, fmt.Sprintf(msg, arg1, arg2))
	})
}
