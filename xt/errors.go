// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"errors"
	"fmt"
	"io"
	"os"
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

	TestHelper(t)

	if err != nil {
		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}

		fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected no error, got:\u001B[0m\n%s", err.Error()), messages...)
	}
}

// KO checks whether err is not nil.
//
// Reverse of OK, and also, well, you know 💥🥊.
//
// This could be replaced with using Eq, but since checking the error in Go
// is done lots, it is nicer to read in tests.
func KO(t *testing.T, err error, messages ...string) {

	TestHelper(t)

	ko(t, nil, err, messages...)
}

func ko(t *testing.T, out io.Writer, err error, messages ...string) {

	TestHelper(t)

	if err == nil {
		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}
		fatal(t, out, fmt.Sprintf("\u001B[31;1mexpected error\u001B[0m"), messages...)
	}
}

// ErrorIs checks whether the provided errors are the same when not wrapping or whether
// the have-error is wrapped in the have-error.
func ErrorIs(t *testing.T, want, have error, messages ...string) {

	TestHelper(t)

	errorIs(t, nil, want, have, messages...)
}

func errorIs(t *testing.T, out io.Writer, want, have error, messages ...string) {

	TestHelper(t)

	if errors.Unwrap(want) == nil && errors.Unwrap(have) == nil {

		var msg string

		switch {
		case want == nil && have == nil:
			return
		case want != nil && have != nil && want.Error() == have.Error():
			return
		}

		msg = fmt.Sprintf("\u001B[31;1m\nwant error: %v\nhave error: %v\u001B[0m", want, have)

		if _, ok := os.LookupEnv(EnvNoColors); ok {
			msg = fmt.Sprintf("\nwant error: %v\nhave error: %v", want, have)
		}

		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}

		fatal(t, out, msg, messages...)

		return
	}

	if !errors.Is(have, want) {
		diff := fmt.Sprintf("\n\u001b[31;1mwant error:\t\u001b[0m%v\n\u001b[31;1mwrapped in:\t\u001b[0m%v\n", want, have)

		if _, ok := os.LookupEnv(EnvNoColors); ok {
			diff = fmt.Sprintf("\nwant error:\t%v\nwrapped in:\t%v\n", want, have)
		}

		if len(messages) > 0 {
			messages = append([]string{"--"}, messages...)
		}

		fatal(t, out, diff, messages...)

		return
	}
}
