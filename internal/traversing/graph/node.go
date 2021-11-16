package graph

import (
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/dependents"
)

type Node struct {
	BuildInfo    build_info.BuildInfo
	Dependencies []dependents.DependentInfo
}

type NodeLink struct {
	Parent *Node
	Child  *Node
}

type NodeCollection struct {
	Root  *Node
	Nodes []*Node
	Links []*NodeLink
}

func newNode(buildInfo build_info.BuildInfo, dependencies []dependents.DependentInfo) *Node {
	node := new(Node)

	node.BuildInfo = buildInfo
	node.Dependencies = dependencies

	return node
}

func newNodeLink(parent *Node, child *Node) *NodeLink {
	nodeLink := new(NodeLink)

	nodeLink.Parent = parent
	nodeLink.Child = child

	return nodeLink
}

func newNodeCollection() *NodeCollection {
	nodeCollection := new(NodeCollection)

	nodeCollection.Root = nil
	nodeCollection.Nodes = make([]*Node, 0)
	nodeCollection.Links = make([]*NodeLink, 0)

	return nodeCollection
}
