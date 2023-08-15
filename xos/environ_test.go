// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"fmt"
	"os"
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestEnvironMap(t *testing.T) {
	t.Run("get current environment as map", func(t *testing.T) {
		envmap := EnvironMap()

		for name, value := range envmap {
			exp, ok := os.LookupEnv(name)
			xt.Assert(t, ok, fmt.Sprintf("expected envvar %s to be availalbe", name))
			xt.Eq(t, exp, value)
		}
	})
}

func TestEnvironMapLookup(t *testing.T) {
	t.Run("using name", func(t *testing.T) {
		envmap := EnvironMapLookup(func(name, _ string) bool {
			return name == "GOROOT" || name == "GOPATH"
		})

		xt.Eq(t, 2, len(envmap))

		for name, value := range envmap {
			exp, ok := os.LookupEnv(name)
			xt.Assert(t, ok, fmt.Sprintf("expected envvar %s to be availalbe", name))
			xt.Eq(t, exp, value)
		}
	})

	t.Run("using value", func(t *testing.T) {
		envvar := "GOROOT"
		exp, ok := os.LookupEnv(envvar)
		xt.Assert(t, ok, fmt.Sprintf("expected %s in environment", envvar))

		envmap := EnvironMapLookup(func(_, value string) bool {
			return value == exp
		})

		xt.Eq(t, 1, len(envmap))
		xt.Eq(t, exp, envmap[envvar])
	})

	t.Run("nothing matches", func(t *testing.T) {
		envmap := EnvironMapLookup(func(_, value string) bool {
			return false
		})

		xt.Eq(t, 0, len(envmap))
	})

	t.Run("no mapper", func(t *testing.T) {
		envmap := EnvironMapLookup(nil)

		for name, value := range envmap {
			exp, ok := os.LookupEnv(name)
			xt.Assert(t, ok, fmt.Sprintf("expected envvar %s to be availalbe", name))
			xt.Eq(t, exp, value)
		}
	})
}
