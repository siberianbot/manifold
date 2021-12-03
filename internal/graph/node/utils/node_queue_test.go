package utils

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/mock/node"
	"testing"
)

func TestNodeQueue(t *testing.T) {
	queue := NewNodeQueue()

	assert.NotNil(t, queue)
	assert.Equal(t, 0, queue.Len())

	first := new(node.Node)

	queue.Enqueue(first)

	assert.Equal(t, 1, queue.Len())

	second := new(node.Node)

	queue.Enqueue(second)

	assert.Equal(t, 2, queue.Len())

	dequeueFirst := queue.Dequeue()

	assert.Equal(t, first, dequeueFirst)
	assert.Equal(t, 1, queue.Len())

	dequeueSecond := queue.Dequeue()

	assert.Equal(t, second, dequeueSecond)
	assert.Equal(t, 0, queue.Len())
}
