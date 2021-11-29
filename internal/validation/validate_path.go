package validation

import (
	"manifold/internal/errors"
	"manifold/internal/utils"
)

func ValidatePath(path string) error {
	if path == "" {
		return errors.NewError(EmptyPath)
	}

	if !utils.Exists(path) {
		return errors.NewError(InvalidPath, path)
	}

	return nil
}
