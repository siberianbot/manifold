package validations

import (
	"fmt"
	"manifold/internal/validation"
	"manifold/test"
	"os"
	path2 "path"
	"runtime"
	"testing"
)

func TestPathValidation(t *testing.T) {
	t.Run("EmptyPath", func(t *testing.T) {
		path := ""
		expected := validation.EmptyPath

		err := validation.ValidatePath(path)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	t.Run("ValidPath", func(t *testing.T) {
		filename := "bar"
		path := path2.Join(t.TempDir(), filename)

		file, err := os.Create(path)

		if err != nil {
			t.Fatal(err)
			return
		} else {
			_ = file.Close()

			defer func() { _ = os.Remove(path) }()
		}

		err = validation.ValidatePath(path)

		test.Assert(t, err == nil)
	})

	t.Run("ValidDirPath", func(t *testing.T) {
		path := t.TempDir()

		err := validation.ValidatePath(path)

		test.Assert(t, err == nil)
	})

	t.Run("InvalidPath", func(t *testing.T) {
		path := "baz"
		expected := generateInvalidPathMsg(path)

		err := validation.ValidatePath(path)

		test.Assert(t, err != nil)
		test.Assert(t, err.Error() == expected)
	})

	if runtime.GOOS == "windows" {
		t.Run("InvalidSymbolsInPath", func(t *testing.T) {
			path := "baz\":<>|?*"
			expected := generateInvalidPathMsg(path)

			err := validation.ValidatePath(path)

			test.Assert(t, err != nil)
			test.Assert(t, err.Error() == expected)
		})
	}
}

func generateInvalidPathMsg(path string) string {
	return fmt.Sprintf(validation.InvalidPath, path)
}
