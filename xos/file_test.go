// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestIsRegularFile(t *testing.T) {
	t.Run("existing regular file", func(t *testing.T) {
		xt.Assert(t, IsRegularFile("file.go"))
	})

	t.Run("non-existing regular file", func(t *testing.T) {
		xt.Assert(t, !IsRegularFile("filefilefile.go"))
	})

	t.Run("dir is not a regular file", func(t *testing.T) {
		xt.Assert(t, !IsRegularFile("_testdata"))
	})
}
