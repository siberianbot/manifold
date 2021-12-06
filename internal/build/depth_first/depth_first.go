package depth_first

import (
	"fmt"
	"manifold/internal/build"
	"manifold/internal/errors"
	"manifold/internal/graph"
	"manifold/internal/graph/node"
	"manifold/internal/step"
	stepProvider "manifold/internal/step/provider"
	"path/filepath"
)

const (
	buildFailed = "failed to build \"%s\": %v"
)

type depthFirstBuild struct {
	stepProvider stepProvider.Interface
}

func NewDepthFirstBuildStrategy(stepProvider stepProvider.Interface) build.Interface {
	return &depthFirstBuild{stepProvider: stepProvider}
}

func (b *depthFirstBuild) Build(g graph.Interface) error {
	buildContext := &buildContext{graph: g, built: make(map[node.Node]bool)}

	for _, n := range g.Nodes() {
		buildContext.built[n] = false
	}

	return b.build(g.Root(), buildContext)
}

func (b *depthFirstBuild) build(n node.Node, buildCtx *buildContext) error {
	if buildCtx.built[n] {
		return nil
	}

	for _, desc := range buildCtx.graph.DescendantsOf(n) {
		err := b.build(desc, buildCtx)

		if err != nil {
			return err
		}
	}

	for _, s := range n.Steps() {
		executor := b.stepProvider.ExecutorFor(s.Name())

		if executor == nil {
			panic(fmt.Sprintf("no executor for %s", s.Name()))
		}

		ctx := &executorContext{
			step: s,
			node: n,
		}

		err := executor.Execute(ctx)

		if err != nil {
			return errors.NewError(buildFailed, n.Name(), err)
		}
	}

	buildCtx.built[n] = true

	return nil
}

type buildContext struct {
	built map[node.Node]bool
	graph graph.Interface
}

type executorContext struct {
	node node.Node
	step step.Step
}

func (e *executorContext) Step() step.Step {
	return e.step
}

func (e *executorContext) Dir() string {
	return filepath.Dir(e.node.Path())
}
