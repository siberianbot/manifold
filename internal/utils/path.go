package utils

import (
	"os"
	"path/filepath"
)

func BuildPath(baseDir string, relativePath string) string {
	if baseDir == "" {
		return relativePath
	}

	if relativePath == "" {
		return baseDir
	}

	return filepath.Clean(filepath.Join(baseDir, relativePath))
}

func Exists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}
