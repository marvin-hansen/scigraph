package collection_types

import "container/list"

// OrderedMap stores keys in a double linked list and KeyValueHolder as List elements in a hashmap index by those keys
// NON-thread safe / Zero write protection.
type OrderedMap[K comparable, V any] struct {
	store map[K]*list.Element
	keys  *list.List
}

// NewSizedOrderedMap creates a new SyncedOrderedMap for keys of type K and values of type V
// This constructor trades off space for better performance in terms of time complexity
// by over-allocating initially capacity by a multiplier to prevent dynamic map resize.
//
// To create an [int, string] map of a total size of 500 entries, call:
// 	m := collection_types.NewOrderedMap[int, string](100, 5)
func NewSizedOrderedMap[K comparable, V any](size, multiple int) *OrderedMap[K, V] {
	capacity := size * multiple
	return &OrderedMap[K, V]{
		store: make(map[K]*list.Element, capacity),
		keys:  list.New(),
	}
}

// NewOrderedMap creates a new SyncedOrderedMap for keys of type K and values of type V
// To create an [int, string] map call
// 	m := collection_types.NewOrderedMap[int, string]()
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		store: make(map[K]*list.Element),
		keys:  list.New(),
	}
}

// Set updates the value for the given key or inserts the key value pair if the key does not exist yet.
// If the key is already in the map, it replaces the value in the map with the given parameter value.
func (m *OrderedMap[K, V]) Set(key K, val V) {
	var e *list.Element
	if _, exists := m.store[key]; !exists {
		e = m.keys.PushBack(KeyValueHolder[K, V]{
			key:   key,
			value: val,
		})
	} else {
		e = m.store[key]
		e.Value = KeyValueHolder[K, V]{
			key:   key,
			value: val,
		}
	}
	m.store[key] = e
}

// Get returns a value V and true  for the given key if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(n) at worst, or 0(1) at best.
func (m *OrderedMap[K, V]) Get(key K) (value V, ok bool) {
	return m.get(key)
}

// GetFirst returns the first value V and true if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(1).
func (m *OrderedMap[K, V]) GetFirst() (value V, ok bool) {
	k := m.keys.Front().Value.(KeyValueHolder[K, V])
	return m.get(k.key)
}

// GetLast returns the last value V and true if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(1).
func (m *OrderedMap[K, V]) GetLast() (value V, ok bool) {
	k := m.keys.Back().Value.(KeyValueHolder[K, V])
	return m.get(k.key)
}

// get returns a value V and true  for the given key if it exits.
// If the key does not exist, or in case of an empty map,
// get returns an empty value and false.
func (m *OrderedMap[K, V]) get(key K) (value V, ok bool) {
	if m.keys.Len() == 0 {
		return *new(V), false
	}
	if val, exists := m.store[key]; !exists {
		return *new(V), false
	} else {
		value = val.Value.(KeyValueHolder[K, V]).value
		return value, true
	}
}

// Delete removes the key from the list and the value from the map.
// If the key does not exist, nothing happen.
// The complexity is O(1).
func (m *OrderedMap[K, V]) Delete(key K) {
	e, exists := m.store[key]
	if !exists {
		return
	} else {
		m.keys.Remove(e)
		delete(m.store, key)
		return
	}
}

// Size returns the number of elements stored in the OrderedMap.
// The complexity is O(1).
func (m *OrderedMap[K, V]) Size() int {
	return m.keys.Len()
}

// Iterator returns an indexed iterator over the entire key, value collection stored in the SyncedOrderedMap.
// Example usage:
//iterator := m.Iterator()
//for {
//	i, k, v := iterator() // index, key, value
//	if i == nil {
//		break
//	}
//	fmt.Println(*k, v+" is a string")
//}
func (m OrderedMap[K, V]) Iterator() func() (*int, *K, V) {
	e := m.keys.Front()
	j := 0
	return func() (_ *int, _ *K, _ V) {
		if e == nil {
			return
		}
		keyVal := e.Value.(KeyValueHolder[K, V])
		j++
		e = e.Next()
		return func() *int { v := j - 1; return &v }(), &keyVal.key, keyVal.value
	}
}
