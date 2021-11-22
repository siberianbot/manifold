package validation

import (
	"errors"
	"fmt"
	"manifold/internal/utils"
)

func ValidatePath(path string) error {
	if path == "" {
		return errors.New(EmptyPath)
	}

	if !utils.Exists(path) {
		return errors.New(fmt.Sprintf(InvalidPath, path))
	}

	return nil
}
