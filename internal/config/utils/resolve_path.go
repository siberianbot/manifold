package utils

import (
	"manifold/internal/errors"
	"manifold/internal/utils"
	"os"
	"path/filepath"
)

const (
	YmlFilename  = ".manifold.yml"
	YamlFilename = ".manifold.yaml"
)

const (
	NotManifoldPath = "path \"%s\" does not contain manifold configuration"
)

func ResolvePath(path string) (string, error) {
	absPath, absErr := filepath.Abs(path)

	if absErr != nil {
		return "", absErr
	}

	stat, statErr := os.Stat(path)

	if statErr != nil {
		return "", statErr
	}

	if !stat.IsDir() {
		return absPath, nil
	}

	if ymlPath := filepath.Join(absPath, YmlFilename); utils.Exists(ymlPath) {
		return ymlPath, nil
	}

	if yamlPath := filepath.Join(absPath, YamlFilename); utils.Exists(yamlPath) {
		return yamlPath, nil
	}

	return "", errors.NewError(NotManifoldPath, absPath)
}
