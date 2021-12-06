package workspace_node

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/graph/node"
	"os"
	"path/filepath"
	"testing"
)

func TestWorkspaceWithOnlyBasicFields(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Workspace: &config.Workspace{
			Name: name,
		},
	}

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())
	assert.Len(t, n.Dependencies(), 0)
}

func TestWorkspaceWithInvalidIncludes(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Workspace: &config.Workspace{
			Name:     name,
			Includes: []config.WorkspaceInclude{config.WorkspaceInclude(name)},
		},
	}

	n, err := builder.FromConfig(name, cfg)

	assert.Error(t, err)
	assert.Nil(t, n)
}

func TestWorkspaceWithValidIncludes(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)

	includeDir := t.TempDir()
	includePath := filepath.Join(includeDir, ".manifold.yml")
	file, _ := os.Create(includePath)
	_ = file.Close()

	name := "foo"
	cfg := &config.Configuration{
		Workspace: &config.Workspace{
			Name:     name,
			Includes: []config.WorkspaceInclude{config.WorkspaceInclude(includeDir)},
		},
	}

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())

	includes := n.Dependencies()
	assert.Len(t, includes, 1)
	assert.Equal(t, node.PathedDependencyKind, includes[0].Kind())
	assert.Equal(t, includePath, includes[0].Value())
}
