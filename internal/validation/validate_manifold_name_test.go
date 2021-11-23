package validation

import (
	"fmt"
	"testing"
)

func TestManifoldNameValidation(t *testing.T) {
	validTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			err := ValidateManifoldName(name)

			if err != nil {
				t.Errorf("error is %v, not nil", err)
			}
		})
	}

	invalidTest := func(t *testing.T, name string) {
		t.Run(name, func(t *testing.T) {
			expected := fmt.Sprintf(InvalidManifoldName, name, NameRegexPattern)

			err := ValidateManifoldName(name)

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != expected {
				t.Errorf("error is %v, not %s", err, expected)
			}
		})
	}

	t.Run("EmptyName", func(t *testing.T) {
		name := ""
		expected := EmptyManifoldName

		err := ValidateManifoldName(name)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
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
