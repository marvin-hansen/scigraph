package graph_types

// StronglyConnected returns the list of strongly connected components
// within the Graph g. This information is primarily used by this package
// for cycle detection, but strongly connected components have widespread use.
func StronglyConnected[T Hashable](g *Graph[T]) [][]T {
	vs := g.Vertices()
	acct := newSccAcct[T](len(vs))
	for _, v := range vs {
		// Recurse on any non-visited nodes
		if acct.VertexIndex[v.Hashcode()] == 0 {
			stronglyConnected(acct, g, v)
		}
	}
	return acct.SCC
}

func stronglyConnected[T Hashable](acct *sccAcct[T], g *Graph[T], v T) int {
	// Initial vertex visit
	index := acct.visit(v)
	minIdx := index

	for _, raw := range g.downEdgesNoCopy(v) {
		target := raw
		targetIdx := acct.VertexIndex[target.Hashcode()]

		// Recurse on successor if not yet visited
		if targetIdx == 0 {
			minIdx = min(minIdx, stronglyConnected(acct, g, target))
		} else if acct.inStack(target) {
			// Check if the vertex is in the stack
			minIdx = min(minIdx, targetIdx)
		}
	}

	// Pop the strongly connected components off the stack if
	// this is a root vertex
	if index == minIdx {
		var scc []T
		for {
			v2 := acct.pop()
			scc = append(scc, v2)
			if v2.Hashcode() == v.Hashcode() {
				break
			}
		}

		acct.SCC = append(acct.SCC, scc)
	}

	return minIdx
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// sccAcct is used ot pass around accounting information for the StronglyConnectedComponents algorithm
type sccAcct[T Hashable] struct {
	NextIndex   int
	VertexIndex map[string]int
	Stack       []T
	SCC         [][]T
}

func newSccAcct[T Hashable](len int) *sccAcct[T] {
	return &sccAcct[T]{
		NextIndex:   1,
		VertexIndex: make(map[string]int, len),
	}

}

// visit assigns an index and pushes a vertex onto the stack
func (s *sccAcct[T]) visit(v T) int {
	idx := s.NextIndex
	s.VertexIndex[v.Hashcode()] = idx
	s.NextIndex++
	s.push(v)
	return idx
}

// push adds a vertex to the stack
func (s *sccAcct[T]) push(n T) {
	s.Stack = append(s.Stack, n)
}

// pop removes a vertex from the stack
func (s *sccAcct[T]) pop() T {
	n := len(s.Stack)
	if n == 0 {
		var new T
		return new

	}
	vertex := s.Stack[n-1]
	s.Stack = s.Stack[:n-1]
	return vertex
}

// inStack checks if a vertex is in the stack
func (s *sccAcct[T]) inStack(needle T) bool {
	for _, n := range s.Stack {
		if n.Hashcode() == needle.Hashcode() {
			return true
		}
	}
	return false
}
