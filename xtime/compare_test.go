/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xtime

import (
	"testing"
	"time"

	"github.com/golistic/xgo/xt"
)

func TestFirstBeforeSecond(t *testing.T) {
	now := time.Now()
	inOneHour := now.Add(1 * time.Hour)
	inTwoHours := now.Add(2 * time.Hour)

	testCases := []struct {
		name string
		a, b *time.Time
		want bool
	}{
		{
			name: "nil values",
			want: true,
		},
		{
			name: "nil value for a",
			b:    &inOneHour,
			want: true,
		},
		{
			name: "nil value for b",
			a:    &inOneHour,
			want: true,
		},
		{
			name: "a and b are zero",
			a:    &time.Time{},
			b:    &time.Time{},
			want: true,
		},
		{
			name: "a is zero",
			a:    &time.Time{},
			b:    &inOneHour,
			want: true,
		},
		{
			name: "b is zero",
			a:    &inOneHour,
			b:    &time.Time{},
			want: true,
		},
		{
			name: "a before b",
			a:    &inOneHour,
			b:    &inTwoHours,
			want: true,
		},
		{
			name: "b before a",
			a:    &inTwoHours,
			b:    &inOneHour,
			want: false,
		},
		{
			name: "a equals b",
			a:    &now,
			b:    &now,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			xt.Eq(t, tc.want, FirstBeforeSecond(tc.a, tc.b))
		})
	}
}
