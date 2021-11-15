package validation

import (
	"errors"
	"fmt"
	"regexp"
)

func newNameRegexp() *regexp.Regexp {
	r, err := regexp.Compile(nameRegexPattern)

	if err != nil {
		panic(err)
	}

	return r
}

var nameRegexp = newNameRegexp()

func ValidateManifoldName(name string) error {
	if !nameRegexp.MatchString(name) {
		return errors.New(fmt.Sprintf(InvalidManifoldName, name, nameRegexp))
	}

	return nil
}
