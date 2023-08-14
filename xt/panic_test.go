// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"bytes"
	"strings"
	"testing"
)

func TestPanics(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		Panics(t, func() {
			panic("Don't panic!")
		})
	})

	t.Run("panic line matches", func(t *testing.T) {
		PanicsEq(t, "this is the panic line", func() {
			panic("this is the panic line")
		})
	})

	t.Run("no panic", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		panics(t, have, "", func() {
			return
		})
		exp := []byte("\u001B[31;1mexpected panic\u001B[0m (in test panic_test.go:26)")
		if !strings.Contains(have.String(), "expected panic") {
			t.Fatal("expected:", string(exp))
		}
	})
}
