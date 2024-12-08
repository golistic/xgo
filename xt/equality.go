// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

// Eq checks if the `want` value is equal to the `have` value during testing.
// If not equal, it triggers a test failure with optional custom `messages`.
//
// This function uses reflection for comparison and handles nil values.
func Eq(t *testing.T, want, have any, messages ...string) {

	TestHelper(t)
	eq(t, nil, want, have, messages...)
}

func eq(t *testing.T, out io.Writer, want, have any, messages ...string) {

	TestHelper(t)
	diff := fmt.Sprintf("\n\u001b[31;1mexpect:\t\u001b[0m%v\n\u001b[31;1mhave:\t\u001b[0m%v", want, have)

	if isNil(want) || isNil(have) {
		if !(isNil(want) && isNil(have)) {
			fatal(t, out, diff, messages...)
		}
		return
	}

	expVal := reflect.ValueOf(want)
	haveVal := reflect.ValueOf(have)

	if expVal.Type().ConvertibleTo(haveVal.Type()) {
		defer func() {
			if r := recover(); r != nil {
				fatal(t, out, fmt.Sprintf("\u001b[31;1mcannot compare %T with %T\u001b[0m", want, have))
			}
		}()

		// Convert and compare
		expConverted := expVal.Convert(haveVal.Type()).Interface()
		if !reflect.DeepEqual(expConverted, have) {
			fatal(t, out, diff, messages...)
		}
	} else {
		fatal(t, out, fmt.Sprintf("\u001b[31;1mcannot convert %T to %T\u001b[0m", want, have), messages...)
	}
}
