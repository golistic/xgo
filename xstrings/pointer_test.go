// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import (
	"reflect"
	"testing"

	"github.com/golistic/xgo/xt"
)

type stringTyped string

func TestPointer(t *testing.T) {
	t.Run("string as pointer", func(t *testing.T) {
		s := "a string"
		have := Pointer(s)
		rv := reflect.ValueOf(have)
		xt.Eq(t, reflect.Pointer, rv.Kind())
	})

	t.Run("string type as pointer", func(t *testing.T) {
		var s stringTyped = "a string"
		have := Pointer(s)
		rv := reflect.ValueOf(have)
		xt.Eq(t, reflect.Pointer, rv.Kind())
	})
}
