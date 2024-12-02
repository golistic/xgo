/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xslice

import (
	"iter"
)

// RemoveFirst removes the first element from the given slice and returns the resulting slice.
// If the slice is empty, it returns an empty slice.
func RemoveFirst[S ~[]E, E any](s S) []E {

	if len(s) == 0 {
		return []E{}
	}

	return s[1:]
}

// RemoveFirstN removes the first "n" elements from the given slice and returns the resulting slice.
// If the slice is empty or "n" is larger than the length of the slice, it returns an empty slice.
func RemoveFirstN[S ~[]E, E any](s S, n int) []E {

	if len(s) == 0 || n > len(s) {
		return []E{}
	}

	return s[n:]
}

// Exclude returns an iterator that yields the slice elements in order,
// excluding the elements in toExclude.
func Exclude[S ~[]E, E comparable](s S, toExclude []E) iter.Seq[E] {

	excludeSet := make(map[E]struct{}, len(toExclude))
	for _, item := range toExclude {
		excludeSet[item] = struct{}{}
	}

	return func(yield func(E) bool) {
		for _, item := range s {
			if _, found := excludeSet[item]; !found {
				if !yield(item) {
					return
				}
			}
		}
	}
}
