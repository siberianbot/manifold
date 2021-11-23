package config

import "testing"

func TestProjectDependency(t *testing.T) {
	t.Run("EmptyDependency", func(t *testing.T) {
		projectDependency := ProjectDependency{
			Path:    "",
			Project: "",
		}

		dependency := projectDependency.ToDependency()

		if dependency.Kind() != UnknownDependencyKind {
			t.Errorf("dependency.Kind() is %v, not UnknownDependencyKind", dependency.Kind())
		}

		if dependency.Value() != "" {
			t.Errorf("dependency.Value() is %v, not empty string", dependency.Value())
		}
	})

	t.Run("ProjectDependency", func(t *testing.T) {
		projectDependency := ProjectDependency{
			Path:    "",
			Project: "foo",
		}

		dependency := projectDependency.ToDependency()

		if dependency.Kind() != ProjectDependencyKind {
			t.Errorf("dependency.Kind() is %v, not ProjectDependencyKind", dependency.Kind())
		}

		if dependency.Value() != "foo" {
			t.Errorf("dependency.Value() is %v, not foo", dependency.Value())
		}
	})

	t.Run("PathDependency", func(t *testing.T) {
		projectDependency := ProjectDependency{
			Path:    "foo",
			Project: "",
		}

		dependency := projectDependency.ToDependency()

		if dependency.Kind() != PathDependencyKind {
			t.Errorf("dependency.Kind() is %v, not PathDependencyKind", dependency.Kind())
		}

		if dependency.Value() != "foo" {
			t.Errorf("dependency.Value() is %v, not foo", dependency.Value())
		}
	})

	t.Run("AmbiguousDependency", func(t *testing.T) {
		projectDependency := ProjectDependency{
			Path:    "foo",
			Project: "bar",
		}

		dependency := projectDependency.ToDependency()

		if dependency.Kind() != UnknownDependencyKind {
			t.Errorf("dependency.Kind() is %v, not UnknownDependencyKind", dependency.Kind())
		}

		if dependency.Value() != "" {
			t.Errorf("dependency.Value() is %v, not empty string", dependency.Value())
		}
	})
}
