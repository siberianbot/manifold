package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAggregateError(t *testing.T) {
	aerr := errors.New("a")
	berr := errors.New("b")
	err := NewAggregateError("msg", aerr, berr)

	assert.EqualError(t, err, "msg\na\nb")
}
