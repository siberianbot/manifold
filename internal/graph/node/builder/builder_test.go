package builder

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/mock"
	"testing"
)

func TestWithValidProject(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Project: &config.Project{}}

	mockNode := new(mock.Node)

	configLoader := new(mock.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(mock.NodeBuilder)
	projectNodeBuilder.On("FromConfig", path, cfg).Return(mockNode, nil)

	workspaceNodeBuilder := new(mock.NodeBuilder)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	n, err := builder.FromPath(path)

	assert.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, mockNode, n)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestWithValidWorkspace(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Workspace: &config.Workspace{}}

	mockNode := new(mock.Node)

	configLoader := new(mock.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(mock.NodeBuilder)

	workspaceNodeBuilder := new(mock.NodeBuilder)
	workspaceNodeBuilder.On("FromConfig", path, cfg).Return(mockNode, nil)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	n, err := builder.FromPath(path)

	assert.NoError(t, err)
	assert.NotNil(t, n)
	assert.Equal(t, mockNode, n)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestInvalidConfig(t *testing.T) {
	path := "foo"

	cfgErr := errors.New("error")

	configLoader := new(mock.ConfigLoader)
	configLoader.On("FromPath", path).Return(nil, cfgErr)

	projectNodeBuilder := new(mock.NodeBuilder)
	workspaceNodeBuilder := new(mock.NodeBuilder)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	node, err := builder.FromPath(path)

	assert.Nil(t, node)
	assert.Error(t, cfgErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestFailedProjectBuilder(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Project: &config.Project{}}

	builderErr := errors.New("error")

	configLoader := new(mock.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(mock.NodeBuilder)
	projectNodeBuilder.On("FromConfig", path, cfg).Return(nil, builderErr)

	workspaceNodeBuilder := new(mock.NodeBuilder)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	node, err := builder.FromPath(path)

	assert.Nil(t, node)
	assert.Error(t, builderErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestFailedWorkspaceBuilder(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Workspace: &config.Workspace{}}

	builderErr := errors.New("error")

	configLoader := new(mock.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(mock.NodeBuilder)

	workspaceNodeBuilder := new(mock.NodeBuilder)
	workspaceNodeBuilder.On("FromConfig", path, cfg).Return(nil, builderErr)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	node, err := builder.FromPath(path)

	assert.Nil(t, node)
	assert.Error(t, builderErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}
