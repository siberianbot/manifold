package build

import (
	"errors"
	"fmt"
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/graph"
)

type buildContext struct {
	graph      *graph.DependencyGraph
	builtNodes []*graph.Node
}

func (ctx buildContext) isBuilt(node *graph.Node) bool {
	for _, builtNode := range ctx.builtNodes {
		if node == builtNode {
			return true
		}
	}

	return false
}

func Build(path string) error {
	dependencyGraph, traverseErr := graph.Traverse(path)

	if traverseErr != nil {
		return traverseErr.Error()
	}

	ctx := new(buildContext)
	ctx.graph = dependencyGraph
	ctx.builtNodes = make([]*graph.Node, 0)

	buildErr := build(dependencyGraph.Root, ctx)

	if buildErr != nil {
		return errors.New(fmt.Sprintf("build failed: %v", buildErr))
	}

	return nil
}

func build(node *graph.Node, ctx *buildContext) error {
	if ctx.isBuilt(node) {
		return nil
	}

	for _, dependency := range ctx.graph.FindDescendants(node) {
		if dependencyErr := build(dependency, ctx); dependencyErr != nil {
			return dependencyErr
		}
	}

	if node.BuildInfo.Kind() == build_info.ProjectBuildInfoKind {
		projectBuildInfo := node.BuildInfo.(build_info.ProjectBuildInfo)

		for _, step := range projectBuildInfo.Steps {
			if stepErr := step.Execute(); stepErr != nil {
				return stepErr
			}
		}
	}

	ctx.builtNodes = append(ctx.builtNodes, node)

	return nil
}
