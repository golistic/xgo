/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xtime

import "time"

// FirstBeforeSecond returns true if the first time is before the second time.
// If either of the times is nil or zero value, it returns true.
func FirstBeforeSecond(a, b *time.Time) bool {

	if (a == nil || a.IsZero()) || (b == nil || b.IsZero()) {
		return true
	}

	return a.Before(*b)
}
