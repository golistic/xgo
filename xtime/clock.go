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

// UTCMidday returns the current time adjusted to midday (12:00:00) in UTC.
func UTCMidday() time.Time {

	now := time.Now().UTC()

	return time.Date(now.Year(), now.Month(), now.Day(), 12, 0, 0, 0, time.UTC)
}
