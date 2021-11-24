package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestResolvePath(t *testing.T) {
	t.Run("PathIsEmptyDir", func(t *testing.T) {
		path := t.TempDir()

		result, err := ResolvePath(path)

		assert.Empty(t, result)
		assert.EqualError(t, err, fmt.Sprintf(NotManifoldPath, path))
	})

	t.Run("PathIsInvalid", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "foo")

		result, err := ResolvePath(path)

		assert.Empty(t, result)
		assert.True(t, os.IsNotExist(err))
	})

	t.Run("PathIsFile", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, "foo")

		file, _ := os.Create(path)
		_ = file.Close()

		result, err := ResolvePath(path)

		assert.Equal(t, path, result)
		assert.NoError(t, err)
	})

	t.Run("PathIsYml", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, YmlFilename)

		file, _ := os.Create(path)
		_ = file.Close()

		result, err := ResolvePath(dir)

		assert.Equal(t, path, result)
		assert.NoError(t, err)
	})

	t.Run("PathIsYaml", func(t *testing.T) {
		dir := t.TempDir()
		path := filepath.Join(dir, YamlFilename)

		file, _ := os.Create(path)
		_ = file.Close()

		result, err := ResolvePath(dir)

		assert.Equal(t, path, result)
		assert.NoError(t, err)
	})
}
