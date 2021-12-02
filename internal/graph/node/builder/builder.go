package builder

import (
	configLoader "manifold/internal/config/loader"
	"manifold/internal/graph/node"
)

type Interface interface {
	FromPath(path string) (node.Node, error)
}

type builder struct {
	configLoader         configLoader.Interface
	projectNodeBuilder   node.Builder
	workspaceNodeBuilder node.Builder
}

func NewBuilder(configLoader configLoader.Interface, projectNodeBuilder node.Builder, workspaceNodeBuilder node.Builder) Interface {
	return &builder{
		configLoader:         configLoader,
		projectNodeBuilder:   projectNodeBuilder,
		workspaceNodeBuilder: workspaceNodeBuilder,
	}
}

func (b *builder) FromPath(path string) (node.Node, error) {
	cfg, cfgErr := b.configLoader.FromPath(path)

	if cfgErr != nil {
		return nil, cfgErr
	}

	switch {
	case cfg.Project != nil:
		return b.projectNodeBuilder.FromConfig(path, cfg)

	case cfg.Workspace != nil:
		return b.workspaceNodeBuilder.FromConfig(path, cfg)

	default:
		panic("configuration is invalid")
	}
}
