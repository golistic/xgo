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
//
// Only fields with same name and type can be updated.
//
// # When target is a pointer, and
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

		if !patchField.Field.IsExported() || !patchField.Value.IsValid() {
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

		// allow patching of non-pointer fields using pointer
		if patchField.IsPointer() && targetField.Kind() != reflect.Pointer {
			if patchField.Value.IsNil() {
				// we do not touch target
				continue
			}
			patchField.Value = patchField.Value.Elem()
		}

		if pk, tk := patchField.Value.Kind(), targetField.Kind(); pk != tk {
			fmt.Println("### patchField", patchField.Value)
			fmt.Println("### patchField", patchField.Value, patchField.Value.Type(), patchField.Value.Kind())
			return false, fmt.Errorf("patching field %s (kind mismatch; target.%s <> patcher.%s)",
				patchField.Field.Name, tk.String(), pk.String())
		}

		if pk, tk := patchField.Value.Type(), targetField.Type(); pk != tk {
			return false, fmt.Errorf("patching field %s (type mismatch; target.%s <> patcher.%s)",
				patchField.Field.Name, tk.String(), pk.String())
		}

		if !reflect.DeepEqual(patchField.Value.Interface(), targetField) {
			patched = true
		}

		targetField.Set(patchField.Value)
	}

	return patched, nil
}
