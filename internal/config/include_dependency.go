package config

import (
	"manifold/internal/utils"
	"manifold/internal/validation"
)

type includeDependency struct {
	path string
}

func (i *includeDependency) Validate(ctx validation.Context) error {
	return validation.ValidatePath(utils.BuildPath(ctx.Dir(), i.path))
}

func (i *includeDependency) Kind() DependencyKind {
	return PathDependencyKind
}

func (i *includeDependency) Value() string {
	return i.path
}

func newInclude(path string) *includeDependency {
	include := new(includeDependency)
	include.path = path

	return include
}
