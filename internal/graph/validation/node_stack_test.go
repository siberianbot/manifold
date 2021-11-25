package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeStack(t *testing.T) {
	t.Run("NewNodeStack", func(t *testing.T) {
		nodeStack := newNodeStack()

		assert.NotEmpty(t, nodeStack)
		assert.NotNil(t, nodeStack.container)
		assert.Empty(t, nodeStack.container)
	})

	t.Run("PushPop", func(t *testing.T) {
		nodeStack := newNodeStack()

		first := &testNode{}
		nodeStack.push(first)

		assert.Len(t, nodeStack.container, 1)
		assert.Equal(t, first, nodeStack.container[0])

		second := &testNode{}
		nodeStack.push(second)

		assert.Len(t, nodeStack.container, 2)
		assert.Equal(t, first, nodeStack.container[0])
		assert.Equal(t, second, nodeStack.container[1])

		popSecond := nodeStack.pop()
		assert.Len(t, nodeStack.container, 1)
		assert.Equal(t, first, nodeStack.container[0])
		assert.Equal(t, second, popSecond)

		popFirst := nodeStack.pop()
		assert.Empty(t, nodeStack.container)
		assert.Equal(t, first, popFirst)
	})
}
