package graph

type DependencyGraph struct {
	Root  *Node
	Nodes []*Node
	Links []*Link
}

func (g *DependencyGraph) FindNodeByName(name string) *Node {
	for _, node := range g.Nodes {
		if node.Target.Name() == name {
			return node
		}
	}

	return nil
}

func (g *DependencyGraph) FindNodeByPath(path string) *Node {
	for _, node := range g.Nodes {
		if node.Target.Path() == path {
			return node
		}
	}

	return nil
}

func (g *DependencyGraph) FindLink(parent *Node, child *Node) *Link {
	for _, link := range g.Links {
		if link.Parent == parent && link.Child == child {
			return link
		}
	}

	return nil
}

func (g *DependencyGraph) FindDescendants(node *Node) []*Node {
	descendants := make([]*Node, 0)

	for _, link := range g.Links {
		if link.Parent == node {
			descendants = append(descendants, link.Child)
		}
	}

	return descendants
}

func newDependencyGraph() *DependencyGraph {
	dependencyGraph := new(DependencyGraph)

	dependencyGraph.Root = nil
	dependencyGraph.Nodes = make([]*Node, 0)
	dependencyGraph.Links = make([]*Link, 0)

	return dependencyGraph
}
