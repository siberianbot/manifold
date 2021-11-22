package validation

import (
	"manifold/internal/utils"
)

func ValidatePath(path string) error {
	if path == "" {
		return NewError(EmptyPath)
	}

	if !utils.Exists(path) {
		return NewError(InvalidPath, path)
	}

	return nil
}
