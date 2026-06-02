/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xptr_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/golistic/xgo/xt"
)

func TestOf(t *testing.T) {
	var cases = []any{
		"String",
		1234,
		int64(-1234),
		uint16(123),
		float32(12.34),
		12.34,
		true,
		'⌘',
		byte(20),
		time.Now(),
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%T", c), func(t *testing.T) {
			got := new(c)
			xt.Eq(t, reflect.Pointer, reflect.ValueOf(got).Kind())
			xt.Assert(t, c == *new(c))
		})
	}
}
