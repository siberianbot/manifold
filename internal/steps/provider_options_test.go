package steps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProviderOptions(t *testing.T) {
	options := NewProviderOptions()

	assert.NotEmpty(t, options)
	assert.NotNil(t, options.Factories)
	assert.Empty(t, options.Factories)
	assert.NotNil(t, options.Executors)
	assert.Empty(t, options.Executors)
}
