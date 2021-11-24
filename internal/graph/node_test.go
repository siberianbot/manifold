package graph

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"path/filepath"
	"testing"
)

func TestFromConfiguration(t *testing.T) {
	t.Run("ProjectConfig", func(t *testing.T) {
		path := "foo/foo.yml"
		cfg := config.Configuration{
			ProjectTarget: &config.ProjectTarget{
				Name: "foo",
				ProjectDependencies: []config.ProjectDependency{
					{Path: "bar"},
					{Project: "baz"},
				},
			},
		}

		node := FromConfiguration(&cfg, path)

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

		node := FromConfiguration(&cfg, path)

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
