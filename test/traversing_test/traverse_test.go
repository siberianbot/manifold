package traversing

import (
	"manifold/internal/traversing/build_info"
	"manifold/internal/traversing/graph"
	"manifold/test"
	"path/filepath"
	"testing"
)

// TODO:
// 0. successful cases
// 1. circular dependencies: A -> B -> A, A -> B -> C -> A, and so on;
// 2. common failures: invalid names, empty documents;
// 3. (?) logging

func TestTraverse(t *testing.T) {
	t.Run("SuccessfulCases", testSuccessfulCases)
}

func testSuccessfulCases(t *testing.T) {
	t.Run("OneProject_NoDependencies_DirPath", func(t *testing.T) {
		dirPath := filepath.Join(t.TempDir(), "foo")
		path := filepath.Join(dirPath, ".manifold.yml")
		foo := `
project:
  name: foo
`

		test.CreateFile(t, path, foo)

		nodeCollection, err := graph.Traverse(dirPath)

		test.Assert(t, nodeCollection != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(nodeCollection.Nodes) == 1)
		test.Assert(t, len(nodeCollection.Links) == 0)
		test.Assert(t, nodeCollection.Root != nil)
		test.Assert(t, nodeCollection.Root.BuildInfo.Name() == "foo")
		test.Assert(t, nodeCollection.Root.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, nodeCollection.Root.BuildInfo.Path() == path)
		test.Assert(t, len(nodeCollection.Root.Dependencies) == 0)
	})

	t.Run("OneProject_NoDependencies_FilePath", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
project:
  name: foo
`

		test.CreateFile(t, path, foo)

		nodeCollection, err := graph.Traverse(path)

		test.Assert(t, nodeCollection != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(nodeCollection.Nodes) == 1)
		test.Assert(t, len(nodeCollection.Links) == 0)
		test.Assert(t, nodeCollection.Root != nil)
		test.Assert(t, nodeCollection.Root.BuildInfo.Name() == "foo")
		test.Assert(t, nodeCollection.Root.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, nodeCollection.Root.BuildInfo.Path() == path)
		test.Assert(t, len(nodeCollection.Root.Dependencies) == 0)
	})

	t.Run("OneWorkspace_NoIncludes", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
workspace:
  name: foo
`

		test.CreateFile(t, path, foo)

		nodeCollection, err := graph.Traverse(path)

		test.Assert(t, nodeCollection != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(nodeCollection.Nodes) == 1)
		test.Assert(t, len(nodeCollection.Links) == 0)
		test.Assert(t, nodeCollection.Root != nil)
		test.Assert(t, nodeCollection.Root.BuildInfo.Name() == "foo")
		test.Assert(t, nodeCollection.Root.BuildInfo.Kind() == build_info.WorkspaceBuildInfoKind)
		test.Assert(t, nodeCollection.Root.BuildInfo.Path() == path)
		test.Assert(t, len(nodeCollection.Root.Dependencies) == 0)
	})

	t.Run("OneWorkspace_OneDependentProject", func(t *testing.T) {
		tempDir := t.TempDir()
		fooDir := filepath.Join(tempDir, "foo")
		fooPath := filepath.Join(fooDir, ".manifold.yml")
		foo := `
workspace:
  name: foo
  includes:
  - bar/
`
		test.CreateFile(t, fooPath, foo)

		barPath := filepath.Join(fooDir, "bar", ".manifold.yml")
		bar := `
project:
  name: bar
`
		test.CreateFile(t, barPath, bar)

		nodeCollection, err := graph.Traverse(fooDir)

		test.Assert(t, nodeCollection != nil)
		test.Assert(t, err == nil)

		test.Assert(t, len(nodeCollection.Nodes) == 2)
		test.Assert(t, len(nodeCollection.Links) == 1)

		test.Assert(t, nodeCollection.Root != nil)
		test.Assert(t, nodeCollection.Root.BuildInfo.Name() == "foo")
		test.Assert(t, nodeCollection.Root.BuildInfo.Kind() == build_info.WorkspaceBuildInfoKind)
		test.Assert(t, nodeCollection.Root.BuildInfo.Path() == fooPath)
		test.Assert(t, len(nodeCollection.Root.Dependencies) == 1)

		dependencyNode := findNode(nodeCollection, "bar")
		test.Assert(t, dependencyNode != nil)
		test.Assert(t, dependencyNode.BuildInfo.Name() == "bar")
		test.Assert(t, dependencyNode.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, dependencyNode.BuildInfo.Path() == barPath)
		test.Assert(t, len(dependencyNode.Dependencies) == 0)

		test.Assert(t, findLink(nodeCollection, nodeCollection.Root, dependencyNode) != nil)
	})

	t.Run("OneWorkspace_TwoDependentProject_OneCommonProject", func(t *testing.T) {
		tempDir := t.TempDir()
		fooDir := filepath.Join(tempDir, "foo")
		fooPath := filepath.Join(fooDir, ".manifold.yml")
		foo := `
workspace:
  name: foo
  includes:
  - bar/
  - baz/
`
		test.CreateFile(t, fooPath, foo)

		barPath := filepath.Join(fooDir, "bar", ".manifold.yml")
		bar := `
project:
  name: bar
  dependencies:
  - project: baz
`
		test.CreateFile(t, barPath, bar)

		bazPath := filepath.Join(fooDir, "baz", ".manifold.yml")
		baz := `
project:
  name: baz
`
		test.CreateFile(t, bazPath, baz)

		nodeCollection, err := graph.Traverse(fooDir)

		test.Assert(t, nodeCollection != nil)
		test.Assert(t, err == nil)

		test.Assert(t, len(nodeCollection.Nodes) == 3)
		test.Assert(t, len(nodeCollection.Links) == 3)

		workspaceNode := nodeCollection.Root
		test.Assert(t, workspaceNode != nil)
		test.Assert(t, workspaceNode.BuildInfo.Name() == "foo")
		test.Assert(t, workspaceNode.BuildInfo.Kind() == build_info.WorkspaceBuildInfoKind)
		test.Assert(t, workspaceNode.BuildInfo.Path() == fooPath)
		test.Assert(t, len(workspaceNode.Dependencies) == 2)

		barNode := findNode(nodeCollection, "bar")
		test.Assert(t, barNode != nil)
		test.Assert(t, barNode.BuildInfo.Name() == "bar")
		test.Assert(t, barNode.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, barNode.BuildInfo.Path() == barPath)
		test.Assert(t, len(barNode.Dependencies) == 1)

		bazNode := findNode(nodeCollection, "baz")
		test.Assert(t, bazNode != nil)
		test.Assert(t, bazNode.BuildInfo.Name() == "baz")
		test.Assert(t, bazNode.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, bazNode.BuildInfo.Path() == bazPath)
		test.Assert(t, len(bazNode.Dependencies) == 0)

		test.Assert(t, findLink(nodeCollection, workspaceNode, barNode) != nil)
		test.Assert(t, findLink(nodeCollection, workspaceNode, bazNode) != nil)
		test.Assert(t, findLink(nodeCollection, barNode, bazNode) != nil)
	})
}

// TODO: move to node collection, make node collection as receiver
func findNode(collection *graph.NodeCollection, name string) *graph.Node {
	for _, node := range collection.Nodes {
		if node.BuildInfo.Name() == name {
			return node
		}
	}

	return nil
}

// TODO: move to node collection, make node collection as receiver
func findLink(collection *graph.NodeCollection, parent *graph.Node, child *graph.Node) *graph.NodeLink {
	for _, link := range collection.Links {
		if link.Parent == parent && link.Child == child {
			return link
		}
	}

	return nil
}
