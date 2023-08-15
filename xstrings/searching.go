// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

// SliceHas returns true when x is found in slice a.
func SliceHas[T ~string](a []T, x T) bool {
	return Search(a, x) > -1
}

// Search returns the position of x in slice a or -1 when x is not part of a.
func Search[T ~string](a []T, x T) int {
	for i, e := range a {
		if e == x {
			return i
		}
	}

	return -1
}
