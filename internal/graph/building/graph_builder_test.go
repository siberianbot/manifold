package building

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"manifold/internal/steps"
	"os"
	"path/filepath"
	"testing"
)

func TestGraphBuilder(t *testing.T) {
	t.Run("NewGraphBuilder", func(t *testing.T) {
		stepsProvider := steps.NewProvider(steps.NewProviderOptions())
		nodeBuilder := NewNodeBuilder(stepsProvider)
		graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

		assert.NotNil(t, graphBuilder)
		assert.Equal(t, nodeBuilder, graphBuilder.options.NodeBuilder)
	})

	t.Run("Build", func(t *testing.T) {
		t.Run("ProjectWithoutDependencies", func(t *testing.T) {
			dir := t.TempDir()

			fooPath := filepath.Join(dir, ".manifold.yml")
			fooContent := `
project:
  name: foo
`
			fooFile, _ := os.Create(fooPath)
			_, _ = fooFile.WriteString(fooContent)
			_ = fooFile.Close()

			stepsProvider := steps.NewProvider(steps.NewProviderOptions())
			nodeBuilder := NewNodeBuilder(stepsProvider)
			graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

			dependencyGraph, err := graphBuilder.Build(fooPath)

			assert.NoError(t, err)
			assert.NotEmpty(t, dependencyGraph)

			root := dependencyGraph.Root()
			assert.NotEmpty(t, root)
			assert.Equal(t, "foo", root.Name())
			assert.Equal(t, fooPath, root.Path())
			assert.Empty(t, root.Dependencies())

			nodes := dependencyGraph.Nodes()
			assert.NotEmpty(t, nodes)
			assert.Len(t, nodes, 1)
			assert.Equal(t, root, nodes[0])
		})

		t.Run("ProjectWithSinglePathDependency", func(t *testing.T) {
			dir := t.TempDir()

			fooPath := filepath.Join(dir, ".manifold.yml")
			fooContent := `
project:
  name: foo
  dependencies:
  - path: bar/
`
			fooFile, _ := os.Create(fooPath)
			_, _ = fooFile.WriteString(fooContent)
			_ = fooFile.Close()

			barDir := filepath.Join(dir, "bar")
			barPath := filepath.Join(barDir, ".manifold.yml")
			barContent := `
project:
  name: bar
`
			_ = os.Mkdir(barDir, os.ModeDir)
			barFile, _ := os.Create(barPath)
			_, _ = barFile.WriteString(barContent)
			_ = barFile.Close()

			stepsProvider := steps.NewProvider(steps.NewProviderOptions())
			nodeBuilder := NewNodeBuilder(stepsProvider)
			graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

			dependencyGraph, err := graphBuilder.Build(fooPath)

			assert.NoError(t, err)
			assert.NotEmpty(t, dependencyGraph)

			nodes := dependencyGraph.Nodes()
			assert.NotEmpty(t, nodes)
			assert.Len(t, nodes, 2)

			root := dependencyGraph.Root()
			assert.NotEmpty(t, root)
			assert.Equal(t, "foo", root.Name())
			assert.Equal(t, fooPath, root.Path())

			dependencies := root.Dependencies()
			assert.NotEmpty(t, dependencies)
			assert.Len(t, dependencies, 1)

			rootDescendants := dependencyGraph.Descendants(root)
			assert.NotEmpty(t, rootDescendants)
			assert.Len(t, rootDescendants, 1)
			assert.Equal(t, "bar", rootDescendants[0].Name())
			assert.Equal(t, barPath, rootDescendants[0].Path())
		})

		t.Run("ProjectWithManyPathDependencies", func(t *testing.T) {
			dir := t.TempDir()

			fooPath := filepath.Join(dir, ".manifold.yml")
			fooContent := `
project:
  name: foo
  dependencies:
  - path: bar/
  - path: baz/
`
			fooFile, _ := os.Create(fooPath)
			_, _ = fooFile.WriteString(fooContent)
			_ = fooFile.Close()

			barDir := filepath.Join(dir, "bar")
			barPath := filepath.Join(barDir, ".manifold.yml")
			barContent := `
project:
  name: bar
`
			_ = os.Mkdir(barDir, os.ModeDir)
			barFile, _ := os.Create(barPath)
			_, _ = barFile.WriteString(barContent)
			_ = barFile.Close()

			bazDir := filepath.Join(dir, "baz")
			bazPath := filepath.Join(bazDir, ".manifold.yml")
			bazContent := `
project:
  name: baz
`
			_ = os.Mkdir(bazDir, os.ModeDir)
			bazFile, _ := os.Create(bazPath)
			_, _ = bazFile.WriteString(bazContent)
			_ = bazFile.Close()

			stepsProvider := steps.NewProvider(steps.NewProviderOptions())
			nodeBuilder := NewNodeBuilder(stepsProvider)
			graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

			dependencyGraph, err := graphBuilder.Build(fooPath)

			assert.NoError(t, err)
			assert.NotEmpty(t, dependencyGraph)

			nodes := dependencyGraph.Nodes()
			assert.NotEmpty(t, nodes)
			assert.Len(t, nodes, 3)

			root := dependencyGraph.Root()
			assert.NotEmpty(t, root)
			assert.Equal(t, "foo", root.Name())
			assert.Equal(t, fooPath, root.Path())

			dependencies := root.Dependencies()
			assert.NotEmpty(t, dependencies)
			assert.Len(t, dependencies, 2)

			rootDescendants := dependencyGraph.Descendants(root)
			assert.NotEmpty(t, rootDescendants)
			assert.Len(t, rootDescendants, 2)

			barNode := dependencyGraph.FindByName("bar")
			assert.NotEmpty(t, barNode)
			assert.Equal(t, "bar", barNode.Name())
			assert.Equal(t, barPath, barNode.Path())

			bazNode := dependencyGraph.FindByName("baz")
			assert.NotEmpty(t, bazNode)
			assert.Equal(t, "baz", bazNode.Name())
			assert.Equal(t, bazPath, bazNode.Path())
		})

		t.Run("ProjectWithDifferentDependencies", func(t *testing.T) {
			dir := t.TempDir()

			fooPath := filepath.Join(dir, ".manifold.yml")
			fooContent := `
project:
  name: foo
  dependencies:
  - project: bar
  - path: baz/
`
			fooFile, _ := os.Create(fooPath)
			_, _ = fooFile.WriteString(fooContent)
			_ = fooFile.Close()

			bazDir := filepath.Join(dir, "baz")
			bazPath := filepath.Join(bazDir, ".manifold.yml")
			bazContent := `
project:
  name: baz
  dependencies:
  - path: bar/
`
			_ = os.Mkdir(bazDir, os.ModeDir)
			bazFile, _ := os.Create(bazPath)
			_, _ = bazFile.WriteString(bazContent)
			_ = bazFile.Close()

			barDir := filepath.Join(bazDir, "bar")
			barPath := filepath.Join(barDir, ".manifold.yml")
			barContent := `
project:
  name: bar
`
			_ = os.Mkdir(barDir, os.ModeDir)
			barFile, _ := os.Create(barPath)
			_, _ = barFile.WriteString(barContent)
			_ = barFile.Close()

			stepsProvider := steps.NewProvider(steps.NewProviderOptions())
			nodeBuilder := NewNodeBuilder(stepsProvider)
			graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

			dependencyGraph, err := graphBuilder.Build(fooPath)

			assert.NoError(t, err)
			assert.NotEmpty(t, dependencyGraph)

			nodes := dependencyGraph.Nodes()
			assert.NotEmpty(t, nodes)
			assert.Len(t, nodes, 3)

			root := dependencyGraph.Root()
			assert.NotEmpty(t, root)
			assert.Equal(t, "foo", root.Name())
			assert.Equal(t, fooPath, root.Path())

			dependencies := root.Dependencies()
			assert.NotEmpty(t, dependencies)
			assert.Len(t, dependencies, 2)

			rootDescendants := dependencyGraph.Descendants(root)
			assert.NotEmpty(t, rootDescendants)
			assert.Len(t, rootDescendants, 2)

			bazNode := dependencyGraph.FindByName("baz")
			assert.NotEmpty(t, bazNode)
			assert.Equal(t, "baz", bazNode.Name())
			assert.Equal(t, bazPath, bazNode.Path())

			bazDescendants := dependencyGraph.Descendants(bazNode)
			assert.NotEmpty(t, bazDescendants)
			assert.Len(t, bazDescendants, 1)

			barNode := dependencyGraph.FindByName("bar")
			assert.NotEmpty(t, barNode)
			assert.Equal(t, "bar", barNode.Name())
			assert.Equal(t, barPath, barNode.Path())

			barDescendants := dependencyGraph.Descendants(barNode)
			assert.Empty(t, barDescendants)

			assert.Contains(t, rootDescendants, barNode)
			assert.Contains(t, rootDescendants, bazNode)
			assert.Contains(t, bazDescendants, barNode)
		})

		t.Run("InvalidProject", func(t *testing.T) {
			dir := t.TempDir()

			fooPath := filepath.Join(dir, ".manifold.yml")
			fooContent := `
project:
  name: foo
  dependencies:
  - project: bar
`
			fooFile, _ := os.Create(fooPath)
			_, _ = fooFile.WriteString(fooContent)
			_ = fooFile.Close()

			stepsProvider := steps.NewProvider(steps.NewProviderOptions())
			nodeBuilder := NewNodeBuilder(stepsProvider)
			graphBuilder := NewGraphBuilder(GraphBuilderOptions{NodeBuilder: nodeBuilder})

			dependencyGraph, err := graphBuilder.Build(fooPath)

			assert.Nil(t, dependencyGraph)
			assert.EqualError(t, err, fmt.Sprintf(UnknownTarget, "bar"))
		})
	})
}
