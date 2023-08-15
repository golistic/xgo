// Copyright (c) 2022, Geert JM Vanderkelen

package xconv

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestParseBool(t *testing.T) {
	t.Run("returning true", func(t *testing.T) {
		var cases = []string{"1", "t", "T", "trUe", "TRUE", "yes", "y", "Ok"}
		for _, c := range cases {
			t.Run(c, func(t *testing.T) {
				r, err := ParseBool(c)
				xt.OK(t, err)
				xt.Assert(t, r)
			})
		}
	})

	t.Run("returning false", func(t *testing.T) {
		var cases = []string{"0", "falSe", "F", "no", "n"}
		for _, c := range cases {
			t.Run(c, func(t *testing.T) {
				r, err := ParseBool(c)
				xt.OK(t, err)
				xt.Assert(t, !r)
			})
		}
	})

	t.Run("returning false and error", func(t *testing.T) {
		var cases = []string{"5", "untrue", "foo"}
		for _, c := range cases {
			t.Run(c, func(t *testing.T) {
				r, err := ParseBool(c)
				xt.KO(t, err)
				xt.Assert(t, !r)
			})
		}
	})
}
