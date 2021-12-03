package cycle_detection

import (
	"manifold/internal/graph"
	"manifold/internal/graph/node"
	"manifold/internal/graph/node/utils"
)

type Cycle []node.Node

type CycleDetectionAlgorithm func(graph graph.Interface) []Cycle

func GetDefaultCycleDetectionAlgorithm() CycleDetectionAlgorithm {
	return detectCycles
}

func detectCycles(graph graph.Interface) []Cycle {
	state := &cycleDetectionState{
		graph:  graph,
		stack:  utils.NewNodeStack(),
		data:   make(nodeDataMap),
		idx:    0,
		cycles: make([]Cycle, 0),
	}

	for _, n := range graph.Nodes() {
		ndata := state.data.getFor(n)

		if ndata.index == -1 {
			processNode(n, state)
		}
	}

	return state.cycles
}

type nodeData struct {
	index    int
	lowIndex int
	onStack  bool
	selfRef  bool
}

func newNodeData() *nodeData {
	return &nodeData{index: -1, lowIndex: -1, onStack: false, selfRef: false}
}

type nodeDataMap map[node.Node]*nodeData

func (m nodeDataMap) contains(node node.Node) bool {
	return m[node] != nil
}

func (m nodeDataMap) getFor(node node.Node) (data *nodeData) {
	if m.contains(node) {
		data = m[node]
	} else {
		m[node] = newNodeData()
		data = m[node]
	}

	return data
}

type cycleDetectionState struct {
	graph  graph.Interface
	data   nodeDataMap
	idx    int
	stack  *utils.NodeStack
	cycles []Cycle
}

func processNode(v node.Node, state *cycleDetectionState) {
	state.stack.Push(v)

	vdata := state.data.getFor(v)
	vdata.index = state.idx
	vdata.lowIndex = state.idx
	vdata.onStack = true

	state.idx++

	for _, w := range state.graph.DescendantsOf(v) {
		if w == v {
			vdata.selfRef = true
		}

		wdata := state.data.getFor(w)

		if wdata.index == -1 {
			processNode(w, state)
			vdata.lowIndex = min(vdata.lowIndex, wdata.lowIndex)
		} else if wdata.onStack {
			vdata.lowIndex = min(vdata.lowIndex, wdata.index)
		}
	}

	if vdata.index == vdata.lowIndex {
		cycle := make(Cycle, 0)

		for {
			w := state.stack.Pop()
			wdata := state.data.getFor(w)
			wdata.onStack = false

			cycle = append(cycle, w)

			if w == v {
				break
			}
		}

		if len(cycle) > 1 || vdata.selfRef {
			state.cycles = append(state.cycles, cycle)
		}
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
