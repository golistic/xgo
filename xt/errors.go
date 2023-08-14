// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"testing"
)

// OK checks whether err is nil.
// This could be replaced with using Eq, but since checking the error in Go
// is done lots, it is nicer to read in tests.
func OK(t *testing.T, err error, messages ...string) {
	TestHelper(t)

	ok(t, nil, err, messages...)
}

func ok(t *testing.T, out io.Writer, err error, messages ...string) {
	if err != nil {
		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}

		fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected no error, got:\u001B[0m\n%s", err.Error()), messages...)
	}
}

// KO checks whether err is not nil.
//
// Reverse of OK, and also, well, you know 🥊.
//
// This could be replaced with using Eq, but since checking the error in Go
// is done lots, it is nicer to read in tests.
func KO(t *testing.T, err error, messages ...string) {
	TestHelper(t)
	ko(t, nil, err, messages...)
}

func ko(t *testing.T, out io.Writer, err error, messages ...string) {
	if err == nil {
		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}
		fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected error\u001B[0m"), messages...)
	}
}
