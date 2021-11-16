package dependents

import (
	"manifold/internal/traversing"
	"manifold/internal/validation"
	"path/filepath"
)

type DependentPathInfo struct {
	Path string
}

func (DependentPathInfo) Kind() DependentInfoKind {
	return DependentPathInfoKind
}

func newDependentPath(path string, ctx traversing.Context) DependentInfo {
	dir := filepath.Dir(ctx.CurrentFile())
	path = filepath.Clean(filepath.Join(dir, path))

	err := validation.ValidatePath(path)

	if err != nil {
		ctx.AddError(validation.InvalidDependentProject, err)
		return nil
	}

	return DependentPathInfo{
		Path: path,
	}
}
