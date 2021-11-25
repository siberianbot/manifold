package validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/graph"
	"strings"
	"testing"
)

func TestCycleDetector(t *testing.T) {
	t.Run("NewCycleDetector", func(t *testing.T) {
		root := &testNode{}
		dependencyGraph := graph.NewDependencyGraph(root)

		cycleDetector := NewCycleDetector(dependencyGraph)

		assert.NotEmpty(t, cycleDetector)
		assert.Equal(t, dependencyGraph, cycleDetector.graph)
		assert.Equal(t, 0, cycleDetector.idx)
		assert.NotNil(t, cycleDetector.cycles)
		assert.Empty(t, cycleDetector.cycles)
		assert.NotNil(t, cycleDetector.data)
		assert.Empty(t, cycleDetector.data)
		assert.NotNil(t, cycleDetector.stack)
	})

	t.Run("Validation", func(t *testing.T) {
		t.Run("SingleNode", func(t *testing.T) {
			root := newTestNode(1)
			dependencyGraph := graph.NewDependencyGraph(root)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("RootAndOneDescendant", func(t *testing.T) {
			root := newTestNode(1)
			descendant := newTestNode(2)
			dependencyGraph := graph.NewDependencyGraph(root)
			dependencyGraph.AddNode(descendant)
			dependencyGraph.AddDescendant(root, descendant)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("RootAndManyDescendant", func(t *testing.T) {
			root := newTestNode(1)
			descendant1 := newTestNode(2)
			descendant2 := newTestNode(3)
			descendant3 := newTestNode(4)
			dependencyGraph := graph.NewDependencyGraph(root)
			dependencyGraph.AddNode(descendant1)
			dependencyGraph.AddDescendant(root, descendant1)
			dependencyGraph.AddNode(descendant2)
			dependencyGraph.AddDescendant(root, descendant2)
			dependencyGraph.AddNode(descendant3)
			dependencyGraph.AddDescendant(root, descendant3)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("SimpleChain", func(t *testing.T) {
			first := newTestNode(1)
			second := newTestNode(2)
			third := newTestNode(3)
			dependencyGraph := graph.NewDependencyGraph(first)
			dependencyGraph.AddNode(second)
			dependencyGraph.AddNode(third)
			dependencyGraph.AddDescendant(first, second)
			dependencyGraph.AddDescendant(second, third)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("SimpleTree", func(t *testing.T) {
			first := newTestNode(1)
			second := newTestNode(2)
			third := newTestNode(3)
			fourth := newTestNode(4)
			dependencyGraph := graph.NewDependencyGraph(first)
			dependencyGraph.AddNode(second)
			dependencyGraph.AddNode(third)
			dependencyGraph.AddNode(fourth)
			dependencyGraph.AddDescendant(first, second)
			dependencyGraph.AddDescendant(first, third)
			dependencyGraph.AddDescendant(second, fourth)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("OneCommonNode", func(t *testing.T) {
			first := newTestNode(1)
			second := newTestNode(2)
			third := newTestNode(3)
			dependencyGraph := graph.NewDependencyGraph(first)
			dependencyGraph.AddNode(second)
			dependencyGraph.AddNode(third)
			dependencyGraph.AddDescendant(first, second)
			dependencyGraph.AddDescendant(first, third)
			dependencyGraph.AddDescendant(second, third)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("RhombusWithOneCommonNode", func(t *testing.T) {
			first := newTestNode(1)
			second := newTestNode(2)
			third := newTestNode(3)
			fourth := newTestNode(4)
			dependencyGraph := graph.NewDependencyGraph(first)
			dependencyGraph.AddNode(second)
			dependencyGraph.AddNode(third)
			dependencyGraph.AddNode(fourth)
			dependencyGraph.AddDescendant(first, second)
			dependencyGraph.AddDescendant(first, third)
			dependencyGraph.AddDescendant(second, fourth)
			dependencyGraph.AddDescendant(third, fourth)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			assert.NoError(t, err)
		})

		t.Run("SelfReference", func(t *testing.T) {
			node := newTestNode(1)
			dependencyGraph := graph.NewDependencyGraph(node)
			dependencyGraph.AddDescendant(node, node)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			expected := fmt.Sprintf(CyclesDetected,
				fmt.Sprintf(SelfReference, node.Name(), node.Path()))
			assert.EqualError(t, err, expected)
		})

		t.Run("TwoNodesCycle", func(t *testing.T) {
			a := newTestNode(1)
			b := newTestNode(2)
			dependencyGraph := graph.NewDependencyGraph(a)
			dependencyGraph.AddNode(b)
			dependencyGraph.AddDescendant(a, b)
			dependencyGraph.AddDescendant(b, a)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			expected := fmt.Sprintf(CyclesDetected,
				fmt.Sprintf(Cycle, strings.Join([]string{
					fmt.Sprintf(CycleEntry, a.Name(), a.Path()),
					fmt.Sprintf(CycleEntry, b.Name(), b.Path()),
					fmt.Sprintf(CycleEntry, a.Name(), a.Path()),
				}, Separator)))
			assert.EqualError(t, err, expected)
		})

		t.Run("ThreeNodesCycle", func(t *testing.T) {
			a := newTestNode(1)
			b := newTestNode(2)
			c := newTestNode(3)
			dependencyGraph := graph.NewDependencyGraph(a)
			dependencyGraph.AddNode(b)
			dependencyGraph.AddNode(c)
			dependencyGraph.AddDescendant(a, b)
			dependencyGraph.AddDescendant(b, c)
			dependencyGraph.AddDescendant(c, a)
			cycleDetector := NewCycleDetector(dependencyGraph)

			err := cycleDetector.Validate()

			expected := fmt.Sprintf(CyclesDetected,
				fmt.Sprintf(Cycle, strings.Join([]string{
					fmt.Sprintf(CycleEntry, a.Name(), a.Path()),
					fmt.Sprintf(CycleEntry, b.Name(), b.Path()),
					fmt.Sprintf(CycleEntry, c.Name(), c.Path()),
					fmt.Sprintf(CycleEntry, a.Name(), a.Path()),
				}, Separator)))
			assert.EqualError(t, err, expected)
		})
	})
}
