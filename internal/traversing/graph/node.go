package graph

import (
	"manifold/internal/traversing/dependents"
	"manifold/internal/traversing/targets"
)

type Node struct {
	Target       targets.Target
	Dependencies []dependents.DependentInfo

	// for cycle detection algorithm
	index   int
	lowLink int
	onStack bool
}

func newNode(target targets.Target, dependencies []dependents.DependentInfo) *Node {
	node := new(Node)

	node.Target = target
	node.Dependencies = dependencies

	// for cycle detection algorithm
	node.index = -1
	node.lowLink = -1
	node.onStack = false

	return node
}
