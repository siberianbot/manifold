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

		node, err := graph.Traverse(dirPath)

		test.Assert(t, node != nil)
		test.Assert(t, err == nil)
		test.Assert(t, node.BuildInfo.Name() == "foo")
		test.Assert(t, node.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, node.BuildInfo.Path() == path)
		test.Assert(t, len(node.Dependencies) == 0)
	})

	t.Run("OneProject_NoDependencies_FilePath", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
project:
  name: foo
`

		test.CreateFile(t, path, foo)

		node, err := graph.Traverse(path)

		test.Assert(t, node != nil)
		test.Assert(t, err == nil)
		test.Assert(t, node.BuildInfo.Name() == "foo")
		test.Assert(t, node.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, node.BuildInfo.Path() == path)
		test.Assert(t, len(node.Dependencies) == 0)
	})

	t.Run("OneWorkspace_NoIncludes", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
workspace:
  name: foo
`

		test.CreateFile(t, path, foo)

		node, err := graph.Traverse(path)

		test.Assert(t, node != nil)
		test.Assert(t, err == nil)
		test.Assert(t, node.BuildInfo.Name() == "foo")
		test.Assert(t, node.BuildInfo.Kind() == build_info.WorkspaceBuildInfoKind)
		test.Assert(t, node.BuildInfo.Path() == path)
		test.Assert(t, len(node.Dependencies) == 0)
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

		node, err := graph.Traverse(fooDir)

		test.Assert(t, node != nil)
		test.Assert(t, err == nil)
		test.Assert(t, node.BuildInfo.Name() == "foo")
		test.Assert(t, node.BuildInfo.Kind() == build_info.WorkspaceBuildInfoKind)
		test.Assert(t, node.BuildInfo.Path() == fooPath)
		test.Assert(t, len(node.Dependencies) == 1)

		dependencyNode := node.Dependencies[0]
		test.Assert(t, dependencyNode != nil)
		test.Assert(t, dependencyNode.BuildInfo.Name() == "bar")
		test.Assert(t, dependencyNode.BuildInfo.Kind() == build_info.ProjectBuildInfoKind)
		test.Assert(t, dependencyNode.BuildInfo.Path() == barPath)
		test.Assert(t, len(dependencyNode.Dependencies) == 0)
	})
}
