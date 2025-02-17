/*
 * Copyright (c) 2025, Geert JM Vanderkelen
 */

package xjson

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Format formats the given value as an indented JSON string and returns it
// or an error message string if it fails.
func Format(v any) string {

	var b bytes.Buffer

	data, err := json.Marshal(v)
	if err != nil {
		return fmt.Sprintf("JSON encoding failed (marshal): %s", err.Error())
	}

	if err := json.Indent(&b, data, "", "  "); err != nil {
		return fmt.Sprintf("JSON formatting failed (indent): %s", err.Error())
	}

	return b.String()
}
