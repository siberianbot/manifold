package validation

import "manifold/internal/graph"

type nodeStack struct {
	container []graph.Node
}

func newNodeStack() *nodeStack {
	return &nodeStack{
		container: make([]graph.Node, 0),
	}
}

func (stack *nodeStack) push(node graph.Node) {
	stack.container = append(stack.container, node)
}

func (stack *nodeStack) pop() (item graph.Node) {
	if len(stack.container) == 0 {
		return nil
	}

	lastIdx := len(stack.container) - 1
	item = stack.container[lastIdx]
	stack.container = stack.container[:lastIdx]

	return
}
