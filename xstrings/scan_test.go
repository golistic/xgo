/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xstrings

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestScanTokens(t *testing.T) {

	var cases = []struct {
		got  string
		want []string
	}{
		{
			" simple   with  variable spaces in Between    ",
			[]string{"simple", "with", "variable", "spaces", "in", "Between"},
		},
		{
			"with \n newline",
			[]string{"with", "newline"},
		},
		{
			`token is "double quoted"`,
			[]string{"token", "is", "double quoted"},
		},
		{
			`token is "double quoted" in middle of string`,
			[]string{"token", "is", "double quoted", "in", "middle", "of", "string"},
		},
		{
			`token is 'single quoted' in middle of string`,
			[]string{"token", "is", "single quoted", "in", "middle", "of", "string"},
		},
		{
			`single in doubles "a string's token" `,
			[]string{"single", "in", "doubles", "a string's token"},
		},
		{
			`doubles in singles 'token with "double" quotes' `,
			[]string{"doubles", "in", "singles", `token with "double" quotes`},
		},
		{
			``,
			[]string{},
		},
	}

	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			xt.Eq(t, c.want, ScanTokens(c.got))
		})
	}
}
