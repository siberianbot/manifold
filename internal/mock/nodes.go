package mock

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

type NodeBuilder struct {
	mock.Mock
}

func (m *NodeBuilder) FromConfig(path string, cfg *config.Configuration) (node.Node, error) {
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
