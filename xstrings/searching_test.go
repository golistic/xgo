// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestSearch(t *testing.T) {
	mascots := []string{"Go gopher", "Sakila", "Wilber", "Tux"}

	var cases = []struct {
		x   string
		exp int
	}{
		{
			x:   "Go gopher",
			exp: 0,
		},
		{
			x:   "Sakila",
			exp: 1,
		},
		{
			x:   "Tux",
			exp: 3,
		},
		{
			x:   "Duke",
			exp: -1,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			xt.Eq(t, c.exp, Search(mascots, c.x))
		})
	}
}

func TestSliceHas(t *testing.T) {
	mascots := []string{"Go gopher", "Sakila", "Wilber", "Tux"}

	var cases = []struct {
		x   string
		exp bool
	}{
		{
			x:   "Go gopher",
			exp: true,
		},
		{
			x:   "Sakila",
			exp: true,
		},
		{
			x:   "Tux",
			exp: true,
		},
		{
			x:   "Duke",
			exp: false,
		},
	}
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			xt.Eq(t, c.exp, SliceHas(mascots, c.x))
		})
	}
}
