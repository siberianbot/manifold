package validation

import (
	"manifold/internal/errors"
	"regexp"
)

func newNameRegexp() *regexp.Regexp {
	r, err := regexp.Compile(NameRegexPattern)

	if err != nil {
		panic(err)
	}

	return r
}

var nameRegexp = newNameRegexp()

func ValidateManifoldName(name string) error {
	if name == "" {
		return errors.NewError(EmptyManifoldName)
	}

	if !nameRegexp.MatchString(name) {
		return errors.NewError(InvalidManifoldName, name, nameRegexp)
	}

	return nil
}
