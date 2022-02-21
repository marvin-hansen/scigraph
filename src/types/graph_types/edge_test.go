package graph_types

import (
	"testing"
)

func TestNewBasicEdgeHashcode(t *testing.T) {
	e1 := NewBasicEdge(myint(1), myint(2))
	e2 := NewBasicEdge(myint(1), myint(2))
	if e1.Hashcode() != e2.Hashcode() {
		t.Fatalf("bad")
	}
}

type test struct {
	Value string
}

func (t test) Hashcode() string {
	return t.Value
}
func TestNewBasicEdgeHashcode_pointer(t *testing.T) {

	v1, v2 := &test{"foo"}, &test{"bar"}
	e1 := NewBasicEdge(v1, v2)
	e2 := NewBasicEdge(v1, v2)
	if e1.Hashcode() != e2.Hashcode() {
		t.Fatalf("bad")
	}
}
