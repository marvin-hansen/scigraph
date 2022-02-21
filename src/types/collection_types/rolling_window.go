// Copyright (c) 2021. Marvin Friedrich Lars Hansen. All Rights Reserved. Contact: marvin.hansen@gmail.com

package collection_types

import "log"

type RollingWindow[V any] struct {
	arr  []V
	size int
	head int
	tail int
}

// NewRollingWindow creates a new moving window,
// with the size and multiple specified.
//
// This data structures trades off space and copying complexity; more precisely,
// the number of moving windows that can be displayed without having to do any
// array copying is proportional to approx 1/M, where M is the multiple.
func NewRollingWindow[V any](size, multiple int) *RollingWindow[V] {
	if size < 1 || multiple < 1 {
		log.Fatalf("Must have positive size and multiple")
	}
	capacity := size * multiple
	return &RollingWindow[V]{
		arr:  make([]V, capacity, capacity),
		size: size,
	}
}

// Filled returns true if the initial size has been reached
func (m *RollingWindow[V]) Filled() bool {
	if m.tail < m.size {
		return false
	} else {
		return true
	}
}
func (m *RollingWindow[V]) PushBack(v V) {
	// if the array is full, rewind
	if m.tail == len(m.arr) {
		m.rewind()
	}
	// push the value
	m.arr[m.tail] = v
	// check if the window is full,
	// and move head pointer appropriately
	if m.tail-m.head >= m.size {
		m.head++
	}
	m.tail++
}

func (m *RollingWindow[V]) rewind() {
	l := len(m.arr)
	for i := 0; i < m.size-1; i++ {
		m.arr[i] = m.arr[l-m.size+i+1]
	}
	m.head, m.tail = 0, m.size-1
}

// GetSlice Slice will present the MovingWindow in
// the form of a slice. This operation never
// requires array copying of any kind.
//
func (m *RollingWindow[V]) GetSlice() []V {
	return m.arr[m.head:m.tail]
}

func (m *RollingWindow[V]) GetReverseSlice() []V {

	s := m.arr[m.head:m.tail]
	for i, j := 0, len(m.arr[m.head:m.tail])-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Size returns the size of the moving window,
// which is set at initialization
func (m *RollingWindow[V]) Size() int {
	return m.size
}
