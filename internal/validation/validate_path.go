package validation

import (
	"errors"
	"fmt"
	"os"
)

func ValidatePath(path string) error {
	if path == "" {
		return errors.New(EmptyPath)
	}

	_, err := os.Stat(path)

	if err != nil {
		return errors.New(fmt.Sprintf(InvalidPath, path))
	}

	return nil
}
