// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import "os"

// IsRegularFile returns whether path is a regular file.
func IsRegularFile(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		return fi.Mode().IsRegular()
	}
	return false
}
