package builder

import (
	"errors"
	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
	"manifold/internal/graph/node"
	"manifold/internal/mock/graph"
	mockNode "manifold/internal/mock/node"
	"path/filepath"
	"testing"
)

func TestInvalidRoot(t *testing.T) {
	path := "path"
	nodeBuilderErr := errors.New("error")

	nodeBuilder := new(mockNode.NodeBuilder)
	nodeBuilder.On("FromPath", path).Return(nil, nodeBuilderErr)
	validator := new(graph.GraphValidator)

	builder := NewBuilder(nodeBuilder, validator)

	assert.NotNil(t, builder)

	g, err := builder.FromPath(path)

	assert.Nil(t, g)
	assert.EqualError(t, err, "error")

	nodeBuilder.AssertExpectations(t)
	validator.AssertExpectations(t)
}

func TestRootWithoutDependencies(t *testing.T) {
	root := new(mockNode.Node)
	root.On("Name").Return("root")
	root.On("Path").Return("root")
	root.On("Dependencies").Return([]node.Dependency{})

	nodeBuilder := new(mockNode.NodeBuilder)
	nodeBuilder.On("FromPath", "root").Return(root, nil)
	validator := new(graph.GraphValidator)
	validator.On("Validate", testifyMock.Anything).Return(nil)

	builder := NewBuilder(nodeBuilder, validator)

	assert.NotNil(t, builder)

	g, err := builder.FromPath("root")

	assert.NotNil(t, g)
	assert.Equal(t, root, g.Root())
	assert.NoError(t, err)

	nodeBuilder.AssertExpectations(t)
	validator.AssertExpectations(t)
}

func TestInvalidGraph(t *testing.T) {
	root := new(mockNode.Node)
	root.On("Name").Return("root")
	root.On("Path").Return("root")
	root.On("Dependencies").Return([]node.Dependency{})

	validatorErr := errors.New("error")

	nodeBuilder := new(mockNode.NodeBuilder)
	nodeBuilder.On("FromPath", "root").Return(root, nil)
	validator := new(graph.GraphValidator)
	validator.On("Validate", testifyMock.Anything).Return(validatorErr)

	builder := NewBuilder(nodeBuilder, validator)

	assert.NotNil(t, builder)

	g, err := builder.FromPath("root")

	assert.Nil(t, g)
	assert.EqualError(t, err, "error")

	nodeBuilder.AssertExpectations(t)
	validator.AssertExpectations(t)
}

func TestFooWithBar(t *testing.T) {
	barDepPath := filepath.Join(".", "bar")
	barDep := new(mockNode.Dependency)
	barDep.On("Kind").Return(node.PathedDependencyKind)
	barDep.On("Value").Return(barDepPath)

	foo := new(mockNode.Node)
	foo.On("Name").Return("foo")
	foo.On("Path").Return("foo")
	foo.On("Dependencies").Return([]node.Dependency{barDep})

	bar := new(mockNode.Node)
	bar.On("Name").Return("bar")
	bar.On("Path").Return("bar")
	bar.On("Dependencies").Return([]node.Dependency{})

	nodeBuilder := new(mockNode.NodeBuilder)
	nodeBuilder.On("FromPath", "foo").Return(foo, nil)
	nodeBuilder.On("FromPath", barDepPath).Return(bar, nil)
	validator := new(graph.GraphValidator)
	validator.On("Validate", testifyMock.Anything).Return(nil)

	builder := NewBuilder(nodeBuilder, validator)

	assert.NotNil(t, builder)

	g, err := builder.FromPath("foo")

	assert.NotNil(t, g)
	assert.Equal(t, foo, g.Root())
	assert.Contains(t, g.Nodes(), foo)
	assert.Contains(t, g.Nodes(), bar)
	assert.Contains(t, g.DescendantsOf(foo), bar)
	assert.NoError(t, err)

	nodeBuilder.AssertExpectations(t)
	validator.AssertExpectations(t)
}
