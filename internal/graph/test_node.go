package graph

type testNode struct {
	//
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
