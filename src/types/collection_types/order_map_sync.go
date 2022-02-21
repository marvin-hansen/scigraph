// Copyright (c) 2021-2022. Marvin Hansen | marvin.hansen@gmail.com

package collection_types

import (
	"container/list"
	"sync"
)

// DRAFT RELEASE NOTES — Introduction to Go 1.18
// https://tip.golang.org/doc/go1.18

// The Go Programming Language Specification - Go 1.18 Draft (incomplete)
// https://tip.golang.org/ref/spec#Type_parameters

// Know Go: Generics [pre-order]
// https://bitfieldconsulting.com/books/generics

// Code club
// https://www.youtube.com/playlist?list=PLEcwzBXTPUE_YQR7R0BRtHBYJ0LN3Y0i3

// Ordered Maps for Go 2 (Using Generics)
// https://medium.com/swlh/ordered-maps-for-go-using-generics-875ef3816c71

// OrderMap in Go 2.0 / 1.8 using generics with Delete Op in O(1) Time complexity
// https://www.tugberkugurlu.com/archive/implementing-ordered-map-in-go-2-0-by-using-generics-with-delete-operation-in-o-1-time-complexity
// https://gotipplay.golang.org/p/UJZsQnPRmRh

// Runtime overhead of using defer in go
// "doDefer is roughly 16 times more expensive than doNoDefer. [..]
// In short, don’t use defer in hot code paths. The overhead is non-trivial and not obvious."
// https://medium.com/i0exception/runtime-overhead-of-using-defer-in-go-7140d5c40e32

// SyncedOrderedMap stores keys in a double linked list and KeyValueHolder as List elements in a hashmap index by those keys
// Thread safe: All read / write operations to the map and list are mutex protected.
type SyncedOrderedMap[K comparable, V any] struct {
	sync.RWMutex
	store map[K]*list.Element
	keys  *list.List
}

// NewSyncedOrderedMap creates a new SyncedOrderedMap for keys of type K and values of type V
func NewSyncedOrderedMap[K comparable, V any]() *SyncedOrderedMap[K, V] {
	return &SyncedOrderedMap[K, V]{
		store: make(map[K]*list.Element),
		keys:  list.New(),
	}
}

// NewSizedSyncedOrderedMap trades memory space for better performance in terms of time complexity
// by over-allocating initially capacity by a multiplier to prevent dynamic map resize.
//
// To create an [int, string] map of a total size of 500 entries, call:
// 	m := collection_types.NewSyncedOrderedMap[int, string](100, 5)
func NewSizedSyncedOrderedMap[K comparable, V any](size, multiple int) *SyncedOrderedMap[K, V] {
	capacity := size * multiple
	return &SyncedOrderedMap[K, V]{
		store: make(map[K]*list.Element, capacity),
		keys:  list.New(),
	}
}

// Set stores a pair of a key type K and a value of type V if it is not yet in the map.
// If the key is already in the map, it replaces the value in the map with the given parameter value.
func (m *SyncedOrderedMap[K, V]) Set(key K, val V) {
	m.Lock()
	var e *list.Element
	_, exists := m.store[key]
	if !exists {
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
	m.Unlock()
}

// Get returns a value V and true  for the given key if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(n) at worst, or 0(1) at best.
func (m *SyncedOrderedMap[K, V]) Get(key K) (value V, ok bool) {
	m.RLock()
	val, exists := m.store[key]
	if !exists {
		m.RUnlock()
		return *new(V), false
	} else {
		value = val.Value.(KeyValueHolder[K, V]).value
		m.RUnlock()
		return value, true
	}
}

// GetFirst returns the first value V and true if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(1).
func (m *SyncedOrderedMap[K, V]) GetFirst() (value V, ok bool) {
	m.RLock()
	k := m.keys.Front().Value.(KeyValueHolder[K, V])
	val, exists := m.store[k.key]
	if !exists {
		m.RUnlock()
		return *new(V), false
	} else {
		value = val.Value.(KeyValueHolder[K, V]).value
		m.RUnlock()
		return value, true
	}
}

// GetLast returns the last value V and true if it exits.
// Otherwise, it returns an empty value and false.
// The complexity is O(1).
func (m *SyncedOrderedMap[K, V]) GetLast() (value V, ok bool) {
	m.RLock()
	k := m.keys.Back().Value.(KeyValueHolder[K, V])
	val, exists := m.store[k.key]
	if !exists {
		m.RUnlock()
		return *new(V), false
	} else {
		value = val.Value.(KeyValueHolder[K, V]).value
		m.RUnlock()
		return value, true
	}
}

// Delete removes the key from the list and the value from the map.
// If the key does not exist, nothing happen.
// The complexity is O(1).
func (m *SyncedOrderedMap[K, V]) Delete(key K) {
	m.Lock()
	e, exists := m.store[key]
	if !exists {
		m.Unlock()
		return
	} else {
		m.keys.Remove(e)
		delete(m.store, key)
		m.Unlock()
		return
	}
}

// Size returns the number of elements stored in the SyncedOrderedMap.
// The complexity is O(1).
func (m *SyncedOrderedMap[K, V]) Size() (size int) {
	m.RLock()
	size = m.keys.Len()
	m.RUnlock()
	return size
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
func (m *SyncedOrderedMap[K, V]) Iterator() func() (*int, *K, V) {
	m.RLock()
	e := m.keys.Front()
	j := 0
	m.RUnlock()
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
