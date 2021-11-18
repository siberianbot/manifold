package traversing

import (
	"manifold/internal/traversing/graph"
	"manifold/internal/traversing/targets"
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
	t.Run("CircularDependencies", testCircularDependencies)
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

		g, err := graph.Traverse(dirPath)

		test.Assert(t, g != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(g.Nodes) == 1)
		test.Assert(t, len(g.Links) == 0)
		test.Assert(t, g.Root != nil)
		test.Assert(t, g.Root.Target.Name() == "foo")
		test.Assert(t, g.Root.Target.Kind() == targets.ProjectTargetKind)
		test.Assert(t, g.Root.Target.Path() == path)
		test.Assert(t, len(g.Root.Dependencies) == 0)
	})

	t.Run("OneProject_NoDependencies_FilePath", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
project:
  name: foo
`

		test.CreateFile(t, path, foo)

		g, err := graph.Traverse(path)

		test.Assert(t, g != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(g.Nodes) == 1)
		test.Assert(t, len(g.Links) == 0)
		test.Assert(t, g.Root != nil)
		test.Assert(t, g.Root.Target.Name() == "foo")
		test.Assert(t, g.Root.Target.Kind() == targets.ProjectTargetKind)
		test.Assert(t, g.Root.Target.Path() == path)
		test.Assert(t, len(g.Root.Dependencies) == 0)
	})

	t.Run("OneWorkspace_NoIncludes", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "foo", "foo.manifold.yml")
		foo := `
workspace:
  name: foo
`

		test.CreateFile(t, path, foo)

		g, err := graph.Traverse(path)

		test.Assert(t, g != nil)
		test.Assert(t, err == nil)
		test.Assert(t, len(g.Nodes) == 1)
		test.Assert(t, len(g.Links) == 0)
		test.Assert(t, g.Root != nil)
		test.Assert(t, g.Root.Target.Name() == "foo")
		test.Assert(t, g.Root.Target.Kind() == targets.WorkspaceTargetKind)
		test.Assert(t, g.Root.Target.Path() == path)
		test.Assert(t, len(g.Root.Dependencies) == 0)
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

		g, err := graph.Traverse(fooDir)

		test.Assert(t, g != nil)
		test.Assert(t, err == nil)

		test.Assert(t, len(g.Nodes) == 2)
		test.Assert(t, len(g.Links) == 1)

		test.Assert(t, g.Root != nil)
		test.Assert(t, g.Root.Target.Name() == "foo")
		test.Assert(t, g.Root.Target.Kind() == targets.WorkspaceTargetKind)
		test.Assert(t, g.Root.Target.Path() == fooPath)
		test.Assert(t, len(g.Root.Dependencies) == 1)

		dependencyNode := g.FindNodeByName("bar")
		test.Assert(t, dependencyNode != nil)
		test.Assert(t, dependencyNode.Target.Name() == "bar")
		test.Assert(t, dependencyNode.Target.Kind() == targets.ProjectTargetKind)
		test.Assert(t, dependencyNode.Target.Path() == barPath)
		test.Assert(t, len(dependencyNode.Dependencies) == 0)

		test.Assert(t, g.FindLink(g.Root, dependencyNode) != nil)
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

		g, err := graph.Traverse(fooDir)

		test.Assert(t, g != nil)
		test.Assert(t, err == nil)

		test.Assert(t, len(g.Nodes) == 3)
		test.Assert(t, len(g.Links) == 3)

		workspaceNode := g.Root
		test.Assert(t, workspaceNode != nil)
		test.Assert(t, workspaceNode.Target.Name() == "foo")
		test.Assert(t, workspaceNode.Target.Kind() == targets.WorkspaceTargetKind)
		test.Assert(t, workspaceNode.Target.Path() == fooPath)
		test.Assert(t, len(workspaceNode.Dependencies) == 2)

		barNode := g.FindNodeByName("bar")
		test.Assert(t, barNode != nil)
		test.Assert(t, barNode.Target.Name() == "bar")
		test.Assert(t, barNode.Target.Kind() == targets.ProjectTargetKind)
		test.Assert(t, barNode.Target.Path() == barPath)
		test.Assert(t, len(barNode.Dependencies) == 1)

		bazNode := g.FindNodeByName("baz")
		test.Assert(t, bazNode != nil)
		test.Assert(t, bazNode.Target.Name() == "baz")
		test.Assert(t, bazNode.Target.Kind() == targets.ProjectTargetKind)
		test.Assert(t, bazNode.Target.Path() == bazPath)
		test.Assert(t, len(bazNode.Dependencies) == 0)

		test.Assert(t, g.FindLink(workspaceNode, barNode) != nil)
		test.Assert(t, g.FindLink(workspaceNode, bazNode) != nil)
		test.Assert(t, g.FindLink(barNode, bazNode) != nil)
	})
}

func testCircularDependencies(t *testing.T) {
	t.Run("Project_SelfReferences", func(t *testing.T) {
		dirPath := filepath.Join(t.TempDir(), "foo")
		path := filepath.Join(dirPath, ".manifold.yml")
		foo := `
project:
  name: foo
  dependencies:
  - project: foo
`

		test.CreateFile(t, path, foo)

		g, err := graph.Traverse(dirPath)

		test.Assert(t, g == nil)
		test.Assert(t, err != nil)
	})

	t.Run("Project_TransitiveReferences", func(t *testing.T) {
		fooDir := filepath.Join(t.TempDir(), "foo")
		fooPath := filepath.Join(fooDir, ".manifold.yml")
		foo := `
project:
  name: foo
  dependencies:
  - project: bar
`

		test.CreateFile(t, fooPath, foo)

		barPath := filepath.Join(fooDir, "bar", ".manifold.yml")
		bar := `
project:
  name: bar
  dependencies:
  - project: foo
`
		test.CreateFile(t, barPath, bar)

		g, err := graph.Traverse(fooDir)

		test.Assert(t, g == nil)
		test.Assert(t, err != nil)
	})
}
