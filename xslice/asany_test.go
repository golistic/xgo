/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xslice_test

import (
	"testing"

	"github.com/golistic/xgo/xslice"
	"github.com/golistic/xgo/xt"
)

func TestAsAny(t *testing.T) {
	t.Run("slice of strings", func(t *testing.T) {
		have := []string{"foo", "bar"}
		exp := []any{"foo", "bar"}

		xt.Eq(t, exp, xslice.AsAny(have))
	})

	t.Run("slice of int", func(t *testing.T) {
		have := []int{2, 1, -1, 90}
		exp := []any{2, 1, -1, 90}

		xt.Eq(t, exp, xslice.AsAny(have))
	})

	t.Run("slice of []byte", func(t *testing.T) {
		foo := []byte{'f', 'o', 'o'}
		bar := []byte{'b', 'a', 'r'}
		have := [][]byte{foo, bar}
		exp := []any{foo, bar}

		xt.Eq(t, exp, xslice.AsAny(have))
	})
}
