package workspace_node

import (
	"manifold/internal/config"
	configUtils "manifold/internal/config/utils"
	"manifold/internal/errors"
	"manifold/internal/graph/node"
	"manifold/internal/utils"
	"path/filepath"
)

const (
	invalidWorkspaceInclude = "workspace \"%s\" contain invalid include \"%s\": %v"
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

	dir := filepath.Dir(path)

	for idx, workspaceInclude := range cfg.Workspace.Includes {
		includePath := utils.BuildPath(dir, string(workspaceInclude))
		resolvedPath, err := configUtils.ResolvePath(includePath)

		if err != nil {
			return nil, errors.NewError(invalidWorkspaceInclude, cfg.Workspace.Name, string(workspaceInclude), err)
		}

		workspaceNode.includes[idx] = node.NewDependency(node.PathedDependencyKind, resolvedPath)
	}

	return workspaceNode, nil
}
