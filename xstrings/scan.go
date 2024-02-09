/*
 * Copyright (c) 2024, Geert JM Vanderkelen
 */

package xstrings

import (
	"strings"
	"unicode"
)

// ScanTokens scans a string and returns a slice of tokens.
//
// Tokens are separated by whitespace, unless they are within quotes.
// If a token is within quotes, the whitespace is preserved within the token.
// Any character that is not whitespace or a quote is part of a token.
func ScanTokens(s string) []string {

	var tokens []string
	var inQuotes rune

	var currToken strings.Builder

	s = strings.TrimSpace(s)
	if s == "" {
		return []string{}
	}

	for _, r := range s {
		switch {
		case unicode.IsSpace(r):
			if inQuotes != 0 {
				currToken.WriteRune(r)
			} else if currToken.Len() > 0 {
				tokens = append(tokens, currToken.String())
				currToken.Reset()
			}
		case r == '"' || r == '\'':
			if inQuotes == 0 {
				inQuotes = r
				continue
			} else if inQuotes == r {
				inQuotes = 0
				continue
			}
			fallthrough
		default:
			currToken.WriteRune(r)
		}
	}

	if currToken.Len() > 0 {
		tokens = append(tokens, currToken.String())
	}

	return tokens
}
