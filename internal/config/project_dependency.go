package config

import (
	"errors"
	"manifold/internal/utils"
	"manifold/internal/validation"
)

type projectDependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

func (p *projectDependency) Validate(ctx validation.Context) error {
	switch {
	case p.Project != "" && p.Path != "":
		return errors.New(validation.DependencyWithBothProjectAndPath)

	case p.Project != "":
		return validation.ValidateManifoldName(p.Project)

	case p.Path != "":
		return validation.ValidatePath(utils.BuildPath(ctx.Dir(), p.Path))

	default:
		return errors.New(validation.EmptyProjectDependency)
	}
}

func (p *projectDependency) Kind() DependencyKind {
	switch {
	case p.Project != "":
		return ProjectDependencyKind

	case p.Path != "":
		return PathDependencyKind

	default:
		return UnknownDependencyKind
	}
}

func (p *projectDependency) Value() string {
	switch p.Kind() {
	case ProjectDependencyKind:
		return p.Project

	case PathDependencyKind:
		return p.Path

	default:
		panic("unsupported dependency kind")
	}
}
