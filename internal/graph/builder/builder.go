package builder

import (
	"manifold/internal/errors"
	"manifold/internal/graph"
	"manifold/internal/graph/node"
	nodeBuilder "manifold/internal/graph/node/builder"
	nodeUtils "manifold/internal/graph/node/utils"
	graphValidator "manifold/internal/graph/validator"
)

const (
	unknownNamedDependency = "unknown target \"%s\""
)

type Interface interface {
	FromPath(path string) (graph.Interface, error)
}

type builder struct {
	nodeBuilder    nodeBuilder.Interface
	graphValidator graphValidator.Interface
}

func NewBuilder(nodeBuilder nodeBuilder.Interface, graphValidator graphValidator.Interface) Interface {
	return &builder{nodeBuilder: nodeBuilder, graphValidator: graphValidator}
}

func (b *builder) FromPath(path string) (graph.Interface, error) {
	root, buildRootErr := b.nodeBuilder.FromPath(path)

	if buildRootErr != nil {
		return nil, buildRootErr
	}

	g := graph.NewGraph(root)

	if processingErr := b.process(root, g); processingErr != nil {
		return nil, processingErr
	}

	if validationErr := b.graphValidator.Validate(g); validationErr != nil {
		return nil, validationErr
	}

	return g, nil
}

func (b *builder) process(n node.Node, g graph.Interface) error {
	newNodes := nodeUtils.NewNodeQueue()
	namedDeps := make([]node.Dependency, 0)

	for _, dep := range n.Dependencies() {
		switch dep.Kind() {
		case node.NamedDependencyKind:
			namedDeps = append(namedDeps, dep)

		case node.PathedDependencyKind:
			depNode := g.FindByPath(dep.Value())

			if depNode != nil {
				g.AddDescendant(n, depNode)
				return nil
			}

			depNode, buildDepNodeErr := b.nodeBuilder.FromPath(dep.Value())

			if buildDepNodeErr != nil {
				return buildDepNodeErr
			}

			newNodes.Enqueue(depNode)
			g.AddNode(depNode)
			g.AddDescendant(n, depNode)

			return nil

		default:
			panic("not supported dependency")
		}
	}

	for newNodes.Len() > 0 {
		newNode := newNodes.Dequeue()

		if err := b.process(newNode, g); err != nil {
			return err
		}
	}

	for _, namedDep := range namedDeps {
		namedDepNode := g.FindByName(namedDep.Value())

		if namedDepNode == nil {
			return errors.NewError(unknownNamedDependency, namedDep.Value())
		}

		g.AddDescendant(n, namedDepNode)
	}

	return nil
}
