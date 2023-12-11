/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xtime

import (
	"testing"
	"time"

	"github.com/golistic/xgo/xt"
)

func TestUTCMidday(t *testing.T) {

	tests := []struct {
		name string
		want time.Time
	}{
		{
			name: "UTC_Midday_Today",
			want: time.Date(time.Now().UTC().Year(), time.Now().UTC().Month(), time.Now().UTC().Day(), 12, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xt.Eq(t, tt.want, UTCMidday())
		})
	}
}

func TestMidday(t *testing.T) {

	t.Parallel()

	tests := []struct {
		name string
		want func() time.Time
	}{
		{
			name: "local midday",
			want: func() time.Time {
				n := time.Now()
				return time.Date(n.Year(), n.Month(), n.Day(), 12, 0, 0, 0, n.Location())
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xt.Eq(t, tt.want(), Midday())
		})
	}
}
