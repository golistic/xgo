// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"bytes"
	"testing"
)

func TestMatchString(t *testing.T) {
	t.Run("matches", func(t *testing.T) {
		MatchString(t, "^go", "golistic")
	})

	t.Run("does not match", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		matchString(t, have, "^go", "holistic")
		exp := []byte("\n\u001B[31;1mstring:\u001B[39m ```holistic```\n\u001B[31;1m" +
			"must match pattern:\u001B[39m ^go")
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp))
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
