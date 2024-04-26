/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xptr

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestValueOr(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		tests := []struct {
			name     string
			a        *int
			b        int
			expected int
		}{
			{
				name:     "Nil pointer",
				a:        nil,
				b:        3,
				expected: 3,
			},
			{
				name:     "Non-nil pointer",
				a:        Of(2),
				b:        3,
				expected: 2,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := ValueOr(test.a, test.b)
				xt.Eq(t, test.expected, result)
			})
		}
	})

	t.Run("bool", func(t *testing.T) {
		tests := []struct {
			name     string
			a        *bool
			b        bool
			expected bool
		}{
			{
				name:     "Nil pointer",
				a:        nil,
				b:        true,
				expected: true,
			},
			{
				name:     "Non-nil pointer",
				a:        Of(true),
				b:        false,
				expected: true,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				result := ValueOr(test.a, test.b)
				xt.Eq(t, test.expected, result)
			})
		}
	})
}
