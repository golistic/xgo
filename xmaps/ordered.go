// Copyright (c) 2023, 2024, Geert JM Vanderkelen

package xmaps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"slices"
	"strings"
)

var ErrKeysMustBeStrings = fmt.Errorf("keys must be strings")

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {

	return &OrderedMap[K, V]{}
}

// OrderedMap wraps around a Go map keeping the order with which
// elements have been added. Keys must be comparable, but values
// can be anything.
// Unlike map, index assigment is not possible. Use the `Set`
// method to set a key with a particular value.
// Use the Keys method to retrieves keys, Values to get the
// values. To get both, which probably what you want, use
// the KeysValues method.
type OrderedMap[K comparable, V any] struct {
	pairs map[K]V
	order []K
}

func (om *OrderedMap[K, V]) init() {

	om.pairs = map[K]V{}
	om.order = []K{}
}

// Count returns the number of elements in the map.
func (om *OrderedMap[K, V]) Count() int {
	return len(om.order)
}

// Set key in OrderedMap to value. Previously stored values
// are overwritten, but the order does not change.
func (om *OrderedMap[K, V]) Set(key K, value V) {

	if om.pairs == nil {
		om.init()
	}

	om.pairs[key] = value
	if !om.has(key) {
		om.order = append(om.order, key)
	}
}

// Delete deletes the element with the specified key from
// the OrderedMap.
func (om *OrderedMap[K, V]) Delete(key K) {

	om.order = slices.DeleteFunc(om.order, func(t K) bool {
		return t == key
	})

	delete(om.pairs, key)
}

// Keys returns keys as slice of string.
func (om *OrderedMap[K, V]) Keys() []K {

	return om.order
}

func (om *OrderedMap[K, V]) values() []V {

	res := make([]V, len(om.order))
	for i, k := range om.order {
		res[i] = om.pairs[k]
	}

	return res
}

// Values returns the values as slice of interfaces.
func (om *OrderedMap[K, V]) Values() []V {

	return om.values()
}

// KeysValues returns the keys as slice of strings, and values as slice of interfaces.
func (om *OrderedMap[K, V]) KeysValues() ([]K, []V) {

	return om.order, om.values()
}

// Has returns whether the map contains key.
func (om *OrderedMap[K, V]) Has(key K) bool {

	return om.has(key)
}

func (om *OrderedMap[K, V]) has(key K) bool {

	for _, e := range om.order {
		if e == key {
			return true
		}
	}

	return false
}

// Value returns the value for key and also whether it was found.
// The bool is returned because value could be nil.
func (om *OrderedMap[K, V]) Value(key K) (any, bool) {

	return om.pairs[key], om.has(key)
}

func (om *OrderedMap[K, V]) MarshalJSON() ([]byte, error) {

	if om.pairs == nil {
		return []byte("null"), nil
	}

	if _, ok := any(om.order[0]).(string); !ok {
		return nil, ErrKeysMustBeStrings
	}

	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, key := range om.order {
		if i > 0 {
			buf.WriteByte(',')
		}
		keyJSON, err := json.Marshal(key)
		if err != nil {
			return nil, err
		}
		valueJSON, err := json.Marshal(om.pairs[key])
		if err != nil {
			return nil, err
		}
		buf.Write(keyJSON)
		buf.WriteByte(':')
		buf.Write(valueJSON)
	}
	buf.WriteByte('}')

	return buf.Bytes(), nil
}

func (om *OrderedMap[K, V]) UnmarshalJSON(data []byte) error {

	decoder := json.NewDecoder(strings.NewReader(string(data)))
	omTmp := NewOrderedMap[K, V]()

	// opening brace
	if t, err := decoder.Token(); err != nil {
		return err
	} else if t != json.Delim('{') {
		return fmt.Errorf("expected JSON object, got %v", t)
	}

	// key-value pairs (object)
	for decoder.More() {
		keyToken, err := decoder.Token()
		if err != nil {
			return fmt.Errorf("error reading key: %w", err)
		}
		if _, ok := keyToken.(string); !ok {
			return fmt.Errorf("expected string key, got %v", keyToken)
		}

		var value V
		if err := decoder.Decode(&value); err != nil {
			return err
		}
		fmt.Println("### value", value)

		omTmp.Set(keyToken.(K), value)
	}

	// closing brace
	if t, err := decoder.Token(); err != nil {
		return err
	} else if t != json.Delim('}') {
		return fmt.Errorf("expected JSON object, got %v", t)
	}

	*om = *omTmp

	return nil
}
