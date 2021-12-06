package project_node

import (
	"manifold/internal/config"
	"manifold/internal/graph/node"
	"manifold/internal/step"
)

type Node struct {
	project      *config.Project
	path         string
	steps        []step.Step
	dependencies []node.Dependency
}

func (n *Node) Name() string {
	return n.project.Name
}

func (n *Node) Path() string {
	return n.path
}

func (n *Node) Dependencies() []node.Dependency {
	return n.dependencies
}

func (n *Node) Steps() []step.Step {
	return n.steps
}
