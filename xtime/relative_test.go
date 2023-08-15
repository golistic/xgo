// Copyright (c) 2022, Geert JM Vanderkelen

package xtime

import (
	"testing"
	"time"

	"github.com/golistic/xgo/xt"
)

func TestYesterday(t *testing.T) {
	// this could fail at midnight
	xt.Eq(t, time.Now().Day()-1, Yesterday().Day())
	xt.Eq(t, time.Now().Location(), Yesterday().Location())
}

func TestUTCYesterday(t *testing.T) {
	// this could fail at midnight
	xt.Eq(t, time.Now().UTC().Day()-1, UTCYesterday().Day())
	xt.Eq(t, time.UTC, UTCYesterday().Location())
}
