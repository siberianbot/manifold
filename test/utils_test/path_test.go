package utils_test

import (
	"manifold/internal/utils"
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestBuildPath(t *testing.T) {
	t.Run("NoBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := ""
		result := utils.BuildPath(baseDir, relativePath)

		test.Assert(t, result == "")
	})

	t.Run("OnlyBaseDir", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := utils.BuildPath(baseDir, relativePath)

		test.Assert(t, result == baseDir)
	})

	t.Run("OnlyRelativePath", func(t *testing.T) {
		baseDir := ""
		relativePath := "RelativePath"
		result := utils.BuildPath(baseDir, relativePath)

		test.Assert(t, result == relativePath)
	})

	t.Run("WithBaseDirAndRelativePath", func(t *testing.T) {
		baseDir := "BaseDir"
		relativePath := ""
		result := utils.BuildPath(baseDir, relativePath)

		test.Assert(t, result == filepath.Join(baseDir, relativePath))
	})
}
