package graph_types

import (
	"strings"
	"testing"
)

func TestGraph_empty(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphEmptyStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraph_basic(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(3)))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphBasicStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraph_remove(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(3)))
	g.Remove(myint(3))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphRemoveStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraph_replace(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Replace(myint(2), myint(42))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphReplaceStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraph_replaceSelf(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Replace(myint(2), myint(2))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphReplaceSelfStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

// This tests that connecting edges works based on custom Hashcode implementations for uniqueness.
func TestGraph_hashcode(t *testing.T) {
	var g Graph[*HashVertex]
	g.Add(&HashVertex{code: 1})
	g.Add(&HashVertex{code: 2})
	g.Add(&HashVertex{code: 3})
	g.Connect(NewBasicEdge(
		&HashVertex{code: 1},
		&HashVertex{code: 3}))

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphBasicStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraphHasVertex(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))

	if !g.HasVertex(myint(1)) {
		t.Fatal("should have 1")
	}
	if g.HasVertex(myint(2)) {
		t.Fatal("should not have 2")
	}
}

func TestGraphHasEdge(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Connect(NewBasicEdge(myint(1), myint(2)))

	if !g.HasEdge(NewBasicEdge(myint(1), myint(2))) {
		t.Fatal("should have 1,2")
	}
	if g.HasEdge(NewBasicEdge(myint(2), myint(3))) {
		t.Fatal("should not have 2,3")
	}
}

func TestGraphEdgesFrom(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(3)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))

	edges := g.EdgesFrom(myint(1))

	expected := make(Set[Edge[myint]])
	expected.Add(NewBasicEdge(myint(1), myint(3)))

	s := make(Set[Edge[myint]])
	for _, e := range edges {
		s.Add(e)
	}

	if s.Intersection(expected).Len() != expected.Len() {
		t.Fatalf("bad: %#v", edges)
	}
}

func TestGraphEdgesTo(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(3)))
	g.Connect(NewBasicEdge(myint(1), myint(2)))

	edges := g.EdgesTo(myint(3))

	expected := make(Set[Edge[myint]])
	expected.Add(NewBasicEdge(myint(1), myint(3)))

	s := make(Set[Edge[myint]])
	for _, e := range edges {
		s.Add(e)
	}

	if s.Intersection(expected).Len() != expected.Len() {
		t.Fatalf("bad: %#v", edges)
	}
}

func TestGraphUpdownEdges(t *testing.T) {
	// Verify that we can't inadvertently modify the internal graph sets
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))

	up := g.UpEdges(myint(2))
	if up.Len() != 1 || !up.Include(myint(1)) {
		t.Fatalf("expected only an up edge of '1', got %#v", up)
	}
	// modify the up set
	up.Add(myint(9))

	orig := g.UpEdges(myint(2))
	diff := up.Difference(orig)
	if diff.Len() != 1 || !diff.Include(myint(9)) {
		t.Fatalf("expected a diff of only '9', got %#v", diff)
	}

	down := g.DownEdges(myint(2))
	if down.Len() != 1 || !down.Include(myint(3)) {
		t.Fatalf("expected only a down edge of '3', got %#v", down)
	}
	// modify the down set
	down.Add(myint(8))

	orig = g.DownEdges(myint(2))
	diff = down.Difference(orig)
	if diff.Len() != 1 || !diff.Include(myint(8)) {
		t.Fatalf("expected a diff of only '8', got %#v", diff)
	}
}

const testGraphBasicStr = `
1
  3
2
3
`

const testGraphEmptyStr = `
1
2
3
`

const testGraphRemoveStr = `
1
2
`

const testGraphReplaceStr = `
1
  42
3
42
  3
`

const testGraphReplaceSelfStr = `
1
  2
2
  3
3
`
