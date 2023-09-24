/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xreflect

import (
	"reflect"
)

type StructField struct {
	Field reflect.StructField
	Value reflect.Value
}

func GetFields(v reflect.Value) []*StructField {

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	rt := v.Type()
	var fields []*StructField

	for i := 0; i < rt.NumField(); i++ {
		fields = append(fields, &StructField{
			Field: rt.Field(i),
			Value: v.Field(i),
		})
	}

	return fields
}

func (s *StructField) IsPointer() bool {
	return s.Value.Kind() == reflect.Ptr
}
