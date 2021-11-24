package graph

import (
	"manifold/internal/config"
	"path/filepath"
)

type Node interface {
	Path() string
	Name() string
	Dependencies() []NodeDependency
}

func FromConfiguration(cfg *config.Configuration, path string) Node {
	switch {
	case cfg.ProjectTarget != nil:
		return newProjectNode(cfg.ProjectTarget, path)

	case cfg.WorkspaceTarget != nil:
		return newWorkspaceNode(cfg.WorkspaceTarget, path)

	default:
		panic("configuration is empty")
	}
}

type projectNode struct {
	path    string
	project *config.ProjectTarget
}

func newProjectNode(project *config.ProjectTarget, path string) *projectNode {
	return &projectNode{project: project, path: path}
}

func (node *projectNode) Path() string {
	return node.path
}

func (node *projectNode) Name() string {
	return node.project.Name
}

func (node *projectNode) Dependencies() []NodeDependency {
	dir := filepath.Dir(node.path)
	dependencies := make([]NodeDependency, len(node.project.ProjectDependencies))

	for idx, dependency := range node.project.ProjectDependencies {
		dependencies[idx] = FromProjectDependency(dependency, dir)
	}

	return dependencies
}

type workspaceNode struct {
	path      string
	workspace *config.WorkspaceTarget
}

func newWorkspaceNode(workspace *config.WorkspaceTarget, path string) *workspaceNode {
	return &workspaceNode{workspace: workspace, path: path}
}

func (node *workspaceNode) Path() string {
	return node.path
}

func (node *workspaceNode) Name() string {
	return node.workspace.Name
}

func (node *workspaceNode) Dependencies() []NodeDependency {
	dir := filepath.Dir(node.path)
	dependencies := make([]NodeDependency, len(node.workspace.Includes))

	for idx, dependency := range node.workspace.Includes {
		dependencies[idx] = FromWorkspaceInclude(dependency, dir)
	}

	return dependencies
}
