// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

// Pointer returns string typed s as pointer.
func Pointer[T ~string](s T) *T {
	return &s
}
