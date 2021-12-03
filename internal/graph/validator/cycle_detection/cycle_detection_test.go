package cycle_detection

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/graph"
	"manifold/internal/mock/node"
	"testing"
)

func TestSingleNode(t *testing.T) {
	root := new(node.Node)
	g := graph.NewGraph(root)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestRootAndOneDescendant(t *testing.T) {
	root := new(node.Node)
	descendant := new(node.Node)
	g := graph.NewGraph(root)
	g.AddNode(descendant)
	g.AddDescendant(root, descendant)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestRootAndManyDescendant(t *testing.T) {
	root := new(node.Node)
	descendant1 := new(node.Node)
	descendant2 := new(node.Node)
	descendant3 := new(node.Node)
	descendant4 := new(node.Node)
	g := graph.NewGraph(root)
	g.AddNode(descendant1)
	g.AddNode(descendant2)
	g.AddNode(descendant3)
	g.AddNode(descendant4)
	g.AddDescendant(root, descendant1)
	g.AddDescendant(root, descendant2)
	g.AddDescendant(root, descendant3)
	g.AddDescendant(root, descendant4)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestSimpleChain(t *testing.T) {
	first := new(node.Node)
	second := new(node.Node)
	third := new(node.Node)
	g := graph.NewGraph(first)
	g.AddNode(second)
	g.AddNode(third)
	g.AddDescendant(first, second)
	g.AddDescendant(second, third)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestSimpleTree(t *testing.T) {
	first := new(node.Node)
	second := new(node.Node)
	third := new(node.Node)
	fourth := new(node.Node)
	g := graph.NewGraph(first)
	g.AddNode(second)
	g.AddNode(third)
	g.AddNode(fourth)
	g.AddDescendant(first, second)
	g.AddDescendant(first, third)
	g.AddDescendant(second, fourth)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestOneCommonNode(t *testing.T) {
	first := new(node.Node)
	second := new(node.Node)
	third := new(node.Node)
	g := graph.NewGraph(first)
	g.AddNode(second)
	g.AddNode(third)
	g.AddDescendant(first, second)
	g.AddDescendant(first, third)
	g.AddDescendant(second, third)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestRhombusWithOneCommonNode(t *testing.T) {
	first := new(node.Node)
	second := new(node.Node)
	third := new(node.Node)
	fourth := new(node.Node)
	g := graph.NewGraph(first)
	g.AddNode(second)
	g.AddNode(third)
	g.AddNode(fourth)
	g.AddDescendant(first, second)
	g.AddDescendant(first, third)
	g.AddDescendant(second, fourth)
	g.AddDescendant(third, fourth)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 0)
}

func TestSelfReference(t *testing.T) {
	n := new(node.Node)
	g := graph.NewGraph(n)
	g.AddDescendant(n, n)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 1)
	assert.Len(t, cycles[0], 1)
	assert.Contains(t, cycles[0], n)
}

func TestTwoNodesCycle(t *testing.T) {
	a := new(node.Node)
	b := new(node.Node)
	g := graph.NewGraph(a)
	g.AddNode(b)
	g.AddDescendant(a, b)
	g.AddDescendant(b, a)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 1)
	assert.Len(t, cycles[0], 2)
	assert.Equal(t, cycles[0][0], a)
	assert.Equal(t, cycles[0][1], b)
}

func TestThreeNodesCycle(t *testing.T) {
	a := new(node.Node)
	b := new(node.Node)
	c := new(node.Node)
	g := graph.NewGraph(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddDescendant(a, b)
	g.AddDescendant(b, c)
	g.AddDescendant(c, a)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 1)
	assert.Len(t, cycles[0], 3)
	assert.Equal(t, cycles[0][0], a)
	assert.Equal(t, cycles[0][1], b)
	assert.Equal(t, cycles[0][2], c)
}

func TestFullCycle(t *testing.T) {
	a := new(node.Node)
	b := new(node.Node)
	c := new(node.Node)
	g := graph.NewGraph(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddDescendant(a, b)
	g.AddDescendant(b, a)
	g.AddDescendant(b, c)
	g.AddDescendant(c, b)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 1)
	assert.Len(t, cycles[0], 3)
	assert.Equal(t, cycles[0][0], a)
	assert.Equal(t, cycles[0][1], b)
	assert.Equal(t, cycles[0][2], c)
}

func TestMultipleCycles(t *testing.T) {
	a := new(node.Node)
	b := new(node.Node)
	c := new(node.Node)
	d := new(node.Node)
	e := new(node.Node)
	g := graph.NewGraph(a)
	g.AddNode(b)
	g.AddNode(c)
	g.AddNode(d)
	g.AddNode(e)
	g.AddDescendant(a, b)
	g.AddDescendant(b, c)
	g.AddDescendant(c, a)
	g.AddDescendant(c, d)
	g.AddDescendant(d, e)
	g.AddDescendant(e, d)

	cycles := GetDefaultCycleDetectionAlgorithm()(g)

	assert.Len(t, cycles, 2)
	t.Logf("a = %p", a)
	t.Logf("b = %p", b)
	t.Logf("c = %p", c)
	t.Logf("d = %p", d)
	t.Logf("e = %p", e)
	t.Log("cycle 1", cycles[0])
	t.Log("cycle 2", cycles[1])

	if len(cycles[0]) == 2 {
		assert.Len(t, cycles[0], 2)
		assert.Equal(t, cycles[0][0], d)
		assert.Equal(t, cycles[0][1], e)
		assert.Len(t, cycles[1], 3)
		assert.Equal(t, cycles[1][0], a)
		assert.Equal(t, cycles[1][1], b)
		assert.Equal(t, cycles[1][2], c)
	} else if len(cycles[0]) == 3 {
		assert.Len(t, cycles[0], 3)
		assert.Equal(t, cycles[0][0], a)
		assert.Equal(t, cycles[0][1], b)
		assert.Equal(t, cycles[0][2], c)
		assert.Len(t, cycles[1], 2)
		assert.Equal(t, cycles[1][0], d)
		assert.Equal(t, cycles[1][1], e)
	}
}
