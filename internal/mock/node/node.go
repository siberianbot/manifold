package node

import (
	"github.com/stretchr/testify/mock"
	"manifold/internal/config"
	"manifold/internal/graph/node"
)

type Node struct {
	mock.Mock
}

func (m *Node) Name() string {
	return m.Called().String(0)
}

func (m *Node) Path() string {
	return m.Called().String(0)
}

func (m *Node) Dependencies() []node.Dependency {
	return m.Called().Get(0).([]node.Dependency)
}

type Dependency struct {
	mock.Mock
}

func (m *Dependency) Kind() node.DependencyKind {
	return m.Called().Get(0).(node.DependencyKind)
}

func (m *Dependency) Value() string {
	return m.Called().String(0)
}

type ConcreteNodeBuilder struct {
	mock.Mock
}

func (m *ConcreteNodeBuilder) FromConfig(path string, cfg *config.Configuration) (node.Node, error) {
	args := m.Called(path, cfg)

	var n node.Node

	if args.Get(0) == nil {
		n = nil
	} else {
		n = args.Get(0).(node.Node)
	}

	var err error

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Get(1).(error)
	}

	return n, err
}

type NodeBuilder struct {
	mock.Mock
}

func (m *NodeBuilder) FromPath(path string) (node.Node, error) {
	args := m.Called(path)

	var n node.Node

	if args.Get(0) == nil {
		n = nil
	} else {
		n = args.Get(0).(node.Node)
	}

	var err error

	if args.Get(1) == nil {
		err = nil
	} else {
		err = args.Get(1).(error)
	}

	return n, err
}
