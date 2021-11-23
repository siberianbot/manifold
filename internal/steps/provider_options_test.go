package steps

import "testing"

func TestProviderOptions(t *testing.T) {
	options := NewProviderOptions()

	if options.Factories == nil {
		t.Error("options.Factories is nil")
	} else if len(options.Factories) != 0 {
		t.Error("options.Factories is not empty")
	}

	if options.Executors == nil {
		t.Error("options.Executors is nil")
	} else if len(options.Factories) != 0 {
		t.Error("options.Executors is not empty")
	}
}
