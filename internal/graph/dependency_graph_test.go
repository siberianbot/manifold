package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDependencyGraph(t *testing.T) {
	t.Run("NewDependencyGraph", func(t *testing.T) {
		node := &testNode{}
		graph := NewDependencyGraph(node)

		assert.NotEmpty(t, graph)
		assert.Equal(t, node, graph.root)
		assert.NotEmpty(t, graph.nodes)
		assert.Len(t, graph.nodes, 1)
		assert.Equal(t, node, graph.nodes[0])
		assert.NotNil(t, graph.links)
		assert.Empty(t, graph.links)
	})

	t.Run("Root", func(t *testing.T) {
		node := &testNode{}
		graph := NewDependencyGraph(node)

		assert.Equal(t, node, graph.Root())
	})

	t.Run("WithoutDescendants", func(t *testing.T) {
		root := &testNode{}
		graph := NewDependencyGraph(root)

		descendants := graph.Descendants(root)
		assert.NotNil(t, descendants)
		assert.Empty(t, descendants)
	})

	t.Run("WithDescendants", func(t *testing.T) {
		root := &testNode{}
		descendant := &testNode{}
		graph := NewDependencyGraph(root)

		graph.AddNode(descendant)
		assert.Len(t, graph.nodes, 2)

		graph.AddDescendant(root, descendant)
		assert.Len(t, graph.links, 1)

		descendants := graph.Descendants(root)
		assert.NotEmpty(t, descendants)
		assert.Len(t, descendants, 1)
		assert.Equal(t, descendant, descendants[0])
	})

	t.Run("FoundByName", func(t *testing.T) {
		root := &testNode{}
		graph := NewDependencyGraph(root)

		node := graph.FindByName(root.Name())
		assert.NotNil(t, node)
		assert.Equal(t, root, node)
	})

	t.Run("NotFoundByName", func(t *testing.T) {
		root := &testNode{}
		graph := NewDependencyGraph(root)

		node := graph.FindByName("foo" + root.Name())
		assert.Nil(t, node)
	})

	t.Run("FoundByPath", func(t *testing.T) {
		root := &testNode{}
		graph := NewDependencyGraph(root)

		node := graph.FindByPath(root.Path())
		assert.NotNil(t, node)
		assert.Equal(t, root, node)
	})

	t.Run("NotFoundByPath", func(t *testing.T) {
		root := &testNode{}
		graph := NewDependencyGraph(root)

		node := graph.FindByPath("foo" + root.Path())
		assert.Nil(t, node)
	})
}
