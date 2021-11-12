package test

import (
	"runtime/debug"
	"testing"
)

func Assert(t *testing.T, condition bool) {
	if !condition {
		t.Error(string(debug.Stack()))
	}
}
