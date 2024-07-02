// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"bytes"
	"fmt"
	"testing"
)

func TestOK(t *testing.T) {
	t.Run("error is nil", func(t *testing.T) {
		OK(t, nil)
	})

	t.Run("error is not nil", func(t *testing.T) {
		out := bytes.NewBuffer(nil)

		ok(t, out, fmt.Errorf("I am error"), "really wanted no error")
		exp := []byte("\u001B[31;1mexpected no error, got:\u001B[0m\nI am error\n\n--\nreally wanted no error\n")

		if !bytes.Equal(exp, out.Bytes()) {
			t.Fatal("expected:", string(exp), out.String())
		}
	})
}

func TestKO(t *testing.T) {
	t.Run("error is not nil", func(t *testing.T) {
		KO(t, fmt.Errorf("I am error"))
	})

	t.Run("error is nil", func(t *testing.T) {
		out := bytes.NewBuffer(nil)

		ko(t, out, nil, "really wanted an error")
		exp := []byte("\u001B[31;1mexpected error\u001B[0m\n\n--\nreally wanted an error\n")

		fmt.Println(out.String())

		if !bytes.Equal(exp, out.Bytes()) {
			t.Fatal("expected:", string(exp), out.String())
		}
	})
}

func TestErrorIs(t *testing.T) {
	t.Run("non-wrapped error does not match", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		want := fmt.Errorf("I am error")
		have := fmt.Errorf("actual error")

		errorIs(t, out, want, have)

		wantOutput := `
want error: I am error
have error: actual error`

		if wantOutput != out.String() {
			t.Fatal("\nwant:", wantOutput, want, "\nhave:", out.String())
		}
	})

	t.Run("wrapped error does match", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		want := fmt.Errorf("I am error")
		have := fmt.Errorf("actual error: %w", want)

		errorIs(t, out, want, have)

		if out.String() != "" {
			t.Fatal("expected success", "have:", out.String())
		}
	})

	t.Run("not wrapped error does match", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		want := fmt.Errorf("I am error")
		have := fmt.Errorf("I am error")

		errorIs(t, out, want, have)

		if out.String() != "" {
			t.Fatal("expected success", "have:", out.String())
		}
	})

	t.Run("want but have is nil", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		want := fmt.Errorf("I am error")
		var have error

		errorIs(t, out, want, have)

		if out.String() != "\nwant error: I am error\nhave error: <nil>" {
			t.Fatal("expected success", "have:", out.String())
		}
	})

	t.Run("want is null but have not", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		var want error
		have := fmt.Errorf("I am error")

		errorIs(t, out, want, have)

		if out.String() != "\nwant error: <nil>\nhave error: I am error" {
			t.Fatal("expected success", "have:", out.String())
		}
	})
}
