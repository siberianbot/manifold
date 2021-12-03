package utils

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/mock/node"
	"testing"
)

func TestNodeStack(t *testing.T) {
	stack := NewNodeStack()

	assert.NotNil(t, stack)
	assert.Equal(t, 0, stack.Len())

	first := new(node.Node)

	stack.Push(first)

	assert.Equal(t, 1, stack.Len())

	second := new(node.Node)

	stack.Push(second)

	assert.Equal(t, 2, stack.Len())

	popSecond := stack.Pop()

	assert.Equal(t, 1, stack.Len())
	assert.Equal(t, second, popSecond)

	popFirst := stack.Pop()

	assert.Equal(t, 0, stack.Len())
	assert.Equal(t, first, popFirst)
}
