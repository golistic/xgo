// Copyright (c) 2022, Geert JM Vanderkelen

package xstrings

import "github.com/golistic/xgo/xptr"

// Pointer returns string typed s as pointer.
//
// Deprecated: use xptr.To().
func Pointer[T ~string](s T) *T {
	return xptr.Of(s)
}
