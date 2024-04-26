/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xptr

// ValueOr returns the first non-nil value. If the first argument is not nil, it is returned.
// Otherwise, the second argument is returned.
func ValueOr[T any](a *T, b T) T {

	if a != nil {
		return *a
	}

	return b
}
