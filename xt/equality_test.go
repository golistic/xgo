// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestEq(t *testing.T) {
	tsNow := time.Now()

	t.Run("equal", func(t *testing.T) {
		var cases = map[string]struct {
			exp  any
			have any
		}{
			"string":    {exp: "golistic", have: "golistic"},
			"int":       {exp: 123, have: 123},
			"int64":     {exp: int64(1234), have: int64(1234)},
			"bytes":     {exp: []byte("bytes are equal"), have: []byte("bytes are equal")},
			"time.Time": {exp: tsNow, have: tsNow},
		}

		for cn, c := range cases {
			t.Run(cn, func(t *testing.T) {
				Eq(t, c.exp, c.have)
			})
		}
	})

	t.Run("does not match", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		eq(t, have, "not", "equal")
		exp := []byte("\n\u001B[31;1mexpect:\t\u001B[0mnot\n\u001B[31;1m" +
			"have:\t\u001B[0mequal")
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp), have.String())
		}
	})

	t.Run("check nil", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		eq(t, have, nil, "nil")
		fmt.Println(have.String())
		exp := []byte("\n\u001B[31;1mexpect:\t\u001B[0m<nil>\n\u001B[31;1m" +
			"have:\t\u001B[0mnil")
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp), have.String())
		}
	})

	t.Run("not ConvertibleTo", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		v := 5
		eq(t, have, []int{1, 2, 3}, &v)
		fmt.Println(have.String())
		exp := []byte("\u001B[31;1mcannot compare []int with *int\u001B[0m")
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp), have.String())
		}
	})

	t.Run("panics if regular expression is bad", func(t *testing.T) {
		defer func() {
			r := recover()
			if r == nil {
				t.Fatalf("expected panic")
			}
		}()

		matchString(t, nil, "^(go", "golistic")
	})
}
