/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xrand_test

import (
	"fmt"
	"testing"

	"github.com/golistic/xgo/xrand"
	"github.com/golistic/xgo/xt"
)

func TestRandomAlphaNumeric(t *testing.T) {
	t.Run("some uniqueness", func(t *testing.T) {
		s := map[string]bool{}
		for i := 0; i < 100000; i++ {
			r := xrand.AlphaNumeric(16)
			xt.Assert(t, !s[r], "expected at least some uniqueness")
			s[r] = true
		}
	})

	for _, n := range []int{16, 8, 33} {
		t.Run(fmt.Sprintf("length %d", n), func(t *testing.T) {
			r := xrand.AlphaNumeric(n)
			xt.Eq(t, n, len(r))
		})
	}

	t.Run("panics if n < 1", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = xrand.AlphaNumeric(0)
		})

		xt.Panics(t, func() {
			_ = xrand.AlphaNumeric(-20)
		})
	})
}
