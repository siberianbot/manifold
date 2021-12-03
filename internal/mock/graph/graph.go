package graph

import (
	"github.com/stretchr/testify/mock"
	"manifold/internal/graph"
	"manifold/internal/graph/node"
)

type Graph struct {
	mock.Mock
}

func (g *Graph) Root() node.Node {
	args := g.Called()

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(node.Node)
}

func (g *Graph) Nodes() []node.Node {
	args := g.Called()

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).([]node.Node)
}

func (g *Graph) AddNode(n node.Node) {
	g.Called(n)
}

func (g *Graph) FindByName(name string) node.Node {
	args := g.Called(name)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(node.Node)
}

func (g *Graph) FindByPath(path string) node.Node {
	args := g.Called(path)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).(node.Node)
}

func (g *Graph) DescendantsOf(n node.Node) []node.Node {
	args := g.Called(n)

	if args.Get(0) == nil {
		return nil
	}

	return args.Get(0).([]node.Node)
}

func (g *Graph) AddDescendant(parent node.Node, child node.Node) {
	g.Called(parent, child)
}

type GraphValidator struct {
	mock.Mock
}

func (g *GraphValidator) Validate(graph graph.Interface) error {
	args := g.Called(graph)

	if args.Get(0) == nil {
		return nil
	}

	return args.Error(0)
}
