package utils

import (
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestBuildPath(t *testing.T) {
	t.Run("NoBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		test.Assert(t, result == "")
	})

	t.Run("OnlyBaseDir", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		test.Assert(t, result == baseDir)
	})

	t.Run("OnlyRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := "RelativePath"
		result := BuildPath(baseDir, relativePath)

		test.Assert(t, result == relativePath)
	})

	t.Run("WithBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := BuildPath(baseDir, relativePath)

		test.Assert(t, result == filepath.Join(baseDir, relativePath))
	})
}
