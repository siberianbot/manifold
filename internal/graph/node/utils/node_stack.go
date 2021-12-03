package utils

import (
	"manifold/internal/graph/node"
)

type NodeStack struct {
	container []node.Node
}

func NewNodeStack() *NodeStack {
	return &NodeStack{
		container: make([]node.Node, 0),
	}
}

func (stack *NodeStack) Push(item node.Node) {
	stack.container = append(stack.container, item)
}

func (stack *NodeStack) Pop() (item node.Node) {
	if len(stack.container) == 0 {
		return nil
	}

	lastIdx := len(stack.container) - 1
	item = stack.container[lastIdx]
	stack.container = stack.container[:lastIdx]

	return
}

func (stack *NodeStack) Len() int {
	return len(stack.container)
}
