package graph

import (
	"os"
	"path/filepath"
)

func processPath(path string) (absPath string, absDir string, err error) {
	const defaultName = ".manifold.yml"

	stat, err := os.Stat(path)

	if err != nil {
		return "", "", err
	}

	if stat.IsDir() {
		path = filepath.Join(path, defaultName)
	}

	path, err = filepath.Abs(path)

	if err != nil {
		return "", "", err
	}

	return path, filepath.Dir(path), nil
}
