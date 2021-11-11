package test

import (
	"reflect"
	"testing"
)

func AssertNil(t *testing.T, v interface{}) {
	if v == nil {
		return
	}

	if !reflect.ValueOf(v).IsNil() {
		t.Error()
	}
}

//AssertNotNil is obsolete
func AssertNotNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Error()
	}

	if reflect.ValueOf(v).IsNil() {
		t.Error()
	}
}

func Assert(t *testing.T, condition bool) {
	if !condition {
		t.Error()
	}
}
