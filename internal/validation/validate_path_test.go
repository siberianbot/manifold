package validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestPathValidation(t *testing.T) {
	t.Run("EmptyPath", func(t *testing.T) {
		path := ""
		err := ValidatePath(path)

		assert.EqualError(t, err, EmptyPath)
	})

	t.Run("ValidPath", func(t *testing.T) {
		filename := "bar"
		path := filepath.Join(t.TempDir(), filename)

		file, _ := os.Create(path)
		_ = file.Close()

		err := ValidatePath(path)

		assert.NoError(t, err)
	})

	t.Run("ValidDirPath", func(t *testing.T) {
		path := t.TempDir()

		err := ValidatePath(path)

		assert.NoError(t, err)
	})

	t.Run("InvalidPath", func(t *testing.T) {
		path := "baz"
		err := ValidatePath(path)

		assert.EqualError(t, err, fmt.Sprintf(InvalidPath, path))
	})

	if runtime.GOOS == "windows" {
		t.Run("InvalidSymbolsInPath", func(t *testing.T) {
			path := "baz\":<>|?*"
			err := ValidatePath(path)

			assert.EqualError(t, err, fmt.Sprintf(InvalidPath, path))
		})
	}
}
