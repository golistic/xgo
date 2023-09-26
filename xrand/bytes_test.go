/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xrand_test

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/golistic/xgo/xrand"
	"github.com/golistic/xgo/xt"
)

func TestBytes(t *testing.T) {
	t.Run("some uniqueness", func(t *testing.T) {
		s := map[string]bool{}
		for i := 0; i < 100000; i++ {
			r, err := xrand.RandomBytes(16)
			xt.OK(t, err)
			henc := hex.EncodeToString(r)
			xt.Assert(t, !s[henc], "expected at least some uniqueness")
			s[henc] = true
		}
	})

	for _, n := range []int{16, 8, 33} {
		t.Run(fmt.Sprintf("length %d", n), func(t *testing.T) {
			r, err := xrand.RandomBytes(n)
			xt.OK(t, err)
			xt.Eq(t, n, len(r))
		})
	}

	t.Run("panics if n < 1", func(t *testing.T) {
		xt.Panics(t, func() {
			_, _ = xrand.RandomBytes(0)
		})

		xt.Panics(t, func() {
			_, _ = xrand.RandomBytes(-20)
		})
	})
}
