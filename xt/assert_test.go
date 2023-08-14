// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"math/rand"
	"testing"
)

// TestAssert is kept to a minimum since it uses Eq.
func TestAssert(t *testing.T) {
	t.Run("true", func(t *testing.T) {
		Assert(t, rand.Int() > -1)
	})
}
