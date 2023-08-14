// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"path/filepath"
	"runtime"
	"testing"
)

// Panics checks if function f panics.
func Panics(t *testing.T, f func()) {
	TestHelper(t)
	panics(t, nil, "", f)
}

// PanicsEq checks if function f panics and whether the panic message is equal to exp.
func PanicsEq(t *testing.T, exp string, f func()) {
	TestHelper(t)
	panics(t, nil, exp, f)
}

func panics(t *testing.T, out io.Writer, exp string, f func()) {
	_, file, lineNr, _ := runtime.Caller(1)

	defer func() {
		r := recover()
		if r == nil {
			fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected panic\u001B[0m (in test %s:%d)", filepath.Base(file), lineNr))
		}
		if exp != "" {
			Eq(t, exp, fmt.Sprintf("%v", r))
		}
	}()

	f()
}
