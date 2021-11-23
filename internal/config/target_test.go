package config

import "testing"

func TestProjectTarget(t *testing.T) {
	t.Run("NoDependencies", func(t *testing.T) {
		project := ProjectTarget{ProjectDependencies: []ProjectDependency{}}

		dependencies := project.Dependencies()

		if dependencies == nil {
			t.Error("dependencies is nil")
		} else if len(dependencies) != 0 {
			t.Error("dependencies is not empty")
		}
	})

	t.Run("WithDependencies", func(t *testing.T) {
		project := ProjectTarget{
			ProjectDependencies: []ProjectDependency{
				{},
				{Project: "foo"},
				{Path: "bar"},
				{Project: "baz", Path: "baz"},
			}}

		dependencies := project.Dependencies()

		if dependencies == nil {
			t.Error("dependencies is nil")
		} else if len(dependencies) != len(project.ProjectDependencies) {
			t.Errorf("len(dependencies) is %v, not %v", len(dependencies), len(project.ProjectDependencies))
		} else {
			if !containsDependency(dependencies, UnknownDependencyKind, "") {
				t.Error("dependencies doesn't contain empty or ambiguous dependency")
			}

			if !containsDependency(dependencies, PathDependencyKind, "bar") {
				t.Error("dependencies doesn't contain path bar")
			}

			if !containsDependency(dependencies, ProjectDependencyKind, "foo") {
				t.Error("dependencies doesn't contain project foo")
			}
		}
	})
}

func TestWorkspaceTarget(t *testing.T) {
	t.Run("NoDependencies", func(t *testing.T) {
		workspace := WorkspaceTarget{Includes: []string{}}

		dependencies := workspace.Dependencies()

		if dependencies == nil {
			t.Error("dependencies is nil")
		} else if len(dependencies) != 0 {
			t.Error("dependencies is not empty")
		}
	})

	t.Run("WithDependencies", func(t *testing.T) {
		workspace := WorkspaceTarget{
			Includes: []string{
				"",
				"foo",
				"bar",
				"baz",
			}}

		dependencies := workspace.Dependencies()

		if dependencies == nil {
			t.Error("dependencies is nil")
		} else if len(dependencies) != len(workspace.Includes) {
			t.Errorf("len(dependencies) is %v, not %v", len(dependencies), len(workspace.Includes))
		} else {
			if !containsDependency(dependencies, PathDependencyKind, "") {
				t.Error("dependencies doesn't contain empty dependency")
			}

			if !containsDependency(dependencies, PathDependencyKind, "foo") {
				t.Error("dependencies doesn't contain path foo")
			}

			if !containsDependency(dependencies, PathDependencyKind, "bar") {
				t.Error("dependencies doesn't contain path bar")
			}

			if !containsDependency(dependencies, PathDependencyKind, "baz") {
				t.Error("dependencies doesn't contain path baz")
			}
		}
	})
}

func containsDependency(dependencies []Dependency, kind DependencyKind, value string) bool {
	for _, dependency := range dependencies {
		if dependency.Kind() == kind && dependency.Value() == value {
			return true
		}
	}

	return false
}
