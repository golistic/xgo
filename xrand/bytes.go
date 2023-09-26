/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xrand

import "crypto/rand"

// RandomBytes returns a n-length slice containing random bytes.
// Panics when n is not >1.
func RandomBytes(n int) ([]byte, error) {

	if n < 1 {
		panic("n must be >1")
	}

	b := make([]byte, n)

	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
