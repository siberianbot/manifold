package validation

import "manifold/internal/graph"

type nodeData struct {
	index    int
	lowIndex int
	onStack  bool
	selfRef  bool
}

func newNodeData() *nodeData {
	return &nodeData{index: -1, lowIndex: -1, onStack: false, selfRef: false}
}

type nodeDataMap map[graph.Node]*nodeData

func newNodeDataMap() nodeDataMap {
	return make(map[graph.Node]*nodeData)
}
