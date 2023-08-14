// Copyright (c) 2022, Geert JM Vanderkelen

package xt

import (
	"fmt"
	"io"
	"regexp"
	"testing"
)

// MatchString tests whether the string s matches the regular expression pattern.
func MatchString(t *testing.T, pattern, s string, messages ...string) {
	TestHelper(t)
	matchString(t, nil, pattern, s, messages...)
}

func matchString(t *testing.T, out io.Writer, pattern, s string, messages ...string) {
	m, err := regexp.MatchString(pattern, s)
	if err != nil {
		panic(err.Error())
	}

	if !m {
		fatal(t, out, fmt.Sprintf("\n\033[31;1mstring:\033[39m ```%s```\n\033[31;1mmust match pattern:\033[39m %s",
			s, pattern), messages...)
	}
}
