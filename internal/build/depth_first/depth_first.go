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
	return b.build(g.Root(), g)
}

func (b *depthFirstBuild) build(n node.Node, g graph.Interface) error {
	for _, desc := range g.DescendantsOf(n) {
		err := b.build(desc, g)

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

	return nil
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
