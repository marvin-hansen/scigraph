package graph_types

import (
	"fmt"
	"github.com/hashicorp/go-multierror"
	"sort"
	"strings"
)

// AcyclicGraph is a specialization of Graph that cannot have cycles.
type AcyclicGraph[T Hashable] struct {
	Graph[T]
}

// DepthWalkFunc is a walk function that also receives the current depth of the
// walk as an argument
type DepthWalkFunc[T Hashable] func(T, int) error

func (g *AcyclicGraph[T]) DirectedGraph() Grapher {
	return g
}

// Descendents returns a Set that includes every Vertex yielded by walking down from the
// provided starting Vertex v. Descendents will NOT include root vertexes that can be reached
// by walking up from v.
func (g *AcyclicGraph[T]) Descendents(v T) (Set[T], error) {
	s := make(Set[T])
	memoFunc := func(v T, d int) error {
		s.Add(v)
		return nil
	}

	if err := g.DepthFirstWalk(g.downEdgesNoCopy(v), memoFunc); err != nil {
		return nil, err
	}

	return s, nil
}

// Ancestors returns a Set that includes every Vertex yielded by walking up from the
// provided starting Vertex v. Ancestors will include all root vertexes that can be reached
// by walking up from v.
func (g *AcyclicGraph[T]) Ancestors(v T) (Set[T], error) {
	s := make(Set[T])
	memoFunc := func(v T, d int) error {
		s.Add(v)
		return nil
	}

	if err := g.ReverseDepthFirstWalk(g.upEdgesNoCopy(v), memoFunc); err != nil {
		return nil, err
	}

	return s, nil
}

// Roots returns the root of the DAG, or an error.
//
// Complexity: O(V)
func (g *AcyclicGraph[T]) Roots() ([]T, error) {
	roots := make([]T, 0, 1)
	for _, v := range g.Vertices() {
		if g.upEdgesNoCopy(v).Len() == 0 {
			roots = append(roots, v)
		}
	}
	if len(roots) == 0 {
		return []T{}, fmt.Errorf("no roots found")
	}

	return roots, nil
}

// TransitiveReduction performs the transitive reduction of graph g in place.
// The transitive reduction of a graph is a graph with as few edges as
// possible with the same reachability as the original graph. This means
// that if there are three nodes A => B => C, and A connects to both
// B and C, and B connects to C, then the transitive reduction is the
// same graph with only a single edge between A and B, and a single edge
// between B and C.
//
// The graph must be valid for this operation to behave properly. If
// Validate() returns an error, the behavior is undefined and the results
// will likely be unexpected.
//
// Complexity: O(V(V+E)), or asymptotically O(VE)
func (g *AcyclicGraph[T]) TransitiveReduction() {
	// For each vertex u in graph g, do a DFS starting from each vertex
	// v such that the edge (u,v) exists (v is a direct descendant of u).
	//
	// For each v-prime reachable from v, remove the edge (u, v-prime).
	for _, u := range g.Vertices() {
		uTargets := g.downEdgesNoCopy(u)

		g.DepthFirstWalk(g.downEdgesNoCopy(u), func(v T, d int) error {
			shared := uTargets.Intersection(g.downEdgesNoCopy(v))
			for _, vPrime := range shared {
				g.RemoveEdge(NewBasicEdge(u, vPrime))
			}

			return nil
		})
	}
}

// Validate validates the DAG. A DAG is valid if it has at least one root
// and no cycles.
func (g *AcyclicGraph[T]) Validate() error {
	if _, err := g.Roots(); err != nil {
		return err
	}

	// Look for cycles of more than 1 component
	var err error
	cycles := g.Cycles()
	if len(cycles) > 0 {
		for _, cycle := range cycles {
			cycleStr := make([]string, len(cycle))
			for j, vertex := range cycle {
				cycleStr[j] = VertexName(vertex)
			}

			err = multierror.Append(err, fmt.Errorf(
				"Cycle: %s", strings.Join(cycleStr, ", ")))
		}
	}

	// Look for cycles to self
	for _, e := range g.Edges() {
		if e.Source().Hashcode() == e.Target().Hashcode() {
			err = multierror.Append(err, fmt.Errorf(
				"Self reference: %s", VertexName(e.Source())))
		}
	}

	return err
}

func (g *AcyclicGraph[T]) Cycles() [][]T {
	var cycles [][]T
	for _, cycle := range StronglyConnected(&g.Graph) {
		if len(cycle) > 1 {
			cycles = append(cycles, cycle)
		}
	}
	return cycles
}

// // Walk walks the graph, calling your callback as each node is visited.
// // This will walk nodes in parallel if it can. The resulting diagnostics
// // contains problems from all graphs visited, in no particular order.
// func (g *AcyclicGraph[T]) Walk(cb WalkFunc[T]) {
// 	w := &Walker[T]{Callback: cb, Reverse: true}
// 	w.Update(g)
// 	w.Wait()

// }

// simple convenience helper for converting a dag.Set to a []Vertex
func AsVertexList[T Hashable](s Set[T]) []T {
	vertexList := make([]T, 0, len(s))
	for _, raw := range s {
		vertexList = append(vertexList, raw)
	}
	return vertexList
}

type vertexAtDepth[T Hashable] struct {
	Vertex T
	Depth  int
}

// DepthFirstWalk does a depth-first walk of the graph starting from
// the vertices in start.
func (g *AcyclicGraph[T]) DepthFirstWalk(start Set[T], f DepthWalkFunc[T]) error {
	seen := make(map[string]struct{})
	frontier := make([]*vertexAtDepth[T], 0, len(start))
	for _, v := range start {
		frontier = append(frontier, &vertexAtDepth[T]{
			Vertex: v,
			Depth:  0,
		})
	}
	for len(frontier) > 0 {
		// Pop the current vertex
		n := len(frontier)
		current := frontier[n-1]
		frontier = frontier[:n-1]

		// Check if we've seen this already and return...
		if _, ok := seen[current.Vertex.Hashcode()]; ok {
			continue
		}
		seen[current.Vertex.Hashcode()] = struct{}{}

		// Visit the current node
		if err := f(current.Vertex, current.Depth); err != nil {
			return err
		}

		for _, v := range g.downEdgesNoCopy(current.Vertex) {
			frontier = append(frontier, &vertexAtDepth[T]{
				Vertex: v,
				Depth:  current.Depth + 1,
			})
		}
	}

	return nil
}

// SortedDepthFirstWalk does a depth-first walk of the graph starting from
// the vertices in start, always iterating the nodes in a consistent order.
func (g *AcyclicGraph[T]) SortedDepthFirstWalk(start []T, f DepthWalkFunc[T]) error {
	seen := make(map[string]struct{})
	frontier := make([]*vertexAtDepth[T], len(start))
	for i, v := range start {
		frontier[i] = &vertexAtDepth[T]{
			Vertex: v,
			Depth:  0,
		}
	}
	for len(frontier) > 0 {
		// Pop the current vertex
		n := len(frontier)
		current := frontier[n-1]
		frontier = frontier[:n-1]

		// Check if we've seen this already and return...
		if _, ok := seen[current.Vertex.Hashcode()]; ok {
			continue
		}
		seen[current.Vertex.Hashcode()] = struct{}{}

		// Visit the current node
		if err := f(current.Vertex, current.Depth); err != nil {
			return err
		}

		// Visit targets of this in a consistent order.
		targets := AsVertexList(g.downEdgesNoCopy(current.Vertex))
		sort.Sort(byVertexName[T](targets))

		for _, t := range targets {
			frontier = append(frontier, &vertexAtDepth[T]{
				Vertex: t,
				Depth:  current.Depth + 1,
			})
		}
	}

	return nil
}

// ReverseDepthFirstWalk does a depth-first walk _up_ the graph starting from
// the vertices in start.
func (g *AcyclicGraph[T]) ReverseDepthFirstWalk(start Set[T], f DepthWalkFunc[T]) error {
	seen := make(map[string]struct{})
	frontier := make([]*vertexAtDepth[T], 0, len(start))
	for _, v := range start {
		frontier = append(frontier, &vertexAtDepth[T]{
			Vertex: v,
			Depth:  0,
		})
	}
	for len(frontier) > 0 {
		// Pop the current vertex
		n := len(frontier)
		current := frontier[n-1]
		frontier = frontier[:n-1]

		// Check if we've seen this already and return...
		if _, ok := seen[current.Vertex.Hashcode()]; ok {
			continue
		}
		seen[current.Vertex.Hashcode()] = struct{}{}

		for _, t := range g.upEdgesNoCopy(current.Vertex) {
			frontier = append(frontier, &vertexAtDepth[T]{
				Vertex: t,
				Depth:  current.Depth + 1,
			})
		}

		// Visit the current node
		if err := f(current.Vertex, current.Depth); err != nil {
			return err
		}
	}

	return nil
}

// SortedReverseDepthFirstWalk does a depth-first walk _up_ the graph starting from
// the vertices in start, always iterating the nodes in a consistent order.
func (g *AcyclicGraph[T]) SortedReverseDepthFirstWalk(start []T, f DepthWalkFunc[T]) error {
	seen := make(map[string]struct{})
	frontier := make([]*vertexAtDepth[T], len(start))
	for i, v := range start {
		frontier[i] = &vertexAtDepth[T]{
			Vertex: v,
			Depth:  0,
		}
	}
	for len(frontier) > 0 {
		// Pop the current vertex
		n := len(frontier)
		current := frontier[n-1]
		frontier = frontier[:n-1]

		// Check if we've seen this already and return...
		if _, ok := seen[current.Vertex.Hashcode()]; ok {
			continue
		}
		seen[current.Vertex.Hashcode()] = struct{}{}

		// Add next set of targets in a consistent order.
		targets := AsVertexList(g.upEdgesNoCopy(current.Vertex))
		sort.Sort(byVertexName[T](targets))
		for _, t := range targets {
			frontier = append(frontier, &vertexAtDepth[T]{
				Vertex: t,
				Depth:  current.Depth + 1,
			})
		}

		// Visit the current node
		if err := f(current.Vertex, current.Depth); err != nil {
			return err
		}
	}

	return nil
}

// byVertexName implements sort.Interface so a list of Vertices can be sorted
// consistently by their VertexName
type byVertexName[T Hashable] []T

func (b byVertexName[T]) Len() int      { return len(b) }
func (b byVertexName[T]) Swap(i, j int) { b[i], b[j] = b[j], b[i] }
func (b byVertexName[T]) Less(i, j int) bool {
	return VertexName(b[i]) < VertexName(b[j])
}
