package workspace_node

import (
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	node2 "manifold/internal/graph/node"
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

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())
	assert.Len(t, node.Dependencies(), 0)
}

func TestWorkspaceWithIncludes(t *testing.T) {
	builder := NewBuilder()

	assert.NotNil(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Workspace: &config.Workspace{
			Name:     name,
			Includes: []config.WorkspaceInclude{config.WorkspaceInclude(name)},
		},
	}

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())

	includes := node.Dependencies()
	assert.Len(t, includes, 1)
	assert.Equal(t, node2.PathedDependencyKind, includes[0].Kind())
	assert.Equal(t, name, includes[0].Value())
}
