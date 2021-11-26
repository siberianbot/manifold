package graph

import (
	"manifold/internal/config"
	"manifold/internal/steps"
	"path/filepath"
)

type Node interface {
	Path() string
	Name() string
	Dependencies() []NodeDependency
	Build(provider *steps.Provider) error
	IsBuilt() bool
}

func FromConfiguration(cfg *config.Configuration, path string, provider *steps.Provider) (Node, error) {
	switch {
	case cfg.ProjectTarget != nil:
		return newProjectNode(cfg.ProjectTarget, path, provider)

	case cfg.WorkspaceTarget != nil:
		return newWorkspaceNode(cfg.WorkspaceTarget, path), nil

	default:
		panic("configuration is empty")
	}
}

type projectNode struct {
	path    string
	project *config.ProjectTarget
	steps   []steps.Step
	isBuilt bool
}

func newProjectNode(project *config.ProjectTarget, path string, provider *steps.Provider) (*projectNode, error) {
	node := &projectNode{
		project: project,
		path:    path,
		isBuilt: false,
		steps:   make([]steps.Step, len(project.Steps)),
	}

	for idx, configStep := range project.Steps {
		step, err := provider.CreateFrom(configStep)

		if err != nil {
			return nil, err
		}

		node.steps[idx] = step
	}

	return node, nil
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

func (node *projectNode) Build(stepsProvider *steps.Provider) error {
	for _, step := range node.steps {
		err := stepsProvider.Execute(step)

		if err != nil {
			return err
		}
	}

	node.isBuilt = true
	return nil
}

func (node *projectNode) IsBuilt() bool {
	return node.isBuilt
}

type workspaceNode struct {
	path      string
	workspace *config.WorkspaceTarget
	isBuilt   bool
}

func newWorkspaceNode(workspace *config.WorkspaceTarget, path string) *workspaceNode {
	return &workspaceNode{workspace: workspace, path: path, isBuilt: false}
}

func (node *workspaceNode) Path() string {
	return node.path
}

func (node *workspaceNode) Name() string {
	return node.workspace.Name
}

func (node *workspaceNode) Build(_ *steps.Provider) error {
	node.isBuilt = true
	return nil
}

func (node *workspaceNode) IsBuilt() bool {
	return node.isBuilt
}

func (node *workspaceNode) Dependencies() []NodeDependency {
	dir := filepath.Dir(node.path)
	dependencies := make([]NodeDependency, len(node.workspace.Includes))

	for idx, dependency := range node.workspace.Includes {
		dependencies[idx] = FromWorkspaceInclude(dependency, dir)
	}

	return dependencies
}
