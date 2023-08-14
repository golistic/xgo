// Copyright (c) 2020, Geert JM Vanderkelen

package xnet_test

import (
	"testing"

	"github.com/golistic/xgo/xnet"
	"github.com/golistic/xgo/xt"
)

func TestGetTCPPort(t *testing.T) {
	t.Run("get a TCP port", func(t *testing.T) {
		results := map[int]bool{}
		for i := 30; i > 0; i-- {
			p := xnet.MustGetLocalhostTCPPort()
			xt.Assert(t, !results[p], "expected automatically chosen ports to be unique")
			results[p] = true
		}
	})

	t.Run("get TCP port without IP being available", func(t *testing.T) {
		_, err := xnet.GetTCPPort("127.0.0.217")
		xt.KO(t, err)
	})
}
