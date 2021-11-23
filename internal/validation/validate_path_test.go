package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPathValidation(t *testing.T) {
	t.Run("EmptyPath", func(t *testing.T) {
		path := ""
		expected := EmptyPath

		err := ValidatePath(path)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	t.Run("ValidPath", func(t *testing.T) {
		filename := "bar"
		path := filepath.Join(t.TempDir(), filename)

		file, _ := os.Create(path)
		_ = file.Close()

		err := ValidatePath(path)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}
	})

	t.Run("ValidDirPath", func(t *testing.T) {
		path := t.TempDir()

		err := ValidatePath(path)

		if err != nil {
			t.Errorf("error is %v, not nil", err)
		}
	})

	t.Run("InvalidPath", func(t *testing.T) {
		path := "baz"
		expected := fmt.Sprintf(InvalidPath, path)

		err := ValidatePath(path)

		if err == nil {
			t.Error("error is nil")
		} else if err.Error() != expected {
			t.Errorf("error is %v, not %s", err, expected)
		}
	})

	if runtime.GOOS == "windows" {
		t.Run("InvalidSymbolsInPath", func(t *testing.T) {
			path := "baz\":<>|?*"
			expected := fmt.Sprintf(InvalidPath, path)

			err := ValidatePath(path)

			if err == nil {
				t.Error("error is nil")
			} else if err.Error() != expected {
				t.Errorf("error is %v, not %s", err, expected)
			}
		})
	}
}
