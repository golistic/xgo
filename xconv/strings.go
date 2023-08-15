// Copyright (c) 2022, Geert JM Vanderkelen

package xconv

import (
	"fmt"
	"strings"
)

// ParseBool returns the boolean value represented by the string.
// It accepts for true: 1, t, true, OK, yes; for false: 0, f, false, no, n (case-insensitive).
// Any other value returns false and an error.
// This works very similar to Go's strconv.ParseBool except that more
// values are accepted, and it is case-insensitive.
func ParseBool(s string) (bool, error) {
	str := strings.ToLower(s)
	switch str {
	case "1", "t", "true", "yes", "y", "ok":
		return true, nil
	case "0", "f", "false", "no", "n":
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean string; was '%s'", s)
}
