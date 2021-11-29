package building

import (
	"manifold/internal/config/utils"
	"manifold/internal/errors"
	"manifold/internal/graph"
	"manifold/internal/graph/validation"
)

const (
	UnknownTarget = "target \"%s\" not found"
)

type GraphBuilder struct {
	options GraphBuilderOptions
}

func NewGraphBuilder(options GraphBuilderOptions) *GraphBuilder {
	return &GraphBuilder{options: options}
}

func (builder *GraphBuilder) Build(path string) (*graph.DependencyGraph, error) {
	root, processRootErr := builder.options.NodeBuilder.FromPath(path)

	if processRootErr != nil {
		return nil, processRootErr
	}

	dependencyGraph := graph.NewDependencyGraph(root)

	processNodeErr := builder.processNode(root, dependencyGraph)

	if processNodeErr != nil {
		return nil, processNodeErr
	}

	cycleErr := validation.NewCycleDetector(dependencyGraph).Validate()

	if cycleErr != nil {
		return nil, cycleErr
	}

	return dependencyGraph, nil
}

func (builder *GraphBuilder) processNode(node graph.Node, dependencyGraph *graph.DependencyGraph) error {
	newNodes := make([]graph.Node, 0)
	namedDependencies := make([]graph.NodeDependency, 0)

	for _, dependency := range node.Dependencies() {
		switch dependency.Kind() {
		case graph.ByNameDependencyKind:
			namedDependencies = append(namedDependencies, dependency)

		case graph.ByPathDependencyKind:
			dependencyPath, dependencyPathErr := utils.ResolvePath(dependency.Value())

			if dependencyPathErr != nil {
				return dependencyPathErr
			}

			dependencyNode := dependencyGraph.FindByPath(dependencyPath)

			if dependencyNode != nil {
				dependencyGraph.AddDescendant(node, dependencyNode)
				continue
			}

			dependencyNode, dependencyNodeErr := builder.options.NodeBuilder.FromPath(dependencyPath)

			if dependencyNodeErr != nil {
				return dependencyNodeErr
			}

			newNodes = append(newNodes, dependencyNode)
			dependencyGraph.AddNode(dependencyNode)
			dependencyGraph.AddDescendant(node, dependencyNode)

		default:
			panic("unknown dependency kind")
		}
	}

	for _, newNode := range newNodes {
		processErr := builder.processNode(newNode, dependencyGraph)

		if processErr != nil {
			return processErr
		}
	}

	for _, namedDependency := range namedDependencies {
		dependencyNode := dependencyGraph.FindByName(namedDependency.Value())

		if dependencyNode == nil {
			return errors.NewError(UnknownTarget, namedDependency.Value())
		}

		dependencyGraph.AddDescendant(node, dependencyNode)
	}

	return nil
}
