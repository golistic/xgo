// Copyright (c) 2022, Geert JM Vanderkelen

package xconv

import (
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestUnsignedAsUint64(t *testing.T) {
	uint64_ := uint64(123)
	uint_ := uint(123)
	uint8_ := uint8(123)
	uint16_ := uint16(123)
	uint32_ := uint32(123)

	t.Run("valid argument type", func(t *testing.T) {
		var cases = []any{
			uint_,
			uint8_,
			uint16_,
			uint32_,
			uint64_,
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, uint64_, UnsignedAsUint64(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			1, // int
			int8(1),
			int16(1),
			int32(1),
			int64(1),
			float32(1),
			float64(1),
			&uint_,
			&uint8_,
			&uint16_,
			&uint32_,
			&uint64_,
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = UnsignedAsUint64(c)
				})
			})
		}
	})
}

func TestUnsignedAsUint64Ptr(t *testing.T) {
	uint64_ := uint64(123)
	uint_ := uint(123)
	uint8_ := uint8(123)
	uint16_ := uint16(123)
	uint32_ := uint32(123)

	t.Run("valid argument type", func(t *testing.T) {
		var cases = []any{
			&uint_,
			&uint8_,
			&uint16_,
			&uint32_,
			&uint64_,
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, uint64_, *UnsignedAsUint64Ptr(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			1, // int
			int8(1),
			int16(1),
			int32(1),
			int64(1),
			float32(1),
			float64(1),
			uint_,
			uint8_,
			uint16_,
			uint32_,
			uint64_,
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = UnsignedAsUint64Ptr(c)
				})
			})
		}
	})
}

func TestSignedAsInt64(t *testing.T) {
	int64_ := int64(-123)
	int_ := -123
	int8_ := int8(-123)
	int16_ := int16(-123)
	int32_ := int32(-123)

	t.Run("valid argument type", func(t *testing.T) {
		var cases = []any{
			int_,
			int8_,
			int16_,
			int32_,
			int64_,
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, int64_, SignedAsInt64(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			uint(1),
			uint8(1),
			uint16(1),
			uint32(1),
			uint64(1),
			float32(1),
			float64(1),
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = SignedAsInt64(c)
				})
			})
		}
	})
}

func TestSignedAsInt64Ptr(t *testing.T) {
	int64_ := int64(-123)
	int_ := -123
	int8_ := int8(-123)
	int16_ := int16(-123)
	int32_ := int32(-123)

	t.Run("valid argument type", func(t *testing.T) {
		var cases = []any{
			&int_,
			&int8_,
			&int16_,
			&int32_,
			&int64_,
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Eq(t, int64_, *SignedAsInt64Ptr(c))
			})
		}
	})

	t.Run("invalid argument type", func(t *testing.T) {
		var cases = []any{
			int_,
			int8_,
			int16_,
			int32_,
			int64_,
			float32(1),
			float64(1),
			"string",
			[]byte("byte"),
		}

		for _, c := range cases {
			t.Run("", func(t *testing.T) {
				xt.Panics(t, func() {
					_ = SignedAsInt64Ptr(c)
				})
			})
		}
	})
}
