package validation

import (
	"manifold/internal/validation"
	"manifold/test"
	"path/filepath"
	"testing"
)

func TestContextFactory(t *testing.T) {
	t.Run("InstantiationCheck", func(t *testing.T) {
		factory := validation.NewDefaultContextFactory()

		test.Assert(t, factory != nil)
	})

	t.Run("NewContext", func(t *testing.T) {
		factory := validation.NewDefaultContextFactory()
		dir := "foo"
		path := filepath.Join(dir, "bar")

		ctx := factory.New(path)

		test.Assert(t, ctx != nil)
		test.Assert(t, ctx.Dir() == dir)
	})
}
