package utils

import "path/filepath"

func BuildPath(baseDir string, relativePath string) string {
	return filepath.Clean(filepath.Join(baseDir, relativePath))
}
