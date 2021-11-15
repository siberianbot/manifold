package validation

import (
	"errors"
	"fmt"
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
		return errors.New(EmptyManifoldName)
	}

	if !nameRegexp.MatchString(name) {
		return errors.New(fmt.Sprintf(InvalidManifoldName, name, nameRegexp))
	}

	return nil
}
