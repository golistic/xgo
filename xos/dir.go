// Copyright (c) 2023, Geert JM Vanderkelen

package xos

import (
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// IsDir returns whether path is a directory.
func IsDir(path string) bool {
	if fi, err := os.Stat(path); err == nil {
		return fi.Mode().IsDir()
	}
	return false
}

// RegularFilesInDir returns regular files found in directory path.
func RegularFilesInDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var l []string
	for _, entry := range entries {
		if !entry.Type().IsRegular() {
			continue
		}
		l = append(l, entry.Name())
	}

	sort.Strings(l)
	return l, nil
}

// RegularFilesInDirWithFullPath returns regular files found in directory path with
// each path included in the filename.
func RegularFilesInDirWithFullPath(path string) ([]string, error) {
	files, err := RegularFilesInDir(path)
	if err != nil {
		return nil, err
	}

	absFiles := make([]string, len(files))
	for i, f := range files {
		absFiles[i] = filepath.Join(path, f)
	}
	return absFiles, nil
}

// FilesInDir returns files which are not directories found in directory path.
// The result is alphabetically sorted.
func FilesInDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var l []string
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		l = append(l, entry.Name())
	}

	sort.Strings(l)
	return l, nil
}

// AllFilenamesInDir returns files which are not directories recursively
// found in directory path.
// The result is alphabetically sorted.
func AllFilenamesInDir(root string) ([]string, error) {
	var result []string

	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			result = append(result, strings.TrimPrefix(strings.TrimPrefix(path, root), "/"))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Strings(result)

	return result, nil
}
