package workspace_node

import (
	"manifold/internal/config"
	"manifold/internal/graph/node"
)

type builder struct {
	//
}

func NewBuilder() node.Builder {
	return &builder{}
}

func (builder) FromConfig(path string, cfg *config.Configuration) (node.Node, error) {
	workspaceNode := &Node{
		path:      path,
		workspace: cfg.Workspace,
		includes:  make([]node.Dependency, len(cfg.Workspace.Includes)),
	}

	for idx, workspaceInclude := range cfg.Workspace.Includes {
		workspaceNode.includes[idx] = node.NewDependency(node.PathedDependencyKind, string(workspaceInclude))
	}

	return workspaceNode, nil
}
