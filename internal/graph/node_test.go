package graph

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/steps"
	"path/filepath"
	"testing"
)

func TestFromConfiguration(t *testing.T) {
	t.Run("ProjectConfig", func(t *testing.T) {
		stepsProviderOptions := steps.NewProviderOptions()
		stepsProviderOptions.Factories["foo"] = func(definition interface{}) (steps.Step, error) {
			return newTestStep("foo"), nil
		}
		stepsProvider := steps.NewProvider(stepsProviderOptions)
		path := "foo/foo.yml"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Path: "bar"},
					{Project: "baz"},
				},
				Steps: []config.Step{
					map[string]interface{}{"foo": "foo"},
				},
			},
		}

		node, err := FromConfiguration(&cfg, path, stepsProvider)

		assert.NoError(t, err)
		assert.NotEmpty(t, node)
		assert.Equal(t, cfg.ProjectTarget.Name, node.Name())
		assert.Equal(t, path, node.Path())

		dependencies := node.Dependencies()
		assert.NotEmpty(t, dependencies)
		assert.Len(t, dependencies, 2)
		assert.True(t, containsDependency(dependencies, ByPathDependencyKind, filepath.Join("foo", "bar")))
		assert.True(t, containsDependency(dependencies, ByNameDependencyKind, "baz"))
	})

	t.Run("WorkspaceConfig", func(t *testing.T) {
		path := "foo/foo.yml"
		cfg := config.Configuration{
			WorkspaceTarget: &config.WorkspaceTarget{
				Name: "foo",
				Includes: []string{
					"bar",
					"baz",
				},
			},
		}

		node, err := FromConfiguration(&cfg, path, nil)

		assert.NoError(t, err)
		assert.NotEmpty(t, node)
		assert.Equal(t, cfg.WorkspaceTarget.Name, node.Name())
		assert.Equal(t, path, node.Path())

		dependencies := node.Dependencies()
		assert.NotEmpty(t, dependencies)
		assert.Len(t, dependencies, 2)
		assert.True(t, containsDependency(dependencies, ByPathDependencyKind, filepath.Join("foo", "bar")))
		assert.True(t, containsDependency(dependencies, ByPathDependencyKind, filepath.Join("foo", "baz")))
	})
}

func containsDependency(dependencies []NodeDependency, kind NodeDependencyKind, value string) bool {
	for _, dependency := range dependencies {
		if dependency.Kind() == kind && dependency.Value() == value {
			return true
		}
	}

	return false
}

type testStep struct {
	name string
}

func (t *testStep) Name() string {
	return t.name
}

func newTestStep(name string) steps.Step {
	return &testStep{name: name}
}
