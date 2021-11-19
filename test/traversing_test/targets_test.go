package traversing_test

import (
	"fmt"
	"manifold/internal/config"
	"manifold/internal/traversing/dependents"
	"manifold/internal/traversing/targets"
	"manifold/internal/validation"
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestTargetFactory(t *testing.T) {
	t.Run("EmptyDocument", func(t *testing.T) {
		expected := validation.EmptyDocument
		ctx := test.NewFakeContext()

		document := config.ConfigurationDefinition{}

		target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

		test.Assert(t, target == nil)
		test.Assert(t, dependencies != nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("BothProjectAndWorkspace", func(t *testing.T) {
		expected := validation.DocumentWithBothProjectAndWorkspace
		ctx := test.NewFakeContext()

		project := config.ProjectDefinition{}
		workspace := config.WorkspaceDefinition{}
		document := config.ConfigurationDefinition{
			Project:   &project,
			Workspace: &workspace,
		}

		target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

		test.Assert(t, target == nil)
		test.Assert(t, dependencies != nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("ProjectDocument", func(t *testing.T) {
		t.Run("InvalidName", func(t *testing.T) {
			projectName := "!!!"
			expected := fmt.Sprintf(validation.InvalidProject, fmt.Sprintf(validation.InvalidManifoldName, projectName, validation.NameRegexPattern))
			ctx := test.NewFakeContext()

			project := config.ProjectDefinition{
				Name: projectName,
			}
			document := config.ConfigurationDefinition{
				Project: &project,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target == nil)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 1)
			test.Assert(t, len(ctx.Warnings) == 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidNameOnly", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			project := config.ProjectDefinition{
				Name: "foo",
			}
			document := config.ConfigurationDefinition{
				Project: &project,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target != nil)
			test.Assert(t, target.Name() == project.Name)
			test.Assert(t, target.Path() == ctx.File)
			test.Assert(t, target.Kind() == targets.ProjectTargetKind)
			test.Assert(t, len(target.(targets.ProjectTarget).Steps) == 0)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 1)
			test.Assert(t, ctx.Warnings[0] == validation.ProjectWithoutSteps)
		})

		t.Run("WithSteps", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			step := config.StepDefinition{
				"cmd": "foo",
			}
			project := config.ProjectDefinition{
				Name:  "foo",
				Steps: []config.StepDefinition{step},
			}
			document := config.ConfigurationDefinition{
				Project: &project,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target != nil)
			test.Assert(t, target.Name() == project.Name)
			test.Assert(t, target.Path() == ctx.File)
			test.Assert(t, target.Kind() == targets.ProjectTargetKind)
			test.Assert(t, len(target.(targets.ProjectTarget).Steps) > 0)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})

		t.Run("WithDependencies", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			step := config.StepDefinition{
				"cmd": "foo",
			}
			dependency := config.DependencyDefinition{
				Project: "bar",
			}
			project := config.ProjectDefinition{
				Name:         "foo",
				Dependencies: []config.DependencyDefinition{dependency},
				Steps:        []config.StepDefinition{step},
			}
			document := config.ConfigurationDefinition{
				Project: &project,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target != nil)
			test.Assert(t, target.Name() == project.Name)
			test.Assert(t, target.Path() == ctx.File)
			test.Assert(t, target.Kind() == targets.ProjectTargetKind)
			test.Assert(t, len(target.(targets.ProjectTarget).Steps) > 0)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 1)
			test.Assert(t, dependencies[0].Kind() == dependents.DependentProjectInfoKind)
			test.Assert(t, dependencies[0].(dependents.DependentProjectInfo).Project == dependency.Project)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})
	})

	t.Run("WorkspaceDocument", func(t *testing.T) {
		t.Run("InvalidName", func(t *testing.T) {
			workspaceName := "!!!"
			expected := fmt.Sprintf(validation.InvalidWorkspace, fmt.Sprintf(validation.InvalidManifoldName, workspaceName, validation.NameRegexPattern))
			ctx := test.NewFakeContext()

			workspace := config.WorkspaceDefinition{
				Name: workspaceName,
			}
			document := config.ConfigurationDefinition{
				Workspace: &workspace,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target == nil)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 1)
			test.Assert(t, len(ctx.Warnings) == 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidNameOnly", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			workspace := config.WorkspaceDefinition{
				Name: "foo",
			}
			document := config.ConfigurationDefinition{
				Workspace: &workspace,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, target != nil)
			test.Assert(t, target.Name() == workspace.Name)
			test.Assert(t, target.Path() == ctx.File)
			test.Assert(t, target.Kind() == targets.WorkspaceTargetKind)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})

		t.Run("WithIncludes", func(t *testing.T) {
			testDir := t.TempDir()

			ctx := test.NewFakeContext()
			ctx.File = filepath.Join(testDir, "foo", ".manifold.yml")
			test.CreateFile(t, ctx.File, "")

			include := "bar"
			includePath := filepath.Join(testDir, "foo", include)
			test.CreateDir(t, includePath)

			workspace := config.WorkspaceDefinition{
				Name:     "foo",
				Includes: []config.IncludeDefinition{config.IncludeDefinition(include)},
			}
			document := config.ConfigurationDefinition{
				Workspace: &workspace,
			}

			target, dependencies := targets.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)

			test.Assert(t, target != nil)
			test.Assert(t, target.Name() == workspace.Name)
			test.Assert(t, target.Path() == ctx.File)
			test.Assert(t, target.Kind() == targets.WorkspaceTargetKind)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 1)
			test.Assert(t, dependencies[0].Kind() == dependents.DependentPathInfoKind)
			test.Assert(t, dependencies[0].(dependents.DependentPathInfo).Path == includePath)
		})
	})
}
