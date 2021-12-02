package project_node

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	node2 "manifold/internal/graph/node"
	"manifold/internal/mock"
	"testing"
)

func TestProjectWithOnlyBasicFields(t *testing.T) {
	stepBuilder := new(mock.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Project: &config.Project{
			Name: name,
		},
	}

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())
	assert.Len(t, node.Dependencies(), 0)

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithValidStep(t *testing.T) {
	stepBuilder := new(mock.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	projectStep := config.ProjectStep{name: name}
	cfg := &config.Configuration{
		Project: &config.Project{
			Name: name,
			Steps: []config.ProjectStep{
				projectStep,
			},
		},
	}

	stepMock := new(mock.Step)
	stepBuilder.On("FromConfig", projectStep).Return(stepMock, nil)

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())

	projectNode := node.(*Node)
	assert.Len(t, projectNode.steps, 1)
	assert.Equal(t, stepMock, projectNode.steps[0])

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithInvalidStep(t *testing.T) {
	stepBuilder := new(mock.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	projectStep := config.ProjectStep{name: name}
	cfg := &config.Configuration{
		Project: &config.Project{
			Name: name,
			Steps: []config.ProjectStep{
				projectStep,
			},
		},
	}

	stepBuilderErr := errors.New("error")
	stepBuilder.On("FromConfig", projectStep).Return(nil, stepBuilderErr)

	node, err := builder.FromConfig(name, cfg)

	assert.Nil(t, node)
	assert.EqualError(t, err, fmt.Sprintf(invalidProjectStep, name, stepBuilderErr))

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithPathedDependency(t *testing.T) {
	stepBuilder := new(mock.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Project: &config.Project{
			Name:  name,
			Steps: []config.ProjectStep{},
			Dependencies: []config.ProjectDependency{
				{Path: name, Project: ""},
			},
		},
	}

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())

	dependencies := node.Dependencies()
	assert.Len(t, dependencies, 1)
	assert.Equal(t, node2.PathedDependencyKind, dependencies[0].Kind())
	assert.Equal(t, name, dependencies[0].Value())

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithNamedDependency(t *testing.T) {
	stepBuilder := new(mock.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Project: &config.Project{
			Name:  name,
			Steps: []config.ProjectStep{},
			Dependencies: []config.ProjectDependency{
				{Path: "", Project: name},
			},
		},
	}

	node, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, node)
	assert.Equal(t, name, node.Name())
	assert.Equal(t, name, node.Path())

	dependencies := node.Dependencies()
	assert.Len(t, dependencies, 1)
	assert.Equal(t, node2.NamedDependencyKind, dependencies[0].Kind())
	assert.Equal(t, name, dependencies[0].Value())

	stepBuilder.AssertExpectations(t)
}
