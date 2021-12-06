package project_node

import (
	"manifold/internal/config"
	configUtils "manifold/internal/config/utils"
	"manifold/internal/errors"
	"manifold/internal/graph/node"
	"manifold/internal/step"
	stepBuilder "manifold/internal/step/builder"
	"manifold/internal/utils"
	"path/filepath"
)

const (
	invalidProjectStep       = "project \"%s\" contain invalid step: %v"
	invalidProjectDependency = "project \"%s\" contain invalid dependency \"%s\": %v"
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

	dir := filepath.Dir(path)

	for idx, projectDependency := range cfg.Project.Dependencies {
		var kind node.DependencyKind
		var value string

		switch {
		case projectDependency.Project != "":
			kind, value = node.NamedDependencyKind, projectDependency.Project

		case projectDependency.Path != "":
			depPath := utils.BuildPath(dir, projectDependency.Path)
			resolvedPath, err := configUtils.ResolvePath(depPath)

			if err != nil {
				return nil, errors.NewError(invalidProjectDependency, cfg.Project.Name, projectDependency.Path, err)
			}

			kind, value = node.PathedDependencyKind, resolvedPath

		default:
			panic("project dependency is invalid")
		}

		projectNode.dependencies[idx] = node.NewDependency(kind, value)
	}

	return projectNode, nil
}
