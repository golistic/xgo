/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xt_test

import (
	"fmt"
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestLogAgg(t *testing.T) {
	t.Run("write, find, reset", func(t *testing.T) {
		la := xt.NewLogAgg()
		for i := 1; i <= 3; i++ {
			_, err := la.Write([]byte(fmt.Sprintf("line %d", i)))
			xt.OK(t, err)
		}

		for i := 1; i <= 3; i++ {
			exp := fmt.Sprintf("line %d", i)
			xt.Eq(t, exp, la.Find(t, exp))
		}

		la.Reset()

		for i := 1; i <= 3; i++ {
			xt.Eq(t, "", la.Find(t, fmt.Sprintf("line %d", i)))
		}
	})

	t.Run("find JSON", func(t *testing.T) {
		la := xt.NewLogAgg()
		for i := 1; i <= 3; i++ {
			_, err := la.Write([]byte(fmt.Sprintf(`{"line": "data %d"}`, i)))
			xt.OK(t, err)
		}

		for i := 1; i <= 3; i++ {
			exp := fmt.Sprintf("data %d", i)
			have := la.FindJSON(t, exp)
			xt.Eq(t, exp, have["line"])
		}
	})

	t.Run("get all entries", func(t *testing.T) {
		la := xt.NewLogAgg()

		exp := []string{
			"line 1",
			"line 2",
			"line 3",
		}

		for _, entry := range exp {
			_, err := la.Write([]byte(entry))
			xt.OK(t, err)
		}

		xt.Eq(t, exp, la.Entries())

		exp = append(exp, "line 4")
		_, err := la.Write([]byte(exp[len(exp)-1]))
		xt.OK(t, err)

		xt.Eq(t, exp, la.Entries())
	})
}
