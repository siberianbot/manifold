package utils

import (
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"testing"
)

func TestBuildPath(t *testing.T) {
	t.Run("NoBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		assert.Equal(t, "", result)
	})

	t.Run("OnlyBaseDir", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		assert.Equal(t, baseDir, result)
	})

	t.Run("OnlyRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := "RelativePath"
		result := BuildPath(baseDir, relativePath)

		assert.Equal(t, relativePath, result)
	})

	t.Run("WithBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		assert.Equal(t, filepath.Join(baseDir, relativePath), result)
	})
}
