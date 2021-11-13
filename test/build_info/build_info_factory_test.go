package build_info

import (
	build_info2 "manifold/internal/build_info"
	"manifold/internal/document_definition"
	"manifold/test"
	"testing"
)

func TestBuildInfoFactory(t *testing.T) {
	t.Run("EmptyDocument", func(t *testing.T) {
		context := test.NewFakeTraverseContext()
		document := document_definition.DocumentDefinition{}

		buildInfo := build_info2.FromDocumentDefinition(&document, &context)

		test.Assert(t, buildInfo == nil)
		test.Assert(t, len(context.Errors) > 0)
		test.Assert(t, context.Errors[0] == "document is empty")
	})

	t.Run("ProjectDocument", func(t *testing.T) {
		t.Run("EmptyName", func(t *testing.T) {
			context := test.NewFakeTraverseContext()

			project := document_definition.ProjectDefinition{}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo := build_info2.FromDocumentDefinition(&document, &context)

			test.Assert(t, buildInfo == nil)
			test.Assert(t, len(context.Errors) > 0)
			test.Assert(t, len(context.Warnings) == 0)
			test.Assert(t, context.Errors[0] == "project does not contain name")
		})

		t.Run("WithNameWithoutSteps", func(t *testing.T) {
			context := test.NewFakeTraverseContext()
			context.File = "foo"

			project := document_definition.ProjectDefinition{
				Name: "foo",
			}
			document := document_definition.DocumentDefinition{
				Project: &project,
			}

			buildInfo := build_info2.FromDocumentDefinition(&document, &context)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == project.Name)
			test.Assert(t, buildInfo.Path() == context.File)
			test.Assert(t, buildInfo.Kind() == build_info2.ProjectBuildInfoKind)
			test.Assert(t, len(buildInfo.(build_info2.ProjectBuildInfo).Steps) == 0)

			test.Assert(t, len(context.Errors) == 0)
			test.Assert(t, len(context.Warnings) > 0)
			test.Assert(t, context.Warnings[0] == "project does not contain any steps")
		})

		t.Run("WithNameAndSteps", func(t *testing.T) {
			context := test.NewFakeTraverseContext()
			context.File = "foo"

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

			buildInfo := build_info2.FromDocumentDefinition(&document, &context)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == project.Name)
			test.Assert(t, buildInfo.Path() == context.File)
			test.Assert(t, buildInfo.Kind() == build_info2.ProjectBuildInfoKind)
			test.Assert(t, len(buildInfo.(build_info2.ProjectBuildInfo).Steps) > 0)

			test.Assert(t, len(context.Errors) == 0)
			test.Assert(t, len(context.Warnings) == 0)
		})
	})

	t.Run("WorkspaceDocument", func(t *testing.T) {
		t.Run("EmptyName", func(t *testing.T) {
			context := test.NewFakeTraverseContext()

			workspace := document_definition.WorkspaceDefinition{}
			document := document_definition.DocumentDefinition{
				Workspace: &workspace,
			}

			buildInfo := build_info2.FromDocumentDefinition(&document, &context)

			test.Assert(t, buildInfo == nil)
			test.Assert(t, len(context.Errors) > 0)
			test.Assert(t, len(context.Warnings) == 0)
			test.Assert(t, context.Errors[0] == "workspace does not contain name")
		})

		t.Run("WithName", func(t *testing.T) {
			context := test.NewFakeTraverseContext()
			context.File = "foo"

			workspace := document_definition.WorkspaceDefinition{
				Name: "foo",
			}
			document := document_definition.DocumentDefinition{
				Workspace: &workspace,
			}

			buildInfo := build_info2.FromDocumentDefinition(&document, &context)

			test.Assert(t, buildInfo != nil)
			test.Assert(t, buildInfo.Name() == workspace.Name)
			test.Assert(t, buildInfo.Path() == context.File)
			test.Assert(t, buildInfo.Kind() == build_info2.WorkspaceBuildInfoKind)

			test.Assert(t, len(context.Errors) == 0)
			test.Assert(t, len(context.Warnings) == 0)
		})
	})
}