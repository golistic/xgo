// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestInDockerContainer(t *testing.T) {
	// this works differently when in container or not
	exp := IsRegularFile("/.dockerenv")
	xt.Eq(t, exp, InDockerContainer())
}
