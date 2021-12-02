package validation

import (
	"manifold/internal/config"
	"manifold/internal/errors"
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
	case cfg.Project != nil && cfg.Workspace != nil:
		return errors.NewError(AmbiguousConfiguration)

	case cfg.Project != nil:
		return validateProject(cfg.Project, dir)

	case cfg.Workspace != nil:
		return validateWorkspace(cfg.Workspace, dir)

	default:
		return errors.NewError(EmptyConfiguration)
	}
}

func validateProject(project *config.Project, dir string) error {
	if err := validation.ValidateManifoldName(project.Name); err != nil {
		return errors.NewError(InvalidProject, err)
	}

	for _, dependency := range project.Dependencies {
		if err := validateProjectDependency(dependency, dir); err != nil {
			return err
		}
	}

	return nil
}

func validateWorkspace(workspace *config.Workspace, dir string) error {
	if err := validation.ValidateManifoldName(workspace.Name); err != nil {
		return errors.NewError(InvalidWorkspace, err)
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
		return errors.NewError(AmbiguousProjectDependency)

	case dependency.Project != "":
		if err := validation.ValidateManifoldName(dependency.Project); err != nil {
			return errors.NewError(InvalidProjectDependency, err)
		}

		return nil

	case dependency.Path != "":
		path := utils.BuildPath(dir, dependency.Path)

		if err := validation.ValidatePath(path); err != nil {
			return errors.NewError(InvalidProjectDependency, err)
		}

		return nil

	default:
		return errors.NewError(EmptyProjectDependency)
	}
}

func validateInclude(include config.WorkspaceInclude, dir string) error {
	if include == "" {
		return errors.NewError(EmptyWorkspaceInclude)
	}

	path := utils.BuildPath(dir, string(include))

	if err := validation.ValidatePath(path); err != nil {
		return errors.NewError(InvalidWorkspaceInclude, err)
	}

	return nil
}
