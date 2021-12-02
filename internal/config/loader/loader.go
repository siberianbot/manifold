package loader

import (
	"manifold/internal/config"
	"manifold/internal/config/utils"
	"manifold/internal/config/validation"
	"os"
)

type Interface interface {
	FromPath(path string) (*config.Configuration, error)
}

type loader struct {
	//
}

func NewLoader() Interface {
	return &loader{}
}

func (l *loader) FromPath(path string) (*config.Configuration, error) {
	path, pathErr := utils.ResolvePath(path)

	if pathErr != nil {
		return nil, pathErr
	}

	file, fileErr := os.Open(path)
	defer func() { _ = file.Close() }()

	if fileErr != nil {
		return nil, fileErr
	}

	cfg, cfgErr := config.Read(file)

	if cfgErr != nil {
		return nil, cfgErr
	}

	validationErr := validation.Validate(cfg, path)

	if validationErr != nil {
		return nil, validationErr
	}

	return cfg, nil
}
