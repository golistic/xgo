// Copyright (c) 2023, 2024, Geert JM Vanderkelen

package xmaps

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/golistic/xgo/xt"
)

func TestOrderedMap(t *testing.T) {
	t.Run("empty", func(t *testing.T) {

		om := OrderedMap[string, any]{}
		xt.Eq(t, 0, om.Count())
		xt.Eq(t, 0, len(om.Keys()))
		xt.Eq(t, 0, len(om.Values()))
		keys, values := om.KeysValues()
		xt.Eq(t, 0, len(keys))
		xt.Eq(t, 0, len(values))
		xt.Eq(t, false, om.Has("somekey"))
	})

	t.Run("retrieve keys and values", func(t *testing.T) {

		om := OrderedMap[string, any]{}
		om.Set("key3", "value3")
		om.Set("key1", 1.1)
		om.Set("key2", 2)
		xt.Eq(t, true, om.Has("key2"))

		expKeys := []string{"key3", "key1", "key2"}
		expValues := []any{"value3", 1.1, 2}

		xt.Eq(t, expKeys, om.Keys())
		xt.Eq(t, expValues, om.Values())
		keys, values := om.KeysValues()
		xt.Eq(t, expKeys, keys)
		xt.Eq(t, expValues, values)
	})

	t.Run("set already set does not change order", func(t *testing.T) {

		om := OrderedMap[string, any]{}
		om.Set("key3", "value3")
		om.Set("key1", 1.1)
		om.Set("key2", 2)

		om.Set("key3", "value number 3")

		expKeys := []string{"key3", "key1", "key2"}
		expValues := []any{"value number 3", 1.1, 2}

		keys, values := om.KeysValues()
		xt.Eq(t, expKeys, keys)
		xt.Eq(t, expValues, values)
		xt.Eq(t, len(expKeys), om.Count())
	})

	t.Run("retrieve key", func(t *testing.T) {

		om := OrderedMap[string, any]{}
		om.Set("key3", "value3")
		om.Set("key1", 1.1)
		om.Set("key2", 2)
		om.Set("key4", nil)

		cases := map[string]struct {
			key     string
			exp     any
			expHave bool
		}{
			"nil":      {key: "key4", exp: nil, expHave: true},
			"notkey":   {key: "notkey", exp: nil, expHave: false},
			"key3":     {key: "key3", exp: "value3", expHave: true},
			"nilnokey": {key: "nilnokey", exp: nil, expHave: false},
		}

		for name, cs := range cases {
			t.Run(name, func(t *testing.T) {
				v, have := om.Value(cs.key)
				xt.Eq(t, cs.expHave, have)
				xt.Eq(t, cs.exp, v)
			})
		}
	})

	t.Run("delete element", func(t *testing.T) {

		om := OrderedMap[string, any]{}
		om.Set("key3", "value3")
		om.Set("key1", 1.1)
		om.Set("key2", 2)
		om.Set("key4", nil)

		xt.Eq(t, 4, om.Count())
		om.Delete("noSuchKey")

		xt.Eq(t, 4, om.Count())

		om.Delete("key2")
		xt.Eq(t, 3, om.Count())
		_, have := om.Value("key2")
		xt.Eq(t, false, have)

		om.Delete("key3")
		xt.Eq(t, 2, om.Count())
		_, have = om.Value("key3")
		xt.Eq(t, false, have)

		om.Delete("key3")
		xt.Eq(t, 2, om.Count())
		_, have = om.Value("key3")
		xt.Eq(t, false, have)
	})

	t.Run("map[int] retrieve keys and values", func(t *testing.T) {

		om := OrderedMap[int, any]{}
		om.Set(3, "value3")
		om.Set(1, 1.1)
		om.Set(2, 2)
		xt.Eq(t, true, om.Has(2))

		expKeys := []int{3, 1, 2}
		expValues := []any{"value3", 1.1, 2}

		xt.Eq(t, expKeys, om.Keys())
		xt.Eq(t, expValues, om.Values())
		keys, values := om.KeysValues()
		xt.Eq(t, expKeys, keys)
		xt.Eq(t, expValues, values)
	})

	t.Run("using approximation", func(t *testing.T) {

		type foo string

		om := OrderedMap[foo, any]{}
		om.Set("3", "value3")
		om.Set("1", 1.1)
		om.Set("2", 2)
		xt.Eq(t, true, om.Has("2"))
	})

	t.Run("decoded from JSON", func(t *testing.T) {
		want := NewOrderedMap[string, any]()
		want.Set("key3", "value")
		want.Set("key1", 1.1)
		want.Set("key2", 2)

		for range 100 {
			// run a few times to test if it is idempotent
			wantJson, err := want.MarshalJSON()
			xt.OK(t, err)

			have := NewOrderedMap[string, any]()

			xt.OK(t, have.UnmarshalJSON(wantJson))

			haveJson, err := have.MarshalJSON()
			xt.OK(t, err)
			xt.Eq(t, wantJson, haveJson)
		}
	})

	t.Run("encoded to JSON keys must be strings", func(t *testing.T) {
		om := NewOrderedMap[int, any]()
		om.Set(1, "value")
		_, err := om.MarshalJSON()
		xt.ErrorIs(t, err, ErrKeysMustBeStrings)
	})

	t.Run("decoded from JSON errors with invalid keys", func(t *testing.T) {
		om := NewOrderedMap[int, any]()
		err := om.UnmarshalJSON([]byte(`{1:"value"}`))
		xt.KO(t, err)
		var syntaxErr *json.SyntaxError
		xt.Assert(t, errors.As(err, &syntaxErr))
	})
}
