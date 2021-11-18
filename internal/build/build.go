package build

import (
	"errors"
	"fmt"
	"manifold/internal/traversing/graph"
	"manifold/internal/traversing/targets"
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

	if node.Target.Kind() == targets.ProjectTargetKind {
		projectTarget := node.Target.(targets.ProjectTarget)

		for _, step := range projectTarget.Steps {
			if stepErr := step.Execute(); stepErr != nil {
				return stepErr
			}
		}
	}

	ctx.builtNodes = append(ctx.builtNodes, node)

	return nil
}
