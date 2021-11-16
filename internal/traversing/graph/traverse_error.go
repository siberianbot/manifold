package graph

import (
	"errors"
	"fmt"
)

type TraverseError struct {
	path string
	err  error
}

func (err TraverseError) Error() error {
	return errors.New(fmt.Sprintf("%s: %v", err.path, err.err))
}

func newTraverseError(path string, err error) (traverseErr *TraverseError) {
	traverseErr = new(TraverseError)
	traverseErr.path = path
	traverseErr.err = err

	return traverseErr
}
