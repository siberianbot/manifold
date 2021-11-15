package dependents

import (
	"manifold/internal/traversing"
	"manifold/internal/validation"
)

type DependentPathInfo struct {
	Path string
}

func (DependentPathInfo) Kind() DependentInfoKind {
	return DependentPathInfoKind
}

func newDependentPath(path string, ctx traversing.Context) DependentInfo {
	err := validation.ValidatePath(path)

	if err != nil {
		ctx.AddError(validation.InvalidDependentProject, err)
		return nil
	}

	return DependentPathInfo{
		Path: path,
	}
}
