package validation

import (
	"manifold/internal/config"
	"manifold/internal/utils"
)

func ValidateConfiguration(cfg *config.Configuration, ctx *Context) error {
	switch {
	case cfg.ProjectTarget != nil && cfg.WorkspaceTarget != nil:
		return NewError(ConfigurationWithProjectAndWorkspace)

	case cfg.ProjectTarget != nil:
		return validateProject(cfg.ProjectTarget, ctx)

	case cfg.WorkspaceTarget != nil:
		return validateWorkspace(cfg.WorkspaceTarget, ctx)

	default:
		return NewError(EmptyConfiguration)
	}
}

func validateProject(project *config.ProjectTarget, ctx *Context) error {
	if err := ValidateManifoldName(project.Name); err != nil {
		return NewError(InvalidProject, err)
	}

	for _, dependency := range project.ProjectDependencies {
		if err := validateProjectDependency(dependency, ctx); err != nil {
			return NewError(InvalidProjectDependency, err)
		}
	}

	return nil
}

func validateWorkspace(workspace *config.WorkspaceTarget, ctx *Context) error {
	if err := ValidateManifoldName(workspace.Name); err != nil {
		return NewError(InvalidWorkspace, err)
	}

	for _, include := range workspace.Includes {
		if err := validateInclude(include, ctx); err != nil {
			return NewError(InvalidWorkspaceInclude, err)
		}
	}

	return nil
}

func validateProjectDependency(dependency config.ProjectDependency, ctx *Context) error {
	switch {
	case dependency.Project != "" && dependency.Path != "":
		return NewError(ProjectDependencyWithBothProjectAndPath)

	case dependency.Project != "":
		return ValidateManifoldName(dependency.Project)

	case dependency.Path != "":
		return ValidatePath(utils.BuildPath(ctx.Dir(), dependency.Path))

	default:
		return NewError(EmptyProjectDependency)
	}
}

func validateInclude(include string, ctx *Context) error {
	if include == "" {
		return NewError(EmptyWorkspaceInclude)
	}

	return ValidatePath(utils.BuildPath(ctx.Dir(), include))
}
