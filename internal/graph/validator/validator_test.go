package validator

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/graph"
	"manifold/internal/graph/node"
	"manifold/internal/graph/validator/cycle_detection"
	graph2 "manifold/internal/mock/graph"
	node2 "manifold/internal/mock/node"
	"testing"
)

func TestNoCycles(t *testing.T) {
	algorithmCalled := false
	algorithmFn := func(_ graph.Interface) []cycle_detection.Cycle {
		algorithmCalled = true
		return []cycle_detection.Cycle{}
	}

	validator := NewValidator(algorithmFn)

	assert.NotNil(t, validator)

	g := new(graph2.Graph)

	err := validator.Validate(g)

	assert.True(t, algorithmCalled)
	assert.NoError(t, err)
}

func TestSingleCycle(t *testing.T) {
	a := new(node2.Node)
	a.On("Name").Return("a")
	b := new(node2.Node)
	b.On("Name").Return("b")

	algorithmCalled := false
	algorithmFn := func(_ graph.Interface) []cycle_detection.Cycle {
		algorithmCalled = true
		return []cycle_detection.Cycle{
			[]node.Node{a, b},
		}
	}

	validator := NewValidator(algorithmFn)

	assert.NotNil(t, validator)

	g := new(graph2.Graph)

	err := validator.Validate(g)

	assert.True(t, algorithmCalled)
	assert.EqualError(t, err, "cycles detected:\n\tcycle: a -> b")
}

func TestSelfReference(t *testing.T) {
	a := new(node2.Node)
	a.On("Name").Return("a")

	algorithmCalled := false
	algorithmFn := func(_ graph.Interface) []cycle_detection.Cycle {
		algorithmCalled = true
		return []cycle_detection.Cycle{
			[]node.Node{a},
		}
	}

	validator := NewValidator(algorithmFn)

	assert.NotNil(t, validator)

	g := new(graph2.Graph)

	err := validator.Validate(g)

	assert.True(t, algorithmCalled)
	assert.EqualError(t, err, "cycles detected:\n\tproject \"a\" self references")
}

func TestManyCycles(t *testing.T) {
	a := new(node2.Node)
	a.On("Name").Return("a")
	b := new(node2.Node)
	b.On("Name").Return("b")
	c := new(node2.Node)
	c.On("Name").Return("c")
	d := new(node2.Node)
	d.On("Name").Return("d")

	algorithmCalled := false
	algorithmFn := func(_ graph.Interface) []cycle_detection.Cycle {
		algorithmCalled = true
		return []cycle_detection.Cycle{
			[]node.Node{a, b},
			[]node.Node{b, c},
			[]node.Node{c, d},
		}
	}

	validator := NewValidator(algorithmFn)

	assert.NotNil(t, validator)

	g := new(graph2.Graph)

	err := validator.Validate(g)

	assert.True(t, algorithmCalled)
	assert.EqualError(t, err, "cycles detected:\n\tcycle: a -> b\n\tcycle: b -> c\n\tcycle: c -> d")
}
