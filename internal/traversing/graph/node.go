package graph

import (
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/dependents"
)

type Node struct {
	BuildInfo    build_info.BuildInfo
	Dependencies []dependents.DependentInfo

	// for cycle detection algorithm
	index   int
	lowLink int
	onStack bool
}

func newNode(buildInfo build_info.BuildInfo, dependencies []dependents.DependentInfo) *Node {
	node := new(Node)

	node.BuildInfo = buildInfo
	node.Dependencies = dependencies

	// for cycle detection algorithm
	node.index = -1
	node.lowLink = -1
	node.onStack = false

	return node
}
