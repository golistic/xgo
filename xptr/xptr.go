/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xptr

// Of returns pointer to value.
func Of[T any](value T) *T {
	return &value
}
