/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xslice

// AsAny takes any kind of slice and returns it as []any.
func AsAny[S ~[]E, E any](s S) []any {
	r := make([]any, len(s))

	for i, e := range s {
		r[i] = any(e)
	}

	return r
}
