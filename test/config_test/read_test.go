package config_test

import (
	"bytes"
	"manifold/internal/config"
	"manifold/test"
	"math/rand"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("EmptyFile", func(t *testing.T) {
		reader := strings.NewReader("")
		configuration, err := config.Read(reader)

		test.Assert(t, configuration == nil)
		test.Assert(t, err != nil)
	})

	t.Run("RandomFile", func(t *testing.T) {
		data := make([]byte, 1024)
		rand.Read(data)

		reader := bytes.NewReader(data)
		configuration, err := config.Read(reader)

		test.Assert(t, configuration == nil)
		test.Assert(t, err != nil)
	})

	t.Run("Project", testReadProject)
	t.Run("Workspace", testReadWorkspace)
}

func testReadProject(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		yaml := `project:`

		reader := strings.NewReader(yaml)
		configuration, err := config.Read(reader)

		test.Assert(t, err == nil)
		test.Assert(t, configuration != nil)
		test.Assert(t, configuration.Target() == nil)
	})

	t.Run("Full", func(t *testing.T) {
		yaml := `
project:
  name: foo
  dependencies:
  - project: bar
  - path: baz
  steps:
  - foo: foo
  - bar: bar
  - baz: baz
`

		reader := strings.NewReader(yaml)
		configuration, err := config.Read(reader)

		test.Assert(t, err == nil)
		test.Assert(t, configuration != nil)

		target, targetOk := configuration.Target().(config.ProjectTarget)
		test.Assert(t, target != nil)
		test.Assert(t, targetOk)
		test.Assert(t, target.Name() == "foo")
		test.Assert(t, target.Kind() == config.ProjectTargetKind)

		steps := target.Steps()
		test.Assert(t, steps != nil)
		test.Assert(t, len(steps) == 3)
		test.Assert(t, containsNamedStep(steps, "foo"))
		test.Assert(t, containsNamedStep(steps, "bar"))
		test.Assert(t, containsNamedStep(steps, "baz"))

		dependencies := target.Dependencies()
		test.Assert(t, dependencies != nil)
		test.Assert(t, len(dependencies) == 2)
		test.Assert(t, containsDependency(dependencies, config.ProjectDependencyKind, "bar"))
		test.Assert(t, containsDependency(dependencies, config.PathDependencyKind, "baz"))
	})
}

func testReadWorkspace(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		yaml := `workspace:`

		reader := strings.NewReader(yaml)
		configuration, err := config.Read(reader)

		test.Assert(t, err == nil)
		test.Assert(t, configuration != nil)
		test.Assert(t, configuration.Target() == nil)
	})

	t.Run("Full", func(t *testing.T) {
		yaml := `
workspace:
  name: foo
  includes:
  - bar
  - baz
`

		reader := strings.NewReader(yaml)
		configuration, err := config.Read(reader)

		test.Assert(t, err == nil)
		test.Assert(t, configuration != nil)

		target, targetOk := configuration.Target().(config.WorkspaceTarget)
		test.Assert(t, target != nil)
		test.Assert(t, targetOk)
		test.Assert(t, target.Name() == "foo")
		test.Assert(t, target.Kind() == config.WorkspaceTargetKind)

		includes := target.Includes()
		test.Assert(t, includes != nil)
		test.Assert(t, len(includes) == 2)
		test.Assert(t, containsInclude(includes, "bar"))
		test.Assert(t, containsInclude(includes, "baz"))
	})
}

func containsNamedStep(steps []config.Step, name string) bool {
	for _, step := range steps {
		if step[name] == name {
			return true
		}
	}

	return false
}

func containsDependency(dependencies []config.Dependency, kind config.DependencyKind, value string) bool {
	for _, dependency := range dependencies {
		if dependency.Kind() == kind && dependency.Value() == value {
			return true
		}
	}

	return false
}

func containsInclude(includes []config.Include, path string) bool {
	for _, include := range includes {
		if include.Path() == path {
			return true
		}
	}

	return false
}
