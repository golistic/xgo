// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"
)

var TestHelper = (*testing.T).Helper

func fatal(t *testing.T, out io.Writer, mainMsg string, messages ...string) {
	TestHelper(t)

	var m string
	if len(messages) > 0 {
		m = "\n\n" + strings.Join(messages, "\n") + "\n"
	}

	msg := mainMsg + m
	if out != nil {
		if _, err := out.Write([]byte(msg)); err != nil {
			panic(fmt.Sprintf("failed writing to out: %s", err))
		}
	} else {
		// this cannot be tested
		t.Fatal(msg)
	}
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	rv, ok := v.(reflect.Value)
	if !ok {
		rv = reflect.ValueOf(v)
	}

	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Slice:
		return rv.IsNil()
	}

	return false
}
