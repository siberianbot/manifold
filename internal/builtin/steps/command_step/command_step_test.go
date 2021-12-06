package command_step

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProxyFactoryWithInvalidInput(t *testing.T) {
	proxy := commandProxy{}

	test := func(t *testing.T, definition interface{}) {
		step, err := proxy.CreateFrom(definition)

		assert.Empty(t, step)
		assert.EqualError(t, err, stepIsInvalid)
	}

	t.Run("Number", func(t *testing.T) { test(t, 42) })
	t.Run("Boolean", func(t *testing.T) { test(t, true) })
	t.Run("Map", func(t *testing.T) { test(t, make(map[interface{}]interface{})) })
	t.Run("Slice", func(t *testing.T) { test(t, make([]interface{}, 3)) })
	t.Run("Struct", func(t *testing.T) { test(t, struct{}{}) })
	t.Run("EmptyString", func(t *testing.T) { test(t, "") })
}

func TestProxyFactoryWithValidInput(t *testing.T) {
	proxy := commandProxy{}
	cmd := "foo"

	step, err := proxy.CreateFrom(cmd)

	assert.NotEmpty(t, step)
	assert.Equal(t, name, step.Name())
	assert.IsType(t, &commandStep{cmd: cmd}, step)
	assert.NoError(t, err)
}
