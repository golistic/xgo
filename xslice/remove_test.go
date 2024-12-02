/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xslice

import (
	"slices"
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestRemoveFirst(t *testing.T) {
	tests := []struct {
		name string
		have []int
		want []int
	}{
		{
			name: "empty slice",
			have: []int{},
			want: []int{},
		},
		{
			name: "single element",
			have: []int{1},
			want: []int{},
		},
		{
			name: "multiple elements",
			have: []int{1, 2, 3, 4},
			want: []int{2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xt.Eq(t, tt.want, RemoveFirst(tt.have))
		})
	}
}

func TestRemoveFirstN(t *testing.T) {

	type test struct {
		name string
		have []int
		n    int
		want []int
	}

	tests := []test{
		{
			name: "empty slice",
			have: []int{},
			n:    1,
			want: []int{},
		},
		{
			name: "remove more than length of slice",
			have: []int{1, 2, 3},
			n:    4,
			want: []int{},
		},
		{
			name: "remove zero",
			have: []int{1, 2, 3},
			n:    0,
			want: []int{1, 2, 3},
		},
		{
			name: "remove some",
			have: []int{1, 2, 3, 4, 5},
			n:    3,
			want: []int{4, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xt.Eq(t, tt.want, RemoveFirstN(tt.have, tt.n))
		})
	}
}

func TestExclude(t *testing.T) {
	fInt := func(s []int, toExclude []int, want []int) {
		var have []int
		Exclude(s, toExclude)(func(i int) bool {
			have = append(have, i)
			return true
		})
		xt.Assert(t, slices.Equal(want, have))
	}

	fString := func(s []string, toExclude []string, want []string) {
		var have []string

		for el := range Exclude(s, toExclude) {
			have = append(have, el)
		}

		xt.Assert(t, slices.Equal(want, have))
	}

	// empty slice
	fInt([]int{}, []int{1, 2}, []int{})
	fString([]string{}, []string{"a", "b"}, []string{})

	// empty toExclude
	fInt([]int{1, 2, 3}, []int{}, []int{1, 2, 3})
	fString([]string{"a", "b", "c"}, []string{}, []string{"a", "b", "c"})

	// exclude some
	fInt([]int{1, 2, 3, 4, 5}, []int{2, 4}, []int{1, 3, 5})
	fString([]string{"a", "b", "c", "d", "e"}, []string{"b", "d"}, []string{"a", "c", "e"})

	// exclude all
	fInt([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}, []int{})
	fString([]string{"a", "b", "c", "d", "e"}, []string{"a", "b", "c", "d", "e"}, []string{})
}
