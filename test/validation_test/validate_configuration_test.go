package validation

import (
	"fmt"
	"manifold/internal/config"
	"manifold/internal/validation"
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestValidateConfiguration(t *testing.T) {
	t.Run("EmptyConfig", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		cfg := config.Configuration{}

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == validation.EmptyConfiguration)
	})

	t.Run("BothProjectAndWorkspaceInConfig", func(t *testing.T) {
		ctx := validation.NewContext("foo")

		cfg := config.Configuration{
			ProjectTarget:   &config.ProjectTarget{},
			WorkspaceTarget: &config.WorkspaceTarget{},
		}

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == validation.ConfigurationWithProjectAndWorkspace)
	})

	t.Run("ProjectConfig", testProjectValidation)
	t.Run("WorkspaceConfig", testWorkspaceValidation)
}

func testProjectValidation(t *testing.T) {
	t.Run("EmptyProject", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{},
		}
		expected := fmt.Sprintf(validation.InvalidProject, validation.EmptyManifoldName)

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("InvalidName", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo!"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
			},
		}
		expected := fmt.Sprintf(validation.InvalidProject, fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern))

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ValidName", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
			},
		}

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err == nil)
	})

	t.Run("EmptyProjectDependency", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
				ProjectDependencies: []config.ProjectDependency{
					{},
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidProjectDependency, validation.EmptyProjectDependency)

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ProjectDependencyWithProjectAndPath", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
				ProjectDependencies: []config.ProjectDependency{
					{Project: "bar", Path: "bar"},
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidProjectDependency, validation.ProjectDependencyWithBothProjectAndPath)

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ProjectDependencyWithInvalidProject", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "bar!"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Project: name},
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidProjectDependency, fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern))

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ProjectDependencyWithInvalidPath", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		path := "bar"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Path: path},
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidProjectDependency, fmt.Sprintf(validation.InvalidPath, filepath.Join(ctx.Dir(), path)))

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})
}

func testWorkspaceValidation(t *testing.T) {
	t.Run("EmptyWorkspace", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{},
		}
		expected := fmt.Sprintf(validation.InvalidWorkspace, validation.EmptyManifoldName)

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("InvalidName", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo!"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: name,
			},
		}
		expected := fmt.Sprintf(validation.InvalidWorkspace, fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern))

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ValidName", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: name,
			},
		}

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err == nil)
	})

	t.Run("EmptyInclude", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: "foo",
				Includes: []string{
					"",
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidWorkspaceInclude, validation.EmptyWorkspaceInclude)

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("InvalidInclude", func(t *testing.T) {
		ctx := validation.NewContext("foo")
		include := "bar"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: "foo",
				Includes: []string{
					include,
				},
			},
		}
		expected := fmt.Sprintf(validation.InvalidWorkspaceInclude, fmt.Sprintf(validation.InvalidPath, include))

		err := validation.ValidateConfiguration(&cfg, ctx)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})
}
