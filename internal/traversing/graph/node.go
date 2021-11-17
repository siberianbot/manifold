package graph

import (
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/dependents"
)

type Node struct {
	BuildInfo    build_info.BuildInfo
	Dependencies []dependents.DependentInfo
}

func newNode(buildInfo build_info.BuildInfo, dependencies []dependents.DependentInfo) *Node {
	node := new(Node)

	node.BuildInfo = buildInfo
	node.Dependencies = dependencies

	return node
}
