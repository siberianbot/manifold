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

		assert.NotEmpty(t, configuration.ProjectTarget)

		assert.Equal(t, "foo", configuration.ProjectTarget.Name)

		assert.NotEmpty(t, configuration.ProjectTarget.Steps)
		assert.Len(t, configuration.ProjectTarget.Steps, 3)
		assert.True(t, containsNamedStep(configuration.ProjectTarget.Steps, "foo"), "configuration.ProjectTarget.Steps doesn't contains foo")
		assert.True(t, containsNamedStep(configuration.ProjectTarget.Steps, "bar"), "configuration.ProjectTarget.Steps doesn't contains bar")
		assert.True(t, containsNamedStep(configuration.ProjectTarget.Steps, "baz"), "configuration.ProjectTarget.Steps doesn't contains baz")

		assert.NotEmpty(t, configuration.ProjectTarget.ProjectDependencies)
		assert.Len(t, configuration.ProjectTarget.ProjectDependencies, 2)
		assert.True(t, containsProjectDependency(configuration.ProjectTarget.ProjectDependencies, "", "bar"), "configuration.ProjectTarget.ProjectDependencies doesn't contains project bar")
		assert.True(t, containsProjectDependency(configuration.ProjectTarget.ProjectDependencies, "baz", ""), "configuration.ProjectTarget.ProjectDependencies doesn't contains path baz")
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

		assert.NotEmpty(t, configuration.WorkspaceTarget)

		assert.Equal(t, "foo", configuration.WorkspaceTarget.Name)

		assert.NotEmpty(t, configuration.WorkspaceTarget.Includes)
		assert.Len(t, configuration.WorkspaceTarget.Includes, 2)
		assert.True(t, containsInclude(configuration.WorkspaceTarget.Includes, "bar"), "configuration.WorkspaceTarget.Includes doesn't contains bar")
		assert.True(t, containsInclude(configuration.WorkspaceTarget.Includes, "baz"), "configuration.WorkspaceTarget.Includes doesn't contains baz")
	})
}

func containsNamedStep(steps []Step, name string) bool {
	for _, step := range steps {
		if step[name] == name {
			return true
		}
	}

	return false
}

func containsInclude(includes []string, include string) bool {
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
