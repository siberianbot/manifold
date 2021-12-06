package project_node

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/config"
	"manifold/internal/graph/node"
	"manifold/internal/mock/step"
	"os"
	"path/filepath"
	"testing"
)

func TestProjectWithOnlyBasicFields(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	name := "foo"
	cfg := &config.Configuration{
		Project: &config.Project{
			Name: name,
		},
	}

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())
	assert.Len(t, n.Dependencies(), 0)

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithValidStep(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

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

	stepMock := new(step.Step)
	stepBuilder.On("FromConfig", projectStep).Return(stepMock, nil)

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())

	projectNode := n.(*Node)
	assert.Len(t, projectNode.steps, 1)
	assert.Equal(t, stepMock, projectNode.steps[0])

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithInvalidStep(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

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

	n, err := builder.FromConfig(name, cfg)

	assert.Nil(t, n)
	assert.EqualError(t, err, fmt.Sprintf(invalidProjectStep, name, stepBuilderErr))

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithValidPathedDependency(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

	builder := NewBuilder(stepBuilder)

	assert.NotEmpty(t, builder)

	depDir := t.TempDir()
	depPath := filepath.Join(depDir, ".manifold.yml")
	file, _ := os.Create(depPath)
	_ = file.Close()

	name := "foo"
	cfg := &config.Configuration{
		Project: &config.Project{
			Name:  name,
			Steps: []config.ProjectStep{},
			Dependencies: []config.ProjectDependency{
				{Path: depDir, Project: ""},
			},
		},
	}

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())

	dependencies := n.Dependencies()
	assert.Len(t, dependencies, 1)
	assert.Equal(t, node.PathedDependencyKind, dependencies[0].Kind())
	assert.Equal(t, depPath, dependencies[0].Value())

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithInvalidPathedDependency(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

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

	n, err := builder.FromConfig(name, cfg)

	assert.Error(t, err)
	assert.Nil(t, n)

	stepBuilder.AssertExpectations(t)
}

func TestProjectWithNamedDependency(t *testing.T) {
	stepBuilder := new(step.StepBuilder)

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

	n, err := builder.FromConfig(name, cfg)

	assert.NoError(t, err)
	assert.NotEmpty(t, n)
	assert.Equal(t, name, n.Name())
	assert.Equal(t, name, n.Path())

	dependencies := n.Dependencies()
	assert.Len(t, dependencies, 1)
	assert.Equal(t, node.NamedDependencyKind, dependencies[0].Kind())
	assert.Equal(t, name, dependencies[0].Value())

	stepBuilder.AssertExpectations(t)
}
