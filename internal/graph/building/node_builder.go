package building

import (
	"manifold/internal/config"
	"manifold/internal/config/utils"
	"manifold/internal/config/validation"
	"manifold/internal/graph"
	"os"
)

type NodeBuilder struct {
	//
}

func NewNodeBuilder() *NodeBuilder {
	return &NodeBuilder{}
}

func (builder *NodeBuilder) FromPath(path string) (graph.Node, error) {
	path, pathErr := utils.ResolvePath(path)

	if pathErr != nil {
		return nil, pathErr
	}

	file, fileErr := os.Open(path)
	defer func() { _ = file.Close() }()

	if fileErr != nil {
		return nil, fileErr
	}

	cfg, cfgErr := config.Read(file)

	if cfgErr != nil {
		return nil, cfgErr
	}

	validationErr := validation.Validate(cfg, path)

	if validationErr != nil {
		return nil, validationErr
	}

	return graph.FromConfiguration(cfg, path), nil
}