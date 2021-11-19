package config

import (
	"errors"
	"manifold/internal/validation"
)

type Configuration interface {
	Validatable

	Target() Target
}

type configuration struct {
	ProjectTarget   *projectTarget   `yaml:"project"`
	WorkspaceTarget *workspaceTarget `yaml:"workspace"`
}

func (c *configuration) Validate(ctx ValidationContext) error {
	switch {
	case c.ProjectTarget != nil && c.WorkspaceTarget != nil:
		return errors.New(validation.ConfigurationWithProjectAndWorkspace)

	case c.ProjectTarget != nil:
		return c.ProjectTarget.Validate(ctx)

	case c.WorkspaceTarget != nil:
		return c.WorkspaceTarget.Validate(ctx)

	default:
		return errors.New(validation.EmptyConfiguration)
	}
}

func (c *configuration) Target() Target {
	switch {
	case c.ProjectTarget != nil:
		return c.ProjectTarget

	case c.WorkspaceTarget != nil:
		return c.WorkspaceTarget

	default:
		return nil
	}
}
