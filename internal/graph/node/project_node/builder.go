package project_node

import (
	"manifold/internal/config"
	"manifold/internal/errors"
	"manifold/internal/graph/node"
	"manifold/internal/step"
	stepBuilder "manifold/internal/step/builder"
)

const (
	invalidProjectStep = "project \"%s\" contain invalid step: %v"
)

type builder struct {
	stepBuilder stepBuilder.Interface
}

func NewBuilder(stepBuilder stepBuilder.Interface) node.Builder {
	return &builder{stepBuilder: stepBuilder}
}

func (b *builder) FromConfig(path string, cfg *config.Configuration) (node.Node, error) {
	projectNode := &Node{
		path:         path,
		project:      cfg.Project,
		steps:        make([]step.Step, len(cfg.Project.Steps)),
		dependencies: make([]node.Dependency, len(cfg.Project.Dependencies)),
	}

	for idx, stepDefinition := range cfg.Project.Steps {
		s, err := b.stepBuilder.FromConfig(stepDefinition)

		if err != nil {
			return nil, errors.NewError(invalidProjectStep, cfg.Project.Name, err)
		}

		projectNode.steps[idx] = s
	}

	for idx, projectDependency := range cfg.Project.Dependencies {
		var kind node.DependencyKind
		var value string

		switch {
		case projectDependency.Project != "":
			kind, value = node.NamedDependencyKind, projectDependency.Project

		case projectDependency.Path != "":
			kind, value = node.PathedDependencyKind, projectDependency.Path

		default:
			panic("project dependency is invalid")
		}

		projectNode.dependencies[idx] = node.NewDependency(kind, value)
	}

	return projectNode, nil
}
