package validation

import (
	"manifold/internal/config"
	"manifold/internal/utils"
	"manifold/internal/validation"
	"path/filepath"
)

func Validate(cfg *config.Configuration, path string) error {
	dir := filepath.Dir(path)

	return validateConfiguration(cfg, dir)
}

func validateConfiguration(cfg *config.Configuration, dir string) error {
	switch {
	case cfg.ProjectTarget != nil && cfg.WorkspaceTarget != nil:
		return validation.NewError(AmbiguousConfiguration)

	case cfg.ProjectTarget != nil:
		return validateProject(cfg.ProjectTarget, dir)

	case cfg.WorkspaceTarget != nil:
		return validateWorkspace(cfg.WorkspaceTarget, dir)

	default:
		return validation.NewError(EmptyConfiguration)
	}
}

func validateProject(project *config.ProjectTarget, dir string) error {
	if err := validation.ValidateManifoldName(project.Name); err != nil {
		return validation.NewError(InvalidProject, err)
	}

	for _, dependency := range project.ProjectDependencies {
		if err := validateProjectDependency(dependency, dir); err != nil {
			return err
		}
	}

	return nil
}

func validateWorkspace(workspace *config.WorkspaceTarget, dir string) error {
	if err := validation.ValidateManifoldName(workspace.Name); err != nil {
		return validation.NewError(InvalidWorkspace, err)
	}

	for _, include := range workspace.Includes {
		if err := validateInclude(include, dir); err != nil {
			return err
		}
	}

	return nil
}

func validateProjectDependency(dependency config.ProjectDependency, dir string) error {
	switch {
	case dependency.Project != "" && dependency.Path != "":
		return validation.NewError(AmbiguousProjectDependency)

	case dependency.Project != "":
		if err := validation.ValidateManifoldName(dependency.Project); err != nil {
			return validation.NewError(InvalidProjectDependency, err)
		}

		return nil

	case dependency.Path != "":
		path := utils.BuildPath(dir, dependency.Path)

		if err := validation.ValidatePath(path); err != nil {
			return validation.NewError(InvalidProjectDependency, err)
		}

		return nil

	default:
		return validation.NewError(EmptyProjectDependency)
	}
}

func validateInclude(include string, dir string) error {
	if include == "" {
		return validation.NewError(EmptyWorkspaceInclude)
	}

	path := utils.BuildPath(dir, include)

	if err := validation.ValidatePath(path); err != nil {
		return validation.NewError(InvalidWorkspaceInclude, err)
	}

	return nil
}
