package loader

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
)

func TestFileNotExists(t *testing.T) {
	path := filepath.Join(t.TempDir(), "invalid.yml")
	loader := NewLoader()

	cfg, err := loader.FromPath(path)

	assert.Nil(t, cfg)
	assert.Error(t, err)
}

func TestFileExistsButInvalid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "invalid.yml")
	loader := NewLoader()

	data := make([]byte, 1024)
	rand.Read(data)
	file, _ := os.Create(path)
	_, _ = file.Write(data)
	_ = file.Close()

	cfg, err := loader.FromPath(path)

	assert.Nil(t, cfg)
	assert.Error(t, err)
}

func TestFileExistsAndValid(t *testing.T) {
	path := filepath.Join(t.TempDir(), "valid.yml")
	loader := NewLoader()

	file, _ := os.Create(path)
	_, _ = file.WriteString(`
project:
  name: foo
`)
	_ = file.Close()

	cfg, err := loader.FromPath(path)

	assert.NoError(t, err)
	assert.NotEmpty(t, cfg)
	assert.Equal(t, "foo", cfg.Project.Name)
}
