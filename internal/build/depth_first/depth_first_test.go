package depth_first

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"manifold/internal/graph/node"
	mockGraph "manifold/internal/mock/graph"
	mockNode "manifold/internal/mock/node"
	mockStep "manifold/internal/mock/step"
	"manifold/internal/step"
	"testing"
)

func TestDepthFirst(t *testing.T) {
	stepName := "foo"
	s := new(mockStep.Step)
	s.On("Name").Return(stepName)

	root := new(mockNode.Node)
	root.On("Steps").Return([]step.Step{s})

	g := new(mockGraph.Graph)
	g.On("Root").Return(root)
	g.On("DescendantsOf", root).Return([]node.Node{})

	executor := new(mockStep.StepExecutor)
	executor.On("Execute", mock.Anything).Return(nil)

	provider := new(mockStep.StepProvider)
	provider.On("ExecutorFor", stepName).Return(executor)

	strategy := NewDepthFirstBuildStrategy(provider)

	assert.NotNil(t, strategy)

	err := strategy.Build(g)

	assert.NoError(t, err)

	g.AssertExpectations(t)
	executor.AssertExpectations(t)
	provider.AssertExpectations(t)
}
