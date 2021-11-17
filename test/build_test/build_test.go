package build_test

import (
	"manifold/internal/build"
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestBuild(t *testing.T) {
	t.Run("SuccessfulCases", testSuccessfulCases)
}

func testSuccessfulCases(t *testing.T) {
	t.Run("OneProject", func(t *testing.T) {
		fooPath := filepath.Join(t.TempDir(), "foo", ".manifold.yml")
		foo := `
project:
  name: foo
  steps:
  - cmd: echo "Hello, World!"
`

		test.CreateFile(t, fooPath, foo)

		err := build.Build(fooPath)

		test.Assert(t, err == nil)
	})

	t.Run("ManyProjects", func(t *testing.T) {
		fooDir := filepath.Join(t.TempDir(), "foo")
		fooPath := filepath.Join(fooDir, ".manifold.yml")
		foo := `
project:
  name: foo
  dependencies:
  - path: bar/
  steps:
  - cmd: echo "foo"
`

		test.CreateFile(t, fooPath, foo)

		barPath := filepath.Join(fooDir, "bar", ".manifold.yml")
		bar := `
project:
  name: bar
  steps:
  - cmd: echo "bar"
`

		test.CreateFile(t, barPath, bar)

		err := build.Build(fooPath)

		test.Assert(t, err == nil)
	})

	t.Run("ManyProjectsInWorkspace", func(t *testing.T) {
		workspaceDir := filepath.Join(t.TempDir(), "workspace")
		workspacePath := filepath.Join(workspaceDir, ".manifold.yml")
		workspace := `
workspace:
  name: workspace
  includes:
  - foo/
  - bar/
  - baz/
`

		test.CreateFile(t, workspacePath, workspace)

		fooPath := filepath.Join(workspaceDir, "foo", ".manifold.yml")
		foo := `
project:
  name: foo
  dependencies:
  - project: bar
  - project: baz
  steps:
  - cmd: echo "foo"
`

		test.CreateFile(t, fooPath, foo)

		barPath := filepath.Join(workspaceDir, "bar", ".manifold.yml")
		bar := `
project:
  name: bar
  dependencies:
  - project: baz
  steps:
  - cmd: echo "bar"
`

		test.CreateFile(t, barPath, bar)

		bazPath := filepath.Join(workspaceDir, "baz", ".manifold.yml")
		baz := `
project:
  name: baz
  steps:
  - cmd: echo "baz"
`

		test.CreateFile(t, bazPath, baz)

		err := build.Build(workspaceDir)

		test.Assert(t, err == nil)
	})
}
