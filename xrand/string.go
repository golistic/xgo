/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xrand

import (
	"math/rand"
)

// AlphaNumeric returns an n-length long string containing random
// alphanumeric characters (both lower and uppercase are possible).
// Panics when n is not >1.
func AlphaNumeric(n int) string {

	if n < 1 {
		panic("n must be >1")
	}

	const abc = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	buf := make([]byte, n)
	for i := range buf {
		buf[i] = abc[rand.Intn(len(abc))]
	}

	return string(buf)
}
