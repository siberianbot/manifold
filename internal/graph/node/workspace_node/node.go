package workspace_node

import (
	"manifold/internal/config"
	"manifold/internal/graph/node"
)

type Node struct {
	workspace *config.Workspace
	path      string
	includes  []node.Dependency
}

func (n *Node) Name() string {
	return n.workspace.Name
}

func (n *Node) Path() string {
	return n.path
}

func (n *Node) Dependencies() []node.Dependency {
	return n.includes
}
