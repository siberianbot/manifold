package errors

import (
	"fmt"
	"strings"
)

type AggregateError struct {
	msg    string
	errors []error
}

func NewAggregateError(msg string, errors ...error) error {
	return &AggregateError{msg: msg, errors: errors}
}

func (a *AggregateError) Error() string {
	errorMsgs := make([]string, len(a.errors))

	for idx, err := range a.errors {
		errorMsgs[idx] = err.Error()
	}

	return fmt.Sprintf("%s\n%s", a.msg, strings.Join(errorMsgs, "\n"))
}
