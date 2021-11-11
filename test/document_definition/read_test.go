package document_definition

import (
	"bytes"
	"manifold/internal/document_definition"
	"manifold/test"
	"math/rand"
	"strings"
	"testing"
)

func TestRead(t *testing.T) {
	t.Run("EmptyFile", func(t *testing.T) {
		reader := strings.NewReader("")
		doc, err := document_definition.Read(reader)

		test.AssertNil(t, doc)
		test.AssertNotNil(t, err)
	})

	t.Run("RandomFile", func(t *testing.T) {
		data := make([]byte, 1024)
		rand.Read(data)

		reader := bytes.NewReader(data)
		doc, err := document_definition.Read(reader)

		test.AssertNil(t, doc)
		test.AssertNotNil(t, err)
	})

	t.Run("EmptyProjectOnly", func(t *testing.T) {
		yaml := `project:`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)
	})

	t.Run("ProjectWithName", func(t *testing.T) {
		yaml := `
project:
  name: foo
`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)

		test.AssertNotNil(t, doc.Project)
		test.Assert(t, doc.Project.Name == "foo")
	})

	t.Run("ProjectWithAllMembers", func(t *testing.T) {
		yaml := `
project:
  name: foo
  dependencies: 
    - path: bar
    - project: baz
  steps:
    - cmd: foo
    - cmd: bar
    - cmd: baz
`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)

		test.AssertNotNil(t, doc.Project)
		test.Assert(t, doc.Project.Name == "foo")

		// TODO: assert items of both slices
		test.AssertNotNil(t, doc.Project.Dependencies)
		test.Assert(t, len(doc.Project.Dependencies) > 0)
		test.AssertNotNil(t, doc.Project.Steps)
		test.Assert(t, len(doc.Project.Steps) > 0)
	})

	t.Run("EmptyWorkspaceOnly", func(t *testing.T) {
		yaml := `workspace:`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)
	})

	t.Run("WorkspaceWithName", func(t *testing.T) {
		yaml := `
workspace:
  name: foo
`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)

		test.AssertNotNil(t, doc.Workspace)
		test.Assert(t, doc.Workspace.Name == "foo")
	})

	t.Run("WorkspaceWithAllMembers", func(t *testing.T) {
		yaml := `
workspace:
  name: foo
  includes:
  - for
  - bar
  - baz
`

		reader := strings.NewReader(yaml)
		doc, err := document_definition.Read(reader)

		test.AssertNotNil(t, doc)
		test.AssertNil(t, err)

		test.AssertNotNil(t, doc.Workspace)
		test.Assert(t, doc.Workspace.Name == "foo")

		// TODO: assert items of slice
		test.AssertNotNil(t, doc.Workspace.Includes)
		test.Assert(t, len(doc.Workspace.Includes) > 0)

	})
}
