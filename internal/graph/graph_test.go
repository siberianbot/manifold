package graph

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/mock/node"
	"testing"
)

func TestGraph(t *testing.T) {
	root := new(node.Node)
	root.On("Name").Return("root")
	root.On("Path").Return("root")

	graph := NewGraph(root)

	assert.NotEmpty(t, graph)
	assert.Equal(t, root, graph.Root())

	descendant := new(node.Node)
	graph.AddDescendant(root, descendant)
	descendant.On("Name").Return("descendant")
	descendant.On("Path").Return("descendant")

	descendants := graph.DescendantsOf(root)
	assert.Len(t, descendants, 1)
	assert.Equal(t, descendant, descendants[0])

	byName := graph.FindByName("descendant")
	assert.Nil(t, byName)

	byPath := graph.FindByPath("descendant")
	assert.Nil(t, byPath)

	graph.AddNode(descendant)

	byName = graph.FindByName("descendant")
	assert.Equal(t, descendant, byName)

	byPath = graph.FindByPath("descendant")
	assert.Equal(t, descendant, byPath)
}
