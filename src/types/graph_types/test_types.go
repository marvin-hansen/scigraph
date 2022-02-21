package graph_types

import (
	"fmt"
	"strconv"
)

type myint int

func (m myint) Hashcode() string {
	return strconv.Itoa(int(m))
}

type HashVertex struct {
	code interface{}
}

func (v *HashVertex) Hashcode() string {
	return fmt.Sprintf("%v", v.code)
}

func (v *HashVertex) Name() string {
	return fmt.Sprintf("%#v", v.code)
}

// use this to simulate slow sort operations
type counter struct {
	Name  string
	Calls int64
}

func (s *counter) Hashcode() string {
	return s.Name
}

func (s *counter) String() string {
	s.Calls++
	return s.Name
}

// NewBasicEdge returns an Edge implementation that simply tracks the source and target given as-is.
func NewBasicEdge[T Hashable](source, target T) Edge[T] {
	return &BasicEdge[T]{Src: source, Trgt: target}
}

// BasicEdge is a basic implementation of Edge that has the source and  target vertex.
type BasicEdge[T Hashable] struct {
	Src  T
	Trgt T
	Hashable
}

func (e BasicEdge[T]) Hashcode() string {
	return e.Src.Hashcode() + "-" + e.Trgt.Hashcode()
}

func (e BasicEdge[T]) Source() T {
	return e.Src
}

func (e BasicEdge[T]) Target() T {
	return e.Trgt
}
