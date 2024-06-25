/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xslice

import (
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
