package build

import "manifold/internal/graph"

type Interface interface {
	Build(g graph.Interface) error
}
