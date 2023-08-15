// Copyright (c) 2022, Geert JM Vanderkelen

package xtime

import "time"

// Yesterday returns the local time of yesterday.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// UTCYesterday returns the local time of yesterday with the location set to UTC.
func UTCYesterday() time.Time {
	return time.Now().UTC().AddDate(0, 0, -1)
}
