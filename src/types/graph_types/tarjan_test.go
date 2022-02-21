package graph_types

import (
	"sort"
	"strings"
	"testing"
)

func TestGraphStronglyConnected(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(1)))

	actual := strings.TrimSpace(testSCCStr(StronglyConnected(&g)))
	expected := strings.TrimSpace(testGraphStronglyConnectedStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraphStronglyConnected_two(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(1)))
	g.Add(myint(3))

	actual := strings.TrimSpace(testSCCStr(StronglyConnected(&g)))
	expected := strings.TrimSpace(testGraphStronglyConnectedTwoStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func TestGraphStronglyConnected_three(t *testing.T) {
	var g Graph[myint]
	g.Add(myint(1))
	g.Add(myint(2))
	g.Connect(NewBasicEdge(myint(1), myint(2)))
	g.Connect(NewBasicEdge(myint(2), myint(1)))
	g.Add(myint(3))
	g.Add(myint(4))
	g.Add(myint(5))
	g.Add(myint(6))
	g.Connect(NewBasicEdge(myint(4), myint(5)))
	g.Connect(NewBasicEdge(myint(5), myint(6)))
	g.Connect(NewBasicEdge(myint(6), myint(4)))

	actual := strings.TrimSpace(testSCCStr(StronglyConnected(&g)))
	expected := strings.TrimSpace(testGraphStronglyConnectedThreeStr)
	if actual != expected {
		t.Fatalf("bad: %s", actual)
	}
}

func testSCCStr[T Hashable](list [][]T) string {
	var lines []string
	for _, vs := range list {
		result := make([]string, len(vs))
		for i, v := range vs {
			result[i] = VertexName(v)
		}

		sort.Strings(result)
		lines = append(lines, strings.Join(result, ","))
	}

	sort.Strings(lines)
	return strings.Join(lines, "\n")
}

const testGraphStronglyConnectedStr = `1,2`

const testGraphStronglyConnectedTwoStr = `
1,2
3
`

const testGraphStronglyConnectedThreeStr = `
1,2
3
4,5,6
`
