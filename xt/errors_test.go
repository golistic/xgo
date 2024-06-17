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
		wantErr := fmt.Errorf("I am error")
		gotErr := fmt.Errorf("actual error")

		errorIs(t, out, wantErr, gotErr)

		wantOutput := `
want error:	I am error
wrapped in:	actual error
`
		if wantOutput != out.String() {
			t.Fatal("want:", wantOutput, out.String())
		}
	})

	t.Run("wrapped error does match", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		wantErr := fmt.Errorf("I am error")
		haveError := fmt.Errorf("actual error: %w", wantErr)

		if !errorIs(t, out, wantErr, haveError) {
			t.Fatal("expected success", "have:", out.String())
		}
	})

	t.Run("not wrapped error does match", func(t *testing.T) {
		out := bytes.NewBuffer(nil)
		wantErr := fmt.Errorf("I am error")
		haveError := wantErr

		if !errorIs(t, out, wantErr, haveError) {
			t.Fatal("expected success", "have:", out.String())
		}
	})
}
