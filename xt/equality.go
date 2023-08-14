// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

// Eq checks whether the expected value is equal to what we have.
// Messages are printed after
func Eq(t *testing.T, expect, have any, messages ...string) {
	TestHelper(t)
	eq(t, nil, expect, have, messages...)
}

func eq(t *testing.T, out io.Writer, expect, have any, messages ...string) {
	diff := fmt.Sprintf("\n\u001b[31;1mexpect:\t\u001b[0m%v\n\u001b[31;1mhave:\t\u001b[0m%v", expect, have)

	if isNil(expect) || isNil(have) {
		if !(isNil(expect) && isNil(have)) {
			fatal(t, out, diff, messages...)
		}
		return
	}

	expVal := reflect.ValueOf(expect)
	haveType := reflect.TypeOf(have)

	if !expVal.Type().ConvertibleTo(haveType) {

	}

	defer func() {
		r := recover()
		if r != nil {
			fatal(t, out, fmt.Sprintf("\u001b[31;1mcannot compare %T with %T\u001b[0m",
				expect, have))
		}
	}()

	expAny := func() any {
		return expVal.Convert(haveType).Interface()
	}()

	if !reflect.DeepEqual(expAny, have) {
		fatal(t, out, diff, messages...)
	}
}
