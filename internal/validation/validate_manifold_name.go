package validation

import (
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
		return NewError(EmptyManifoldName)
	}

	if !nameRegexp.MatchString(name) {
		return NewError(InvalidManifoldName, name, nameRegexp)
	}

	return nil
}
