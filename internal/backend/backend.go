package backend

import (
	"manifold/internal/builtin/command_step"
	"manifold/internal/graph"
	"manifold/internal/graph/building"
	"manifold/internal/steps"
)

type BuildOptions struct {
	Path string
}

type Backend interface {
	Build(options BuildOptions) error
}

type backend struct {
	stepsProvider *steps.Provider
	graphBuilder  *building.GraphBuilder
}

func NewBackend() Backend {
	stepsProviderOptions := steps.NewProviderOptions()
	command_step.PopulateOptions(stepsProviderOptions)

	stepsProvider := steps.NewProvider(stepsProviderOptions)
	graphBuilderOptions := building.GraphBuilderOptions{
		NodeBuilder: building.NewNodeBuilder(stepsProvider),
	}
	graphBuilder := building.NewGraphBuilder(graphBuilderOptions)

	return &backend{
		stepsProvider: stepsProvider,
		graphBuilder:  graphBuilder,
	}
}

func (b *backend) Build(options BuildOptions) error {
	dependencyGraph, graphErr := b.graphBuilder.Build(options.Path)

	if graphErr != nil {
		return graphErr
	}

	return b.buildNode(dependencyGraph.Root(), dependencyGraph)
}

func (b *backend) buildNode(node graph.Node, dependencyGraph *graph.DependencyGraph) error {
	if node.IsBuilt() {
		return nil
	}

	for _, descendant := range dependencyGraph.Descendants(node) {
		descendantBuildErr := b.buildNode(descendant, dependencyGraph)

		if descendantBuildErr != nil {
			return descendantBuildErr
		}
	}

	return node.Build(b.stepsProvider)
}
