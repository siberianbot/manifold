package dependents

import (
	"manifold/internal/config"
	"manifold/internal/traversing"
	"manifold/internal/validation"
)

func FromDependencyDefinition(definition config.DependencyDefinition, ctx traversing.Context) DependentInfo {
	switch {
	case definition.Project != "" && definition.Path != "":
		ctx.AddError(validation.DependencyWithBothProjectAndPath)
		return nil

	case definition.Project != "":
		return newDependentProject(definition.Project, ctx)

	case definition.Path != "":
		return newDependentPath(definition.Path, ctx)

	default:
		ctx.AddError(validation.EmptyProjectDependency)
		return nil
	}
}

func FromIncludeDefinition(definition config.IncludeDefinition, ctx traversing.Context) DependentInfo {
	if definition == "" {
		ctx.AddError(validation.EmptyWorkspaceInclude)
		return nil
	}

	pathStr := string(definition)

	return newDependentPath(pathStr, ctx)
}
