/*
 * Copyright (c) 2025, Geert JM Vanderkelen
 */
package xjson

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestFormat(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{
			name:  "EmptyMap",
			input: map[string]string{},
			want:  "{}",
		},
		{
			name: "SimpleMap",
			input: map[string]string{
				"key": "value",
			},
			want: `{
  "key": "value"
}`,
		},
		{
			name: "NestedMap",
			input: map[string]any{
				"key": map[string]any{
					"nestedKey": "nestedValue",
				},
			},
			want: `{
  "key": {
    "nestedKey": "nestedValue"
  }
}`,
		},
		{
			name: "Array",
			input: []any{
				1, "two", 3.14,
			},
			want: `[
  1,
  "two",
  3.14
]`,
		},
		{
			name:  "InvalidValue",
			input: make(chan int),
			want:  "JSON encoding failed (marshal): json: 2unsupported type: chan int",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			xt.Eq(t, tc.want, Format(tc.input))
		})
	}
}
