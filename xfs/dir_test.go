/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestIsDir(t *testing.T) {

	tempDir, err := os.MkdirTemp("", "testdir")
	xt.OK(t, err, "create temp dir")
	defer func() { _ = os.RemoveAll(tempDir) }()

	tempFile, err := os.CreateTemp(tempDir, "testfile")
	xt.OK(t, err, "create file")
	_ = tempFile.Close()

	xt.OK(t, os.Mkdir(filepath.Join(tempDir, "existingPath"), 0750), "create directory")
	_ = tempFile.Close()

	// Define the tests.
	tests := []struct {
		name          string
		path          string
		expectedIsDir bool
	}{
		{name: "directory", path: "existingPath", expectedIsDir: true},
		{name: "file", path: tempFile.Name(), expectedIsDir: false},
		{name: "non existing file", path: "nonExistingPath", expectedIsDir: false},
	}

	fsys := os.DirFS(tempDir)
	xt.Assert(t, fsys != nil, "fsys is nil")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(fsys, tt.path); got != tt.expectedIsDir {
				t.Errorf("IsDir() = %v, want %v", got, tt.expectedIsDir)
			}
		})
	}
}
