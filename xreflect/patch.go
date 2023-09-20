/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xreflect

import (
	"fmt"
	"reflect"
	"strings"
)

// PatchStruct patches target with values found in patcher with both target and patcher
// being structs.
// Only fields with same name, type, and kind (cannot update pointer values with non
// pointer values) are updated.
//
// The fields-argument can be used to select which fields gets patched (case-insensitive).
// If the fields-argument is not provided, all fields of patcher are used to update target.
//
// If any field got update, this function returns true.
func PatchStruct(target any, patcher any, fields ...string) (bool, error) {

	fieldsToPatch := map[string]bool{}
	for _, f := range fields {
		fieldsToPatch[strings.ToLower(f)] = true
	}

	rvPatch := reflect.ValueOf(patcher)
	patchFields := GetFields(rvPatch)

	rtTarget := reflect.TypeOf(target)
	rvTarget := reflect.ValueOf(target)

	var patched bool

	for i := 0; i < len(patchFields); i++ {
		var patchField = patchFields[i]

		if !patchField.Field.IsExported() {
			continue
		}

		// embedded structs
		if patchField.Field.Anonymous {
			patchFields = append(patchFields, GetFields(patchField.Value)...)
			continue
		}

		// skip fields we do not need to patch, but only skip where are fields given
		if len(fieldsToPatch) > 0 && !fieldsToPatch[strings.ToLower(patchField.Field.Name)] {
			continue
		}

		// check if target has field
		if _, ok := rtTarget.Elem().FieldByName(patchField.Field.Name); !ok {
			continue
		}

		targetField := reflect.Indirect(rvTarget).FieldByName(patchField.Field.Name)
		if pk, tk := patchField.Value.Kind(), targetField.Kind(); pk != tk {
			return false, fmt.Errorf("patching field %s (type mismatch; patcher.%s <> target.%s)",
				patchField.Field.Name, pk.String(), tk.String())
		}

		if !reflect.DeepEqual(patchField.Value.Interface(), targetField) {
			patched = true
		}

		targetField.Set(patchField.Value)
	}

	return patched, nil
}
