package validation

import (
	"fmt"
	"manifold/internal/errors"
	"manifold/internal/graph"
	"strings"
)

const (
	CyclesDetected = "cycles detected:\n%s"
	Cycle          = "cycle:\n%s"
	Separator      = "\n-> "
	CycleEntry     = "%s at %s"
	SelfReference  = "self reference: %s at %s"
)

type cycle []graph.Node

type CycleDetector struct {
	graph  *graph.DependencyGraph
	data   nodeDataMap
	stack  *nodeStack
	cycles []cycle
	idx    int
}

func NewCycleDetector(dependencyGraph *graph.DependencyGraph) *CycleDetector {
	return &CycleDetector{
		graph:  dependencyGraph,
		data:   newNodeDataMap(),
		stack:  newNodeStack(),
		cycles: make([]cycle, 0),
		idx:    0,
	}
}

func (detector *CycleDetector) getDataFor(node graph.Node) *nodeData {
	data := detector.data[node]

	if data == nil {
		data = newNodeData()
		detector.data[node] = data
	}

	return data
}

func (detector *CycleDetector) process(v graph.Node) {
	detector.stack.push(v)

	vdata := detector.getDataFor(v)
	vdata.index = detector.idx
	vdata.lowIndex = detector.idx
	vdata.onStack = true

	detector.idx++

	for _, w := range detector.graph.Descendants(v) {
		if w == v {
			vdata.selfRef = true
		}

		wdata := detector.getDataFor(w)

		if wdata.index == -1 {
			detector.process(w)
			vdata.lowIndex = min(vdata.lowIndex, wdata.lowIndex)
		} else if wdata.onStack {
			vdata.lowIndex = min(vdata.lowIndex, wdata.index)
		}
	}

	if vdata.index == vdata.lowIndex {
		cycle := make(cycle, 0)

		for {
			w := detector.stack.pop()
			wdata := detector.getDataFor(w)
			wdata.onStack = false

			cycle = append(cycle, w)

			if w == v {
				break
			}
		}

		if len(cycle) > 1 || vdata.selfRef {
			detector.cycles = append(detector.cycles, cycle)
		}
	}
}

func (detector *CycleDetector) Validate() error {
	for _, node := range detector.graph.Nodes() {
		nodeData := detector.getDataFor(node)

		if nodeData.index == -1 {
			detector.process(node)
		}
	}

	if len(detector.cycles) > 0 {
		cyclesStrs := make([]string, 0)

		for _, cycle := range detector.cycles {
			if len(cycle) == 1 {
				node := cycle[0]
				cyclesStrs = append(cyclesStrs, fmt.Sprintf(SelfReference, node.Name(), node.Path()))
			} else {
				entries := make([]string, len(cycle))

				for idx, node := range reverse(cycle) {
					entries[idx] = fmt.Sprintf(CycleEntry, node.Name(), node.Path())
				}

				entries = append(entries, entries[0])

				cyclesStrs = append(cyclesStrs, fmt.Sprintf(Cycle, strings.Join(entries, Separator)))
			}
		}

		return errors.NewError(CyclesDetected, strings.Join(cyclesStrs, "\n"))
	}

	return nil
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func reverse(slice []graph.Node) []graph.Node {
	sliceLen := len(slice)
	reversedSlice := make([]graph.Node, sliceLen)

	for i := range slice {
		reversedSlice[sliceLen-1-i] = slice[i]
	}

	return reversedSlice
}
