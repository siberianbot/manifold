package graph

import "manifold/internal/steps"

type testNode struct {
	//
}

func (n testNode) Build(provider *steps.Provider) error {
	return nil
}

func (n testNode) IsBuilt() bool {
	return false
}

func (testNode) Path() string {
	return "path"
}

func (testNode) Name() string {
	return "name"
}

func (testNode) Dependencies() []NodeDependency {
	return make([]NodeDependency, 0)
}
