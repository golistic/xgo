// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import (
	"math"
	"testing"

	"github.com/golistic/xgo/xt"
)

type StringTyped string

func TestRepeatSlice(t *testing.T) {
	t.Run("count is 0", func(t *testing.T) {
		n := RepeatSlice("Go", 0)
		xt.Eq(t, 0, len(n))
	})

	t.Run("create slice with 100 elements", func(t *testing.T) {
		exp := 100
		n := RepeatSlice("Go", exp)
		xt.Eq(t, exp, len(n))
		xt.Eq(t, n[0], n[len(n)-1])
	})

	t.Run("negative count panics", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = RepeatSlice("Go", -1)
		})
	})

	t.Run("len out of range", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = RepeatSlice("Go", math.MaxInt64)
		})
	})
}

func TestRepeatJoin(t *testing.T) {
	t.Run("count is 0", func(t *testing.T) {
		exp := ""
		s := RepeatJoin("Go", 0, ",")
		xt.Eq(t, exp, s)
	})

	t.Run("count is 1", func(t *testing.T) {
		exp := "Go"
		s := RepeatJoin("Go", 1, ",")
		xt.Eq(t, exp, s)
	})

	t.Run("repeat 10 times and join with single character", func(t *testing.T) {
		exp := "?,?,?,?,?,?,?,?,?,?"
		s := RepeatJoin("?", 10, ",")
		xt.Eq(t, exp, s)
	})

	t.Run("repeat 10 times and join with multi character", func(t *testing.T) {
		exp := "Go;;Go;;Go;;Go;;Go;;Go;;Go;;Go;;Go;;Go"
		s := RepeatJoin("Go", 10, ";;")
		xt.Eq(t, exp, s)
	})

	t.Run("negative count panics", func(t *testing.T) {
		xt.Panics(t, func() {
			_ = RepeatJoin("Go", -1, ";")
		})
	})
}

func TestJoin(t *testing.T) {
	t.Run("zero string elements", func(t *testing.T) {
		exp := ""
		xt.Eq(t, exp, Join([]string{}, ","))
	})

	t.Run("1 string element", func(t *testing.T) {
		exp := "first"
		xt.Eq(t, exp, Join([]string{"first"}, ","))
	})

	t.Run("multiple string elements", func(t *testing.T) {
		exp := "first,second,third"
		xt.Eq(t, exp, Join([]string{"first", "second", "third"}, ","))
	})

	t.Run("zero string type elements", func(t *testing.T) {
		exp := ""
		xt.Eq(t, exp, Join([]StringTyped{}, ","))
	})

	t.Run("1 string type element", func(t *testing.T) {
		exp := "first"
		xt.Eq(t, exp, Join([]StringTyped{"first"}, ";"))
	})

	t.Run("multiple string type elements", func(t *testing.T) {
		exp := "first;second;third"
		xt.Eq(t, exp, Join([]StringTyped{"first", "second", "third"}, ";"))
	})
}
