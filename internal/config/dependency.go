package config

import (
	"errors"
	"manifold/internal/utils"
	"manifold/internal/validation"
)

type DependencyKind int8

const (
	UnknownDependencyKind DependencyKind = iota
	ProjectDependencyKind
	PathDependencyKind
)

type Dependency interface {
	Validatable

	Kind() DependencyKind
	Value() string
}

type dependency struct {
	Path    string `yaml:"path"`
	Project string `yaml:"project"`
}

func (d *dependency) Validate(ctx ValidationContext) error {
	switch {
	case d.Project != "" && d.Path != "":
		return errors.New(validation.DependencyWithBothProjectAndPath)

	case d.Project != "":
		return validation.ValidateManifoldName(d.Project)

	case d.Path != "":
		return validation.ValidatePath(utils.BuildPath(ctx.Dir(), d.Path))

	default:
		return errors.New(validation.EmptyProjectDependency)
	}
}

func (d *dependency) Kind() DependencyKind {
	switch {
	case d.Project != "":
		return ProjectDependencyKind

	case d.Path != "":
		return PathDependencyKind

	default:
		return UnknownDependencyKind
	}
}

func (d *dependency) Value() string {
	switch d.Kind() {
	case ProjectDependencyKind:
		return d.Project

	case PathDependencyKind:
		return d.Path

	default:
		panic("unsupported dependency kind")
	}
}
