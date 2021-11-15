package traversing_test_test

import (
	"fmt"
	"manifold/internal/document_definition"
	"manifold/internal/traversing/dependents"
	"manifold/internal/validation"
	"manifold/test"
	"os"
	"testing"
)

func TestDependentInfoFactory(t *testing.T) {
	t.Run("FromIncludeDefinition", func(t *testing.T) {
		t.Run("EmptyDefinition", func(t *testing.T) {
			expected := fmt.Sprintf(validation.InvalidDependentProject, validation.EmptyPath)

			ctx := test.NewFakeContext()
			includeDefinition := document_definition.IncludeDefinition("")

			dependentInfo := dependents.FromIncludeDefinition(includeDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("InvalidDefinition", func(t *testing.T) {
			path := "baz"
			expected := fmt.Sprintf(validation.InvalidDependentProject, fmt.Sprintf(validation.InvalidPath, path))

			ctx := test.NewFakeContext()
			includeDefinition := document_definition.IncludeDefinition(path)

			dependentInfo := dependents.FromIncludeDefinition(includeDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidDefinition", func(t *testing.T) {
			path := "foo"

			file, err := os.Create(path)

			if err != nil {
				t.Fatal(err)
				return
			} else {
				_ = file.Close()

				defer func() { _ = os.Remove(path) }()
			}

			ctx := test.NewFakeContext()
			includeDefinition := document_definition.IncludeDefinition(path)

			dependentInfo := dependents.FromIncludeDefinition(includeDefinition, &ctx)

			test.Assert(t, dependentInfo != nil)
			test.Assert(t, dependentInfo.Kind() == dependents.DependentPathInfoKind)
			test.Assert(t, dependentInfo.(dependents.DependentPathInfo).Path == path)
			test.Assert(t, len(ctx.Errors) == 0)
		})
	})

	t.Run("FromDependencyDefinition", func(t *testing.T) {
		t.Run("EmptyDefinition", func(t *testing.T) {
			expected := validation.EmptyProjectDependency

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("BothProjectAndPath", func(t *testing.T) {
			expected := validation.DependencyWithBothProjectAndPath

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{
				Project: "foo",
				Path:    "bar",
			}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("InvalidProject", func(t *testing.T) {
			project := "foo!"
			expected := fmt.Sprintf(validation.InvalidDependentProject, fmt.Sprintf(validation.InvalidManifoldName, project, validation.NameRegexPattern))

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{
				Project: project,
			}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidProject", func(t *testing.T) {
			project := "foo"

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{
				Project: project,
			}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo != nil)
			test.Assert(t, dependentInfo.Kind() == dependents.DependentProjectInfoKind)
			test.Assert(t, dependentInfo.(dependents.DependentProjectInfo).Project == project)
			test.Assert(t, len(ctx.Errors) == 0)
		})

		t.Run("InvalidPath", func(t *testing.T) {
			path := "baz"
			expected := fmt.Sprintf(validation.InvalidDependentProject, fmt.Sprintf(validation.InvalidPath, path))

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{
				Path: path,
			}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo == nil)
			test.Assert(t, len(ctx.Errors) > 0)
			test.Assert(t, ctx.Errors[0] == expected)
		})

		t.Run("ValidPath", func(t *testing.T) {
			path := "foo"

			file, err := os.Create(path)

			if err != nil {
				t.Fatal(err)
				return
			} else {
				_ = file.Close()

				defer func() { _ = os.Remove(path) }()
			}

			ctx := test.NewFakeContext()
			dependencyDefinition := document_definition.DependencyDefinition{
				Path: path,
			}

			dependentInfo := dependents.FromDependencyDefinition(dependencyDefinition, &ctx)

			test.Assert(t, dependentInfo != nil)
			test.Assert(t, dependentInfo.Kind() == dependents.DependentPathInfoKind)
			test.Assert(t, dependentInfo.(dependents.DependentPathInfo).Path == path)
			test.Assert(t, len(ctx.Errors) == 0)
		})
	})
}
