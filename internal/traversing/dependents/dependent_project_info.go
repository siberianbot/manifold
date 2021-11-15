package dependents

import (
	"manifold/internal/traversing"
	"manifold/internal/validation"
)

type DependentProjectInfo struct {
	Project string
}

func (DependentProjectInfo) Kind() DependentInfoKind {
	return DependentProjectInfoKind
}

func newDependentProject(project string, ctx traversing.Context) DependentInfo {
	err := validation.ValidateManifoldName(project)

	if err != nil {
		ctx.AddError(validation.InvalidDependentProject, err)
		return nil
	}

	return DependentProjectInfo{
		Project: project,
	}
}
