/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xreflect_test

import (
	"testing"
	"time"

	"github.com/golistic/xgo/xptr"
	"github.com/golistic/xgo/xreflect"
	"github.com/golistic/xgo/xt"
)

func TestPatchStruct(t *testing.T) {
	t.Run("patch struct", func(t *testing.T) {

		type targetType struct {
			FieldInt    int
			FieldPtrInt *int
		}

		patcher := &targetType{FieldInt: 23, FieldPtrInt: xptr.Of(2383)}

		t.Run("patch int", func(t *testing.T) {
			target := &targetType{}

			patched, err := xreflect.PatchStruct(target, patcher, "FieldInt")
			xt.OK(t, err)
			xt.Assert(t, patched)
			xt.Eq(t, patcher.FieldInt, target.FieldInt)
			xt.Eq(t, nil, target.FieldPtrInt, "expected no update for field FieldPtrInt")
		})

		t.Run("patch ptr int", func(t *testing.T) {
			target := &targetType{}

			patched, err := xreflect.PatchStruct(target, patcher, "FieldPtrInt")
			xt.OK(t, err)
			xt.Assert(t, patched)
			xt.Eq(t, patcher.FieldPtrInt, target.FieldPtrInt)
			xt.Eq(t, 0, target.FieldInt, "expected no update for field FieldInt")
		})
	})

	t.Run("patching with patcher using pointer-field", func(t *testing.T) {

		type targetType struct {
			FieldInt    int
			FieldPtrInt *int
			FieldTime   time.Time
		}

		target := &targetType{}

		type patcherType struct {
			FieldInt  *int
			FieldTime *time.Time
		}

		expInt := 5
		exptTime := time.Now()
		patcher := &patcherType{
			FieldInt:  xptr.Of(expInt),
			FieldTime: xptr.Of(exptTime),
		}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.OK(t, err)
		xt.Assert(t, patched)
		xt.Eq(t, expInt, target.FieldInt)
		xt.Eq(t, exptTime, target.FieldTime)
	})

	t.Run("fail patching with target having pointer-field", func(t *testing.T) {

		type targetType struct {
			FieldInt    int
			FieldPtrInt *int
		}

		target := &targetType{}

		type patcherType struct {
			FieldPtrInt int
		}

		patcher := &patcherType{FieldPtrInt: 6}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.KO(t, err)
		xt.Assert(t, !patched)
		xt.Eq(t, "patching field FieldPtrInt (kind mismatch; target.ptr <> patcher.int)", err.Error())
	})

	t.Run("using embedded struct", func(t *testing.T) {

		type targetType struct {
			FieldInt    int
			FieldPtrInt *int
		}

		type targetEmbeds struct {
			targetType
		}

		target := &targetEmbeds{}

		type patcherType struct {
			FieldInt    int
			FieldPtrInt *int
		}

		patcher := &patcherType{
			FieldInt:    112,
			FieldPtrInt: xptr.Of(889),
		}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.OK(t, err)
		xt.Assert(t, patched)
		xt.Eq(t, patcher.FieldInt, target.FieldInt)
		xt.Eq(t, patcher.FieldPtrInt, target.FieldPtrInt)
	})

	t.Run("patching pointer fields with same type", func(t *testing.T) {
		type targetType struct {
			FieldPtr *string
		}

		target := &targetType{}

		type patcherType struct {
			FieldPtr *string
		}

		exp := xptr.Of("pointer")
		patcher := &patcherType{FieldPtr: exp}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.OK(t, err)
		xt.Assert(t, patched)
		xt.Eq(t, exp, target.FieldPtr)
	})

	t.Run("patching with nil leaves value in target untouched", func(t *testing.T) {

		type targetType struct {
			Field string
		}

		exp := "leave me be"
		target := &targetType{
			Field: exp,
		}

		type patcherType struct {
			Field *string
		}

		patcher := &patcherType{}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.OK(t, err)
		xt.Assert(t, !patched)
		xt.Eq(t, exp, target.Field)
	})

	t.Run("fail patching with different types", func(t *testing.T) {
		type targetType struct {
			FieldPtr *int
		}

		target := &targetType{}

		type patcherType struct {
			FieldPtr *string
		}

		patcher := &patcherType{FieldPtr: xptr.Of("pointer")}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.KO(t, err)
		xt.Assert(t, !patched)
		xt.Eq(t, "patching field FieldPtr (type mismatch; target.*int <> patcher.*string)", err.Error())
	})
}
