// Copyright (c) 2023, 2024, Geert JM Vanderkelen

package xmaps

import "slices"

// OrderedMap wraps around a Go map keeping the order with which
// elements have been added. Keys must be comparable, but values
// can be anything.
// Unlike map, index assigment is not possible. Use the `Set`
// method to set a key with a particular value.
// Use the Keys method to retrieves keys, Values to get the
// values. To get both, which probably what you want, use
// the KeysValues method.
type OrderedMap[T comparable] struct {
	pairs map[T]any
	order []T
}

// Count returns the number of elements in the map.
func (om *OrderedMap[T]) Count() int {
	return len(om.order)
}

// Set key in OrderedMap to value. Previously stored values
// are overwritten, but the order does not change.
func (om *OrderedMap[T]) Set(key T, value any) {

	if om.pairs == nil {
		om.pairs = map[T]any{}
	}

	om.pairs[key] = value
	if !om.has(key) {
		om.order = append(om.order, key)
	}
}

// Delete deletes the element with the specified key from
// the OrderedMap.
func (om *OrderedMap[T]) Delete(key T) {

	om.order = slices.DeleteFunc(om.order, func(t T) bool {
		return t == key
	})

	delete(om.pairs, key)
}

// Keys returns keys as slice of string.
func (om *OrderedMap[T]) Keys() []T {

	return om.order
}

func (om *OrderedMap[T]) values() []any {

	res := make([]any, len(om.order))
	for i, k := range om.order {
		res[i] = om.pairs[k]
	}

	return res
}

// Values returns the values as slice of interfaces.
func (om *OrderedMap[T]) Values() []any {

	return om.values()
}

// KeysValues returns the keys as slice of strings, and values as slice of interfaces.
func (om *OrderedMap[T]) KeysValues() ([]T, []any) {

	return om.order, om.values()
}

// Has returns whether the map contains key.
func (om *OrderedMap[T]) Has(key T) bool {

	return om.has(key)
}

func (om *OrderedMap[T]) has(key T) bool {

	for _, e := range om.order {
		if e == key {
			return true
		}
	}

	return false
}

// Value returns the value for key and also whether it was found.
// The bool is returned because value could be nil.
func (om *OrderedMap[T]) Value(key T) (any, bool) {

	return om.pairs[key], om.has(key)
}
