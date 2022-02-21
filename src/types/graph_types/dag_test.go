package graph_types

import (
	"flag"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestAcyclicGraphRoot(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(3), myint(2)))
	g.Connect(NewBasicEdge(myint(3), myint(1)))

	if root, err := g.Roots(); err != nil {
		t.Fatalf("err: %s", err)
	} else if root[0] != myint(3) {
		t.Fatalf("bad: %#v", root)
	}
}

func TestAcyclicGraphRoot_cycle(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Connect(NewBasicEdge(myint(3), myint(1)))

	if _, err := g.Roots(); err == nil {
		t.Fatal("should error")
	}
}

func TestAcyclicGraphRoot_multiple(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(3), myint(2)))

	if _, err := g.Roots(); err != nil {
		t.Fatal("should NOT error for multiple roots.")
	}
}

func TestAcyclicGraphTransReduction(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(1), myint(3)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.TransitiveReduction()

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphTransReductionStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestAcyclicGraphTransReduction_more(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Add(myint(4))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(1), myint(3)))
	g.Connect(NewBasicEdge(myint(1), myint(4)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Connect(NewBasicEdge(myint(2), myint(4)))
	g.Connect(NewBasicEdge(myint(3), myint(4)))
	g.TransitiveReduction()

	actual := strings.TrimSpace(g.String())
	expected := strings.TrimSpace(testGraphTransReductionMoreStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

// Make sure we can reduce a sizable, fully-connected graph.
func TestAcyclicGraphTransReduction_fullyConnected(t *testing.T) {
	var g AcyclicGraph[*counter]

	const nodeCount = 200
	nodes := make([]*counter, nodeCount)
	for i := 0; i < nodeCount; i++ {
		nodes[i] = &counter{Name: strconv.Itoa(i)}
	}

	// Add them all to the graph
	for _, n := range nodes {
		g.Add(n)
	}

	// connect them all
	for i := range nodes {
		for j := range nodes {
			if i == j {
				continue
			}
			g.Connect(NewBasicEdge(nodes[i], nodes[j]))
		}
	}

	g.TransitiveReduction()

	vertexNameCalls := int64(0)
	for _, n := range nodes {
		vertexNameCalls += n.Calls
	}

	switch {
	case vertexNameCalls > 2*nodeCount:
		// Make calling it more the 2x per node fatal.
		// If we were sorting this would give us roughly ln(n)(n^3) calls, or
		// >59000000 calls for 200 vertices.
		t.Fatalf("VertexName called %d times", vertexNameCalls)
	case vertexNameCalls > 0:
		// we don't expect any calls, but a change here isn't necessarily fatal
		t.Logf("WARNING: VertexName called %d times", vertexNameCalls)
	}
}

func TestAcyclicGraphValidate(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(3), myint(2)))
	g.Connect(NewBasicEdge(myint(3), myint(1)))

	if err := g.Validate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestAcyclicGraphValidate_cycle(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(3), myint(2)))
	g.Connect(NewBasicEdge(myint(3), myint(1)))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(1)))

	if err := g.Validate(); err == nil {
		t.Fatal("should error")
	}
}

func TestAcyclicGraphValidate_cycleSelf(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Connect(NewBasicEdge(myint(1), myint(1)))

	if err := g.Validate(); err == nil {
		t.Fatal("should error")
	}
}

func TestAcyclicGraphAncestors(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Add(myint(4))
	g.Add(myint(5))
	g.Connect(NewBasicEdge(myint(0), myint(1)))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Connect(NewBasicEdge(myint(3), myint(4)))
	g.Connect(NewBasicEdge(myint(4), myint(5)))

	actual, err := g.Descendents(myint(2))
	if err != nil {
		t.Fatalf("err: %#v", err)
	}

	expected := []myint{myint(3), myint(4), myint(5)}

	if actual.Len() != len(expected) {
		t.Fatalf("bad length! expected %#v to have len %d", actual, len(expected))
	}

	for _, e := range expected {
		if !actual.Include(e) {
			t.Fatalf("expected: %#v to include: %#v", expected, actual)
		}
	}
}

func TestAcyclicGraphDescendents(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Add(myint(4))
	g.Add(myint(5))
	g.Connect(NewBasicEdge(myint(0), myint(1)))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(3)))
	g.Connect(NewBasicEdge(myint(3), myint(4)))
	g.Connect(NewBasicEdge(myint(4), myint(5)))

	actual, err := g.Ancestors(myint(2))
	if err != nil {
		t.Fatalf("err: %#v", err)
	}

	expected := []myint{myint(0), myint(1)}

	if actual.Len() != len(expected) {
		t.Fatalf("bad length! expected %#v to have len %d", actual, len(expected))
	}

	for _, e := range expected {
		if !actual.Include(e) {
			t.Fatalf("expected: %#v to include: %#v", expected, actual)
		}
	}
}

func TestAcyclicGraph_ReverseDepthFirstWalk_WithRemoval(t *testing.T) {
	var g AcyclicGraph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Add(myint(3))
	g.Connect(NewBasicEdge(myint(3), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(1)))

	var visits []myint
	var lock sync.Mutex
	err := g.SortedReverseDepthFirstWalk([]myint{myint(1)}, func(v myint, d int) error {
		lock.Lock()
		defer lock.Unlock()
		visits = append(visits, v)
		g.Remove(v)
		return nil
	})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	expected := []myint{myint(1), myint(2), myint(3)}
	if !reflect.DeepEqual(visits, expected) {
		t.Fatalf("expected: %#v, got: %#v", expected, visits)
	}
}

const testGraphTransReductionStr = `
1
  2
2
  3
3
`

const testGraphTransReductionMoreStr = `
1
  2
2
  3
3
  4
4
`
