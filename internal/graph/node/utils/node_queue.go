package utils

import "manifold/internal/graph/node"

type NodeQueue struct {
	container []node.Node
}

func NewNodeQueue() *NodeQueue {
	return &NodeQueue{container: make([]node.Node, 0)}
}

func (queue *NodeQueue) Enqueue(node node.Node) {
	queue.container = append(queue.container, node)
}

func (queue *NodeQueue) Dequeue() (node node.Node) {
	if len(queue.container) == 0 {
		return nil
	}

	node = queue.container[0]
	queue.container = queue.container[1:]

	return
}

func (queue *NodeQueue) Len() int {
	return len(queue.container)
}
