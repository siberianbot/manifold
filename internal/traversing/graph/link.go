package graph

type Link struct {
	Parent *Node
	Child  *Node
}

func newLink(parent *Node, child *Node) *Link {
	link := new(Link)

	link.Parent = parent
	link.Child = child

	return link
}
