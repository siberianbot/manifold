package traversing_test

import (
	"fmt"
	"manifold/internal/document_definition"
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/dependents"
	"manifold/internal/validation"
	"manifold/test"
	"os"
	"path"
	"testing"
)

func TestBuildInfoFactory(t *testing.T) {
	t.Run("EmptyDocument", func(t *testing.T) {
		expected := validation.EmptyDocument
		ctx := test.NewFakeContext()

		document := document_definition.DocumentDefinition{}

		buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

		test.Assert(t, buildInfo == nil)
		test.Assert(t, dependencies != nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("BothProjectAndWorkspace", func(t *testing.T) {
		expected := validation.DocumentWithBothProjectAndWorkspace
		ctx := test.NewFakeContext()

		project := document_definition.ProjectDefinition{}
		workspace := document_definition.WorkspaceDefinition{}
		document := document_definition.DocumentDefinition{
			Project:   &project,
			Workspace: &workspace,
		}

		buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

		test.Assert(t, buildInfo == nil)
		test.Assert(t, dependencies != nil)
		test.Assert(t, len(ctx.Errors) > 0)
		test.Assert(t, ctx.Errors[0] == expected)
	})

	t.Run("ProjectDocument", func(t *testing.T) {
		t.Run("InvalidName", func(t *testing.T) {
			projectName := "!!!"
			expected := fmt.Sprintf(validation.InvalidProject, fmt.Sprintf(validation.InvalidManifoldName, projectName, validation.NameRegexPattern))
			ctx := test.NewFakeContext()

			project := document_definition.ProjectDefinition{
				Name: projectName,
			}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo == nil)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 1)
			test.Assert(t, len(ctx.Warnings) == 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidNameOnly", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			project := document_definition.ProjectDefinition{
				Name: "foo",
			}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == project.Name)
			test.Assert(t, buildInfo.Path() == ctx.File)
			test.Assert(t, buildInfo.Kind() == build_info.ProjectBuildInfoKind)
			test.Assert(t, len(buildInfo.(build_info.ProjectBuildInfo).Steps) == 0)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 1)
			test.Assert(t, ctx.Warnings[0] == validation.ProjectWithoutSteps)
		})

		t.Run("WithSteps", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			step := document_definition.StepDefinition{
				Command: "foo",
			}
			project := document_definition.ProjectDefinition{
				Name:  "foo",
				Steps: []document_definition.StepDefinition{step},
			}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == project.Name)
			test.Assert(t, buildInfo.Path() == ctx.File)
			test.Assert(t, buildInfo.Kind() == build_info.ProjectBuildInfoKind)
			test.Assert(t, len(buildInfo.(build_info.ProjectBuildInfo).Steps) > 0)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})

		t.Run("WithDependencies", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			step := document_definition.StepDefinition{
				Command: "foo",
			}
			dependency := document_definition.DependencyDefinition{
				Project: "bar",
			}
			project := document_definition.ProjectDefinition{
				Name:         "foo",
				Dependencies: []document_definition.DependencyDefinition{dependency},
				Steps:        []document_definition.StepDefinition{step},
			}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == project.Name)
			test.Assert(t, buildInfo.Path() == ctx.File)
			test.Assert(t, buildInfo.Kind() == build_info.ProjectBuildInfoKind)
			test.Assert(t, len(buildInfo.(build_info.ProjectBuildInfo).Steps) > 0)

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

			workspace := document_definition.WorkspaceDefinition{
				Name: workspaceName,
			}
			document := document_definition.DocumentDefinition{
				Workspace: &workspace,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo == nil)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 1)
			test.Assert(t, len(ctx.Warnings) == 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidNameOnly", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			workspace := document_definition.WorkspaceDefinition{
				Name: "foo",
			}
			document := document_definition.DocumentDefinition{
				Workspace: &workspace,
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == workspace.Name)
			test.Assert(t, buildInfo.Path() == ctx.File)
			test.Assert(t, buildInfo.Kind() == build_info.WorkspaceBuildInfoKind)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 0)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})

		t.Run("WithIncludes", func(t *testing.T) {
			ctx := test.NewFakeContext()
			ctx.File = "foo"

			include := path.Join(t.TempDir(), "bar")
			workspace := document_definition.WorkspaceDefinition{
				Name:     "foo",
				Includes: []document_definition.IncludeDefinition{document_definition.IncludeDefinition(include)},
			}
			document := document_definition.DocumentDefinition{
				Workspace: &workspace,
			}

			file, err := os.Create(include)

			if err != nil {
				t.Fatal(err)
				return
			} else {
				_ = file.Close()

				defer func() { _ = os.Remove(include) }()
			}

			buildInfo, dependencies := build_info.FromDocumentDefinition(&document, &ctx)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == workspace.Name)
			test.Assert(t, buildInfo.Path() == ctx.File)
			test.Assert(t, buildInfo.Kind() == build_info.WorkspaceBuildInfoKind)

			test.Assert(t, dependencies != nil)
			test.Assert(t, len(dependencies) == 1)
			test.Assert(t, dependencies[0].Kind() == dependents.DependentPathInfoKind)
			test.Assert(t, dependencies[0].(dependents.DependentPathInfo).Path == include)

			test.Assert(t, len(ctx.Errors) == 0)
			test.Assert(t, len(ctx.Warnings) == 0)
		})
	})
}
