package graph

type link struct {
	parent Node
	child  Node
}

type DependencyGraph struct {
	root  Node
	nodes []Node
	links []link
}

func (d *DependencyGraph) Root() Node {
	return d.root
}

func (d *DependencyGraph) Descendants(node Node) []Node {
	descendants := make([]Node, 0)

	for _, link := range d.links {
		if link.parent != node {
			continue
		}

		descendants = append(descendants, link.child)
	}

	return descendants
}

func (d *DependencyGraph) FindByName(name string) Node {
	for _, node := range d.nodes {
		if node.Name() == name {
			return node
		}
	}

	return nil
}

func (d *DependencyGraph) FindByPath(path string) Node {
	for _, node := range d.nodes {
		if node.Path() == path {
			return node
		}
	}

	return nil
}

func (d *DependencyGraph) AddDescendant(parent Node, child Node) {
	d.nodes = append(d.nodes, child)
	d.links = append(d.links, link{parent: parent, child: child})
}

func NewDependencyGraph(root Node) *DependencyGraph {
	return &DependencyGraph{
		root:  root,
		nodes: []Node{root},
		links: make([]link, 0),
	}
}
