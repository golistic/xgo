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
		have := bytes.NewBuffer(nil)
		ok(t, have, fmt.Errorf("I am error"), "really wanted no error")
		exp := []byte("\u001B[31;1mexpected no error, got:\u001B[0m\nI am error\n\n--\nreally wanted no error\n")
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp), have.String())
		}
	})
}

func TestKO(t *testing.T) {
	t.Run("error is not nil", func(t *testing.T) {
		KO(t, fmt.Errorf("I am error"))
	})

	t.Run("error is nil", func(t *testing.T) {
		have := bytes.NewBuffer(nil)
		ko(t, have, nil, "really wanted an error")
		exp := []byte("\u001B[31;1mexpected error\u001B[0m\n\n--\nreally wanted an error\n")
		fmt.Println(have.String())
		if !bytes.Equal(exp, have.Bytes()) {
			t.Fatal("expected:", string(exp), have.String())
		}
	})
}
