package errors

import (
	"errors"
	"fmt"
)

func NewError(msg string, args ...interface{}) error {
	if len(args) > 0 {
		return errors.New(fmt.Sprintf(msg, args...))
	} else {
		return errors.New(msg)
	}
}
