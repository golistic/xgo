/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xtime

import "time"

// Midday returns the time corresponding to 12:00:00 on the current local date.
func Midday() time.Time {

	n := time.Now()

	return time.Date(n.Year(), n.Month(), n.Day(), 12, 0, 0, 0, time.Local)
}

// MiddayForDate returns the time corresponding to 12:00:00 on the given year, month, and day.
// The time is returned in the local time zone.
func MiddayForDate(year int, month time.Month, day int) time.Time {

	return time.Date(year, month, day, 12, 0, 0, 0, time.Local)
}

// UTCMidday returns the time corresponding to 12:00:00 in UTC on the current UTC date.
func UTCMidday() time.Time {

	now := time.Now().UTC()

	return time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
}

// UTCMiddayForDate returns the time corresponding to 12:00:00 in UTC for the given year, month, and day.
func UTCMiddayForDate(year int, month time.Month, day int) time.Time {

	return time.Date(year, month, day, 12, 0, 0, 0, time.UTC)
}
