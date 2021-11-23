package validation

import (
	"fmt"
	"manifold/internal/config"
	"path/filepath"
	"testing"
)

func TestValidateConfiguration(t *testing.T) {
	t.Run("EmptyConfig", func(t *testing.T) {
		ctx := NewContext("foo")
		cfg := config.Configuration{}

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != EmptyConfiguration {
			t.Errorf("error is %v, not %s", err, EmptyConfiguration)
		}
	})

	t.Run("BothProjectAndWorkspaceInConfig", func(t *testing.T) {
		ctx := NewContext("foo")

		cfg := config.Configuration{
			ProjectTarget:   &config.ProjectTarget{},
			WorkspaceTarget: &config.WorkspaceTarget{},
		}

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != ConfigurationWithProjectAndWorkspace {
			t.Errorf("error is %v, not %s", err, ConfigurationWithProjectAndWorkspace)
		}
	})

	t.Run("ProjectConfig", testProjectValidation)
	t.Run("WorkspaceConfig", testWorkspaceValidation)
}

func testProjectValidation(t *testing.T) {
	t.Run("EmptyProject", func(t *testing.T) {
		ctx := NewContext("foo")
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{},
		}
		expected := fmt.Sprintf(InvalidProject, EmptyManifoldName)

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("InvalidName", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo!"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
			},
		}
		expected := fmt.Sprintf(InvalidProject, fmt.Sprintf(InvalidManifoldName, name, NameRegexPattern))

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ValidName", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
			},
		}

		err := ValidateConfiguration(&cfg, ctx)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}
	})

	t.Run("EmptyProjectDependency", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
				ProjectDependencies: []config.ProjectDependency{
					{},
				},
			},
		}
		expected := fmt.Sprintf(InvalidProjectDependency, EmptyProjectDependency)

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ProjectDependencyWithProjectAndPath", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: name,
				ProjectDependencies: []config.ProjectDependency{
					{Project: "bar", Path: "bar"},
				},
			},
		}
		expected := fmt.Sprintf(InvalidProjectDependency, ProjectDependencyWithBothProjectAndPath)

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ProjectDependencyWithInvalidProject", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "bar!"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Project: name},
				},
			},
		}
		expected := fmt.Sprintf(InvalidProjectDependency, fmt.Sprintf(InvalidManifoldName, name, NameRegexPattern))

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ProjectDependencyWithInvalidPath", func(t *testing.T) {
		ctx := NewContext("foo")
		path := "bar"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Path: path},
				},
			},
		}
		expected := fmt.Sprintf(InvalidProjectDependency, fmt.Sprintf(InvalidPath, filepath.Join(ctx.Dir(), path)))

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})
}

func testWorkspaceValidation(t *testing.T) {
	t.Run("EmptyWorkspace", func(t *testing.T) {
		ctx := NewContext("foo")
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{},
		}
		expected := fmt.Sprintf(InvalidWorkspace, EmptyManifoldName)

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("InvalidName", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo!"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: name,
			},
		}
		expected := fmt.Sprintf(InvalidWorkspace, fmt.Sprintf(InvalidManifoldName, name, NameRegexPattern))

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ValidName", func(t *testing.T) {
		ctx := NewContext("foo")
		name := "foo"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: name,
			},
		}

		err := ValidateConfiguration(&cfg, ctx)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}
	})

	t.Run("EmptyInclude", func(t *testing.T) {
		ctx := NewContext("foo")
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: "foo",
				Includes: []string{
					"",
				},
			},
		}
		expected := fmt.Sprintf(InvalidWorkspaceInclude, EmptyWorkspaceInclude)

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("InvalidInclude", func(t *testing.T) {
		ctx := NewContext("foo")
		include := "bar"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: "foo",
				Includes: []string{
					include,
				},
			},
		}
		expected := fmt.Sprintf(InvalidWorkspaceInclude, fmt.Sprintf(InvalidPath, include))

		err := ValidateConfiguration(&cfg, ctx)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})
}
