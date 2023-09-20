/*
 * Copyright (c) 2023, Geert JM Vanderkelen
 */

package xreflect_test

import (
	"testing"

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

	t.Run("fail patching with patcher using pointer-field", func(t *testing.T) {

		type targetType struct {
			FieldInt    int
			FieldPtrInt *int
		}

		target := &targetType{}

		type patcherType struct {
			FieldInt *int
		}

		patcher := &patcherType{FieldInt: xptr.Of(5)}

		patched, err := xreflect.PatchStruct(target, patcher)
		xt.KO(t, err)
		xt.Assert(t, !patched)
		xt.Eq(t, "patching field FieldInt (type mismatch; patcher.ptr <> target.int)", err.Error())
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
		xt.Eq(t, "patching field FieldPtrInt (type mismatch; patcher.int <> target.ptr)", err.Error())
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
}
