package graph_types

import (
	"strings"
	"testing"
)

func TestGraphDot_empty(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))

	actual := strings.TrimSpace(string(g.Dot(nil)))
	expected := strings.TrimSpace(testGraphDotEmptyStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraphDot_basic(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(3)))

	actual := strings.TrimSpace(string(g.Dot(nil)))
	expected := strings.TrimSpace(testGraphDotBasicStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

type mystr string

func (m mystr) Hashcode() string {
	return string(m)
}
func TestGraphDot_quoted(t *testing.T) {
	var g Graph[mystr]
	quoted := mystr(`name["with-quotes"]`)
	other := mystr(`other`)
	g.Add(quoted)
	g.Add(other)
	g.Connect(NewBasicEdge(quoted, other))

	actual := strings.TrimSpace(string(g.Dot(nil)))
	expected := strings.TrimSpace(testGraphDotQuotedStr)
	if actual != expected {
		t.Fatalf("\ngot:   %q\nwanted %q\n", actual, expected)
	}
}

func TestGraphDot_attrs(t *testing.T) {
	var g Graph[*testGraphNodeDotter]
	g.Add(&testGraphNodeDotter{
		Result: &DotNode{
			Name:  "foo",
			Attrs: map[string]string{"foo": "bar"},
		},
	})

	actual := strings.TrimSpace(string(g.Dot(nil)))
	expected := strings.TrimSpace(testGraphDotAttrsStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

type testGraphNodeDotter struct{ Result *DotNode }

func (n *testGraphNodeDotter) Hashcode() string { return n.Result.Name }

func (n *testGraphNodeDotter) Name() string                      { return n.Result.Name }
func (n *testGraphNodeDotter) DotNode(string, *DotOpts) *DotNode { return n.Result }

const testGraphDotQuotedStr = `digraph {
	compound = "true"
	newrank = "true"
	subgraph "root" {
		"[root] name[\"with-quotes\"]" -> "[root] other"
	}
}`

const testGraphDotBasicStr = `digraph {
	compound = "true"
	newrank = "true"
	subgraph "root" {
		"[root] 1" -> "[root] 3"
	}
}
`

const testGraphDotEmptyStr = `digraph {
	compound = "true"
	newrank = "true"
	subgraph "root" {
	}
}`

const testGraphDotAttrsStr = `digraph {
	compound = "true"
	newrank = "true"
	subgraph "root" {
		"[root] foo" [foo = "bar"]
	}
}`
