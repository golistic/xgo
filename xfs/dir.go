/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xfs

import (
	"io/fs"
)

// IsDir checks if a directory exists in the given file system
// at the specified path.
func IsDir(fsys fs.FS, path string) bool {

	if fi, err := fs.Stat(fsys, path); err == nil {
		return fi.Mode().IsDir()
	}

	return false
}
