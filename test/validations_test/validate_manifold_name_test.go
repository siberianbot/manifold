package validations_test

import (
	"fmt"
	"manifold/internal/validation"
	"manifold/test"
	"testing"
)

func TestManifoldNameValidation(t *testing.T) {
	validTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			err := validation.ValidateManifoldName(name)
			test.Assert(t, err == nil)
		})
	}

	invalidTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			expectedErrMsg := generateInvalidManifoldNameMsg(name)

			err := validation.ValidateManifoldName(name)

			test.Assert(t, err != nil)
			test.Assert(t, err.Error() == expectedErrMsg)
		})
	}

	t.Run("EmptyName", func(t *testing.T) {
		name := ""
		expectedErrMsg := validation.EmptyManifoldName

		err := validation.ValidateManifoldName(name)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expectedErrMsg)
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

func generateInvalidManifoldNameMsg(name string) string {
	return fmt.Sprintf(validation.InvalidManifoldName, name, validation.NameRegexPattern)
}
