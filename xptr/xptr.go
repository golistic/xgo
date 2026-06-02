/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xptr

// Of returns pointer to value.
// Deprecated: use Go 1.26's new()
func Of[T any](value T) *T {
	return &value
}
