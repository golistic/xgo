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

	now := time.Now().UTC()
	want := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	xt.Eq(t, want, UTCMidday())
}

func TestUTCMiddayForDate(t *testing.T) {

	now := time.Now().UTC()

	tests := []struct {
		name      string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:      "past date",
			wantYear:  now.Year() - 1,
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "future date",
			wantYear:  now.Year() - 1,
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "now",
			wantYear:  now.Year(),
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "month out of range",
			wantYear:  now.Year(),
			wantMonth: time.Month(13),
			wantDay:   now.Day(),
		},
		{
			name:      "day out of range",
			wantYear:  now.Year(),
			wantMonth: now.Month(),
			wantDay:   32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := time.Date(tt.wantYear, tt.wantMonth, tt.wantDay, 12, 0, 0, 0, time.UTC)
			xt.Eq(t, want, UTCMiddayForDate(tt.wantYear, tt.wantMonth, tt.wantDay))
		})
	}
}

func TestMidday(t *testing.T) {

	now := time.Now()
	want := time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, now.Location())
	xt.Eq(t, want, Midday())
}

func TestMiddayForDate(t *testing.T) {

	now := time.Now()

	tests := []struct {
		name      string
		wantYear  int
		wantMonth time.Month
		wantDay   int
	}{
		{
			name:      "past date",
			wantYear:  now.Year() - 1,
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "future date",
			wantYear:  now.Year() - 1,
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "now",
			wantYear:  now.Year(),
			wantMonth: now.Month(),
			wantDay:   now.Day(),
		},
		{
			name:      "month out of range",
			wantYear:  now.Year(),
			wantMonth: time.Month(13),
			wantDay:   now.Day(),
		},
		{
			name:      "day out of range",
			wantYear:  now.Year(),
			wantMonth: now.Month(),
			wantDay:   32,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := time.Date(tt.wantYear, tt.wantMonth, tt.wantDay, 12, 0, 0, 0, now.Location())
			xt.Eq(t, want, MiddayForDate(tt.wantYear, tt.wantMonth, tt.wantDay))
		})
	}
}
