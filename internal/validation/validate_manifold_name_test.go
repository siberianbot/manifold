package validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestManifoldNameValidation(t *testing.T) {
	validTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			err := ValidateManifoldName(name)

			assert.NoError(t, err)
		})
	}

	invalidTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			err := ValidateManifoldName(name)

			assert.EqualError(t, err, fmt.Sprintf(InvalidManifoldName, name, NameRegexPattern))
		})
	}

	t.Run("EmptyName", func(t *testing.T) {
		name := ""
		err := ValidateManifoldName(name)

		assert.EqualError(t, err, EmptyManifoldName)
	})

	t.Run("ValidNames", func(t *testing.T) {
		validTest(t, "ValidName1")
		validTest(t, "1234567890")
		validTest(t, "A")
		validTest(t, "Valid-Name.With_Underscore1111")
	})

	t.Run("InvalidNames", func(t *testing.T) {
		invalidTest(t, "!!!!")
		invalidTest(t, "mail@somewhe.re")
		invalidTest(t, "Th1s!s!nc)RR3ctN$m3!!")
	})
}
