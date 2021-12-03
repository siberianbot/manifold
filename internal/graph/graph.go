package graph

import "manifold/internal/graph/node"

type Interface interface {
	Root() node.Node

	Nodes() []node.Node
	AddNode(n node.Node)

	FindByName(name string) node.Node
	FindByPath(path string) node.Node

	DescendantsOf(n node.Node) []node.Node
	AddDescendant(parent node.Node, child node.Node)
}

type graph struct {
	root  node.Node
	nodes []node.Node
	links []link
}

type link struct {
	parent node.Node
	child  node.Node
}

func NewGraph(root node.Node) Interface {
	return &graph{
		root:  root,
		nodes: []node.Node{root},
		links: make([]link, 0),
	}
}

func (g *graph) Root() node.Node {
	return g.root
}

func (g *graph) Nodes() []node.Node {
	return g.nodes
}

func (g *graph) DescendantsOf(n node.Node) []node.Node {
	descendants := make([]node.Node, 0)

	for _, link := range g.links {
		if link.parent != n {
			continue
		}

		descendants = append(descendants, link.child)
	}

	return descendants
}

func (g *graph) FindByName(name string) node.Node {
	for _, n := range g.nodes {
		if n.Name() == name {
			return n
		}
	}

	return nil
}

func (g *graph) FindByPath(path string) node.Node {
	for _, n := range g.nodes {
		if n.Path() == path {
			return n
		}
	}

	return nil
}

func (g *graph) AddNode(n node.Node) {
	g.nodes = append(g.nodes, n)
}

func (g *graph) AddDescendant(parent node.Node, child node.Node) {
	g.links = append(g.links, link{parent: parent, child: child})
}
