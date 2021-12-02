package config

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("EmptyFile", func(t *testing.T) {
		reader := strings.NewReader("")
		configuration, err := Read(reader)

		assert.Empty(t, configuration)
		assert.Error(t, err)
	})

	t.Run("RandomFile", func(t *testing.T) {
		data := make([]byte, 1024)
		rand.Read(data)

		reader := bytes.NewReader(data)
		configuration, err := Read(reader)

		assert.Empty(t, configuration)
		assert.Error(t, err)
	})

	t.Run("ProjectConfig", func(t *testing.T) {
		content := `
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

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		assert.NotEmpty(t, configuration)
		assert.NoError(t, err)

		assert.NotEmpty(t, configuration.Project)

		assert.Equal(t, "foo", configuration.Project.Name)

		assert.NotEmpty(t, configuration.Project.Steps)
		assert.Len(t, configuration.Project.Steps, 3)
		assert.True(t, containsNamedStep(configuration.Project.Steps, "foo"), "configuration.Project.Steps doesn't contains foo")
		assert.True(t, containsNamedStep(configuration.Project.Steps, "bar"), "configuration.Project.Steps doesn't contains bar")
		assert.True(t, containsNamedStep(configuration.Project.Steps, "baz"), "configuration.Project.Steps doesn't contains baz")

		assert.NotEmpty(t, configuration.Project.Dependencies)
		assert.Len(t, configuration.Project.Dependencies, 2)
		assert.True(t, containsProjectDependency(configuration.Project.Dependencies, "", "bar"), "configuration.Project.Dependencies doesn't contains project bar")
		assert.True(t, containsProjectDependency(configuration.Project.Dependencies, "baz", ""), "configuration.Project.Dependencies doesn't contains path baz")
	})

	t.Run("WorkspaceConfig", func(t *testing.T) {
		content := `
workspace:
  name: foo
  includes:
  - bar
  - baz
`

		reader := strings.NewReader(content)
		configuration, err := Read(reader)

		assert.NotEmpty(t, configuration)
		assert.NoError(t, err)

		assert.NotEmpty(t, configuration.Workspace)

		assert.Equal(t, "foo", configuration.Workspace.Name)

		assert.NotEmpty(t, configuration.Workspace.Includes)
		assert.Len(t, configuration.Workspace.Includes, 2)
		assert.True(t, containsInclude(configuration.Workspace.Includes, "bar"), "configuration.Workspace.Includes doesn't contains bar")
		assert.True(t, containsInclude(configuration.Workspace.Includes, "baz"), "configuration.Workspace.Includes doesn't contains baz")
	})
}

func containsNamedStep(steps []ProjectStep, name string) bool {
	for _, step := range steps {
		if step[name] == name {
			return true
		}
	}

	return false
}

func containsInclude(includes []WorkspaceInclude, include WorkspaceInclude) bool {
	for _, inc := range includes {
		if inc == include {
			return true
		}
	}

	return false
}

func containsProjectDependency(projectDependencies []ProjectDependency, path string, project string) bool {
	for _, dependency := range projectDependencies {
		if dependency.Path == path && dependency.Project == project {
			return true
		}
	}

	return false
}
