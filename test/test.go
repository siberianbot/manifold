package test

import (
	"os"
	"path/filepath"
	"runtime/debug"
	"testing"
)

func Assert(t *testing.T, condition bool) {
	if !condition {
		t.Fatal(string(debug.Stack()))
	}
}

func CreateDir(t *testing.T, dir string) {
	err := os.MkdirAll(dir, os.ModeDir)

	if err != nil {
		t.Fatalf("failed to create %s: %v", dir, err)
	}
}

func CreateFile(t *testing.T, path string, content string) {
	CreateDir(t, filepath.Dir(path))

	file, err := os.Create(path)

	if err != nil {
		t.Fatalf("failed to create %s: %v", path, err)
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	_, err = file.WriteString(content)

	if err != nil {
		t.Fatalf("failed to write content into %s: %v", path, err)
	}
}
