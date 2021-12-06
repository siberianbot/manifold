package backend

import (
	"manifold/internal/build"
	"manifold/internal/build/depth_first"
	"manifold/internal/builtin/steps/command_step"
	"manifold/internal/builtin/steps/dotnet_step"
	configLoader "manifold/internal/config/loader"
	graphBuilder "manifold/internal/graph/builder"
	"manifold/internal/graph/node"
	nodeBuilder "manifold/internal/graph/node/builder"
	"manifold/internal/graph/node/project_node"
	"manifold/internal/graph/node/workspace_node"
	graphValidator "manifold/internal/graph/validator"
	"manifold/internal/graph/validator/cycle_detection"
	"manifold/internal/step"
	stepBuilder "manifold/internal/step/builder"
	stepProvider "manifold/internal/step/provider"
)

type BuildOptions struct {
	Path string
}

type Interface interface {
	Build(options BuildOptions) error
}

type backend struct {
	configLoader            configLoader.Interface
	stepProvider            stepProvider.Interface
	stepBuilder             stepBuilder.Interface
	nodeBuilder             nodeBuilder.Interface
	projectNodeBuilder      node.Builder
	workspaceNodeBuilder    node.Builder
	graphValidator          graphValidator.Interface
	graphBuilder            graphBuilder.Interface
	buildStrategy           build.Interface
	cycleDetectionAlgorithm cycle_detection.CycleDetectionAlgorithm
}

func NewBackend() Interface {
	stepProviderOptions := stepProvider.Options{
		Factories: map[string]step.Factory{},
		Executors: map[string]step.Executor{},
	}

	command_step.PopulateOptions(&stepProviderOptions)
	dotnet_step.PopulateOptions(&stepProviderOptions)

	b := &backend{
		configLoader:            configLoader.NewLoader(),
		stepProvider:            stepProvider.NewProvider(stepProviderOptions),
		cycleDetectionAlgorithm: cycle_detection.GetDefaultCycleDetectionAlgorithm(),
	}

	b.stepBuilder = stepBuilder.NewBuilder(b.stepProvider)
	b.projectNodeBuilder = project_node.NewBuilder(b.stepBuilder)
	b.workspaceNodeBuilder = workspace_node.NewBuilder()
	b.nodeBuilder = nodeBuilder.NewBuilder(b.configLoader, b.projectNodeBuilder, b.workspaceNodeBuilder)
	b.graphValidator = graphValidator.NewValidator(b.cycleDetectionAlgorithm)
	b.graphBuilder = graphBuilder.NewBuilder(b.nodeBuilder, b.graphValidator)
	b.buildStrategy = depth_first.NewDepthFirstBuildStrategy(b.stepProvider)

	return b
}

func (b *backend) Build(options BuildOptions) error {
	graph, err := b.graphBuilder.FromPath(options.Path)

	if err != nil {
		return err
	}

	return b.buildStrategy.Build(graph)
}
