package validation

import (
	"fmt"
	"manifold/internal/graph"
	"manifold/internal/steps"
)

type testNode struct {
	num int
}

func (n testNode) Build(provider *steps.Provider) error {
	return nil
}

func (n testNode) IsBuilt() bool {
	return false
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
