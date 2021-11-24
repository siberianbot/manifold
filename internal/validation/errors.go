package validation

import (
	"errors"
	"fmt"
)

const (
	EmptyManifoldName   = "name is empty"
	InvalidManifoldName = "name \"%s\" does not matches regex pattern \"%s\""
	EmptyPath           = "path is empty"
	InvalidPath         = "invalid path \"%s\""

	UnknownTarget = "target \"%s\" not found"
)

func NewError(msg string, args ...interface{}) error {
	if len(args) > 0 {
		return errors.New(fmt.Sprintf(msg, args...))
	} else {
		return errors.New(msg)
	}
}
