package validation

import (
	"fmt"
	"manifold/internal/graph"
)

type testNode struct {
	num int
}

func (n testNode) Path() string {
	return fmt.Sprintf("path%d", n.num)
}

func (n testNode) Name() string {
	return fmt.Sprintf("name%d", n.num)
}

func (testNode) Dependencies() []graph.NodeDependency {
	return make([]graph.NodeDependency, 0)
}

func newTestNode(num int) *testNode {
	return &testNode{num: num}
}
