package graph

import "errors"

// see also: Tarjan's strongly connected components algorithm
// https://en.wikipedia.org/wiki/Tarjan%27s_strongly_connected_components_algorithm

type cycle []*Node

type cycleContext struct {
	graph  *DependencyGraph
	stack  []*Node
	index  int
	cycles []cycle
}

func strongConnect(v *Node, ctx *cycleContext) {
	v.index = ctx.index
	v.lowLink = ctx.index

	ctx.index++

	ctx.stack = append(ctx.stack, v)
	v.onStack = true

	for _, w := range ctx.graph.FindDescendants(v) {
		if w.index == -1 {
			strongConnect(w, ctx)
			v.lowLink = min(v.lowLink, w.lowLink)
		} else if w.onStack {
			v.lowLink = min(v.lowLink, w.index)
		}
	}

	if v.lowLink == v.index {
		cycle := make(cycle, 0)

		for {
			w := ctx.stack[len(ctx.stack)-1]
			ctx.stack = ctx.stack[:len(ctx.stack)-1]

			w.onStack = false
			cycle = append(cycle, w)

			if v == w {
				break
			}
		}

		if len(cycle) > 1 {
			ctx.cycles = append(ctx.cycles, cycle)
		}
	}
}

func detectCycle(g *DependencyGraph) error {
	ctx := new(cycleContext)
	ctx.graph = g
	ctx.stack = make([]*Node, 0)
	ctx.index = 0
	ctx.cycles = make([]cycle, 0)

	// detect self-references before Tarjan's algorithm...
	for _, link := range g.Links {
		if link.Parent == link.Child {
			ctx.cycles = append(ctx.cycles, cycle{link.Parent})
		}
	}

	// detect complex cycles with Tarjan's algorithm...
	for _, node := range g.Nodes {
		if node.index == -1 {
			strongConnect(node, ctx)
		}
	}

	if len(ctx.cycles) > 0 {
		// TODO: handle error properly
		return errors.New("dependency graph contains cycles")
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
