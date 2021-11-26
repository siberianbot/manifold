package building

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestNodeBuilder(t *testing.T) {
	t.Run("NewNodeBuilder", func(t *testing.T) {
		nodeBuilder := NewNodeBuilder(nil)

		assert.NotNil(t, nodeBuilder)
	})

	t.Run("FromPath", testNodeBuilderFromPath)
}

func testNodeBuilderFromPath(t *testing.T) {
	t.Run("FileNotExists", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "invalid.yml")
		nodeBuilder := NewNodeBuilder(nil)

		node, err := nodeBuilder.FromPath(path)

		assert.Empty(t, node)
		assert.Error(t, err)
	})

	t.Run("FileExistsButInvalid", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "invalid.yml")
		nodeBuilder := NewNodeBuilder(nil)

		data := make([]byte, 1024)
		rand.Read(data)
		file, _ := os.Create(path)
		_, _ = file.Write(data)
		_ = file.Close()

		node, err := nodeBuilder.FromPath(path)

		assert.Empty(t, node)
		assert.Error(t, err)
	})

	t.Run("FileExistsAndValid", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "valid.yml")
		nodeBuilder := NewNodeBuilder(nil)

		file, _ := os.Create(path)
		_, _ = file.WriteString(`
project:
  name: foo
`)
		_ = file.Close()

		node, err := nodeBuilder.FromPath(path)

		assert.NoError(t, err)
		assert.NotEmpty(t, node)
		assert.Equal(t, "foo", node.Name())
		assert.Equal(t, path, node.Path())
	})
}
