package builder

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	mockConfig "manifold/internal/mock/config"
	"manifold/internal/mock/node"
	"testing"
)

func TestWithValidProject(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Project: &config.Project{}}

	mockNode := new(node.Node)

	configLoader := new(mockConfig.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(node.ConcreteNodeBuilder)
	projectNodeBuilder.On("FromConfig", path, cfg).Return(mockNode, nil)

	workspaceNodeBuilder := new(node.ConcreteNodeBuilder)

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

	mockNode := new(node.Node)

	configLoader := new(mockConfig.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(node.ConcreteNodeBuilder)

	workspaceNodeBuilder := new(node.ConcreteNodeBuilder)
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

	configLoader := new(mockConfig.ConfigLoader)
	configLoader.On("FromPath", path).Return(nil, cfgErr)

	projectNodeBuilder := new(node.ConcreteNodeBuilder)
	workspaceNodeBuilder := new(node.ConcreteNodeBuilder)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	n, err := builder.FromPath(path)

	assert.Nil(t, n)
	assert.Error(t, cfgErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestFailedProjectBuilder(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Project: &config.Project{}}

	builderErr := errors.New("error")

	configLoader := new(mockConfig.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(node.ConcreteNodeBuilder)
	projectNodeBuilder.On("FromConfig", path, cfg).Return(nil, builderErr)

	workspaceNodeBuilder := new(node.ConcreteNodeBuilder)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	n, err := builder.FromPath(path)

	assert.Nil(t, n)
	assert.Error(t, builderErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}

func TestFailedWorkspaceBuilder(t *testing.T) {
	path := "foo"
	cfg := &config.Configuration{Workspace: &config.Workspace{}}

	builderErr := errors.New("error")

	configLoader := new(mockConfig.ConfigLoader)
	configLoader.On("FromPath", path).Return(cfg, nil)

	projectNodeBuilder := new(node.ConcreteNodeBuilder)

	workspaceNodeBuilder := new(node.ConcreteNodeBuilder)
	workspaceNodeBuilder.On("FromConfig", path, cfg).Return(nil, builderErr)

	builder := NewBuilder(configLoader, projectNodeBuilder, workspaceNodeBuilder)

	assert.NotNil(t, builder)

	n, err := builder.FromPath(path)

	assert.Nil(t, n)
	assert.Error(t, builderErr, err)

	configLoader.AssertExpectations(t)
	projectNodeBuilder.AssertExpectations(t)
	workspaceNodeBuilder.AssertExpectations(t)
}
