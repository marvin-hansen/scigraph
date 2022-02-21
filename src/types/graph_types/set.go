package graph_types

// Set is a set data structure.
type Set[T Hashable] map[string]T

// Add adds an item to the set
func (s Set[T]) Add(v T) {
	s[v.Hashcode()] = v
}

// Delete removes an item from the set.
func (s Set[T]) Delete(v T) {
	delete(s, v.Hashcode())
}

// Include returns true/false of whether a value is in the set.
func (s Set[T]) Include(v T) bool {
	_, ok := s[v.Hashcode()]
	return ok
}

// Intersection computes the set intersection with other.
func (s Set[T]) Intersection(other Set[T]) Set[T] {
	result := make(Set[T])
	if s == nil || other == nil {
		return result
	}
	// Iteration over a smaller set has better performance.
	if other.Len() < s.Len() {
		s, other = other, s
	}
	for _, v := range s {
		if other.Include(v) {
			result.Add(v)
		}
	}
	return result
}

// Difference returns a set with the elements that s has but
// other doesn't.
func (s Set[T]) Difference(other Set[T]) Set[T] {
	if other == nil || other.Len() == 0 {
		return s.Copy()
	}

	result := make(Set[T])
	for k, v := range s {
		if _, ok := other[k]; !ok {
			result.Add(v)
		}
	}

	return result
}

// Filter returns a set that contains the elements from the receiver
// where the given callback returns true.
func (s Set[T]) Filter(cb func(T) bool) Set[T] {
	result := make(Set[T])

	for _, v := range s {
		if cb(v) {
			result.Add(v)
		}
	}

	return result
}

// Len is the number of items in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// List returns the list of set elements.
func (s Set[T]) List() []T {
	if s == nil {
		return nil
	}

	r := make([]T, 0, len(s))
	for _, v := range s {
		r = append(r, v)
	}

	return r
}

// Copy returns a shallow copy of the set.
func (s Set[T]) Copy() Set[T] {
	c := make(Set[T], len(s))
	for k, v := range s {
		c[k] = v
	}
	return c
}
