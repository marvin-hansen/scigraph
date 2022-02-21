package graph_types

import (
	"bytes"
	"fmt"
	"sort"
)

// Graph is used to represent a dependency graph.
type Graph[T Hashable] struct {
	vertices  Set[T]
	edges     Set[Edge[T]]
	downEdges map[string]Set[T]
	upEdges   map[string]Set[T]
}

func (g *Graph[T]) DirectedGraph() Grapher {
	return g
}

// Vertices returns the list of all the vertices in the graph.
func (g *Graph[T]) Vertices() []T {
	result := make([]T, 0, len(g.vertices))
	for _, v := range g.vertices {
		result = append(result, v)
	}

	return result
}

// Edges returns the list of all the edges in the graph.
func (g *Graph[T]) Edges() []Edge[T] {
	result := make([]Edge[T], 0, len(g.edges))
	for _, v := range g.edges {
		result = append(result, v)
	}

	return result
}

// EdgesFrom returns the list of edges from the given source.
func (g *Graph[T]) EdgesFrom(v T) []Edge[T] {
	var result []Edge[T]
	from := v.Hashcode()
	for _, e := range g.Edges() {
		if e.Source().Hashcode() == from {
			result = append(result, e)
		}
	}

	return result
}

// EdgesTo returns the list of edges to the given target.
func (g *Graph[T]) EdgesTo(v T) []Edge[T] {
	var result []Edge[T]
	search := v.Hashcode()
	for _, e := range g.Edges() {
		if e.Target().Hashcode() == search {
			result = append(result, e)
		}
	}

	return result
}

// HasVertex checks if the given Vertex is present in the graph.
func (g *Graph[T]) HasVertex(v T) bool {
	return g.vertices.Include(v)
}

// HasEdge checks if the given Edge is present in the graph.
func (g *Graph[T]) HasEdge(e Edge[T]) bool {
	return g.edges.Include(e)
}

// Add adds a vertex to the graph. This is safe to call multiple time with
// the same Vertex.
func (g *Graph[T]) Add(v T) T {
	g.init()
	g.vertices.Add(v)
	return v
}

// Remove removes a vertex from the graph. This will also remove any
// edges with this vertex as a source or target.
func (g *Graph[T]) Remove(v T) T {
	// Delete the vertex itself
	g.vertices.Delete(v)

	// Delete the edges to non-existent things
	for _, target := range g.downEdgesNoCopy(v) {
		g.RemoveEdge(NewBasicEdge(v, target))
	}
	for _, source := range g.upEdgesNoCopy(v) {
		g.RemoveEdge(NewBasicEdge(source, v))
	}
	var new T
	return new

}

// Replace replaces the original Vertex with replacement. If the original
// does not exist within the graph, then false is returned. Otherwise, true
// is returned.
func (g *Graph[T]) Replace(original, replacement T) bool {
	// If we don't have the original, we can't do anything
	if !g.vertices.Include(original) {
		return false
	}

	// If they're the same, then don't do anything
	if original.Hashcode() == replacement.Hashcode() {
		return true
	}

	// Add our new vertex, then copy all the edges
	g.Add(replacement)
	for _, target := range g.downEdgesNoCopy(original) {
		g.Connect(NewBasicEdge(replacement, target))
	}
	for _, source := range g.upEdgesNoCopy(original) {
		g.Connect(NewBasicEdge(source, replacement))
	}

	// Remove our old vertex, which will also remove all the edges
	g.Remove(original)

	return true
}

// RemoveEdge removes an edge from the graph.
func (g *Graph[T]) RemoveEdge(edge Edge[T]) {
	g.init()

	// Delete the edge from the set
	g.edges.Delete(edge)

	// Delete the up/down edges
	if s, ok := g.downEdges[edge.Source().Hashcode()]; ok {
		s.Delete(edge.Target())
	}
	if s, ok := g.upEdges[edge.Target().Hashcode()]; ok {
		s.Delete(edge.Source())
	}
}

// UpEdges returns the vertices connected to the outward edges from the source
// Vertex v.
func (g *Graph[T]) UpEdges(v T) Set[T] {
	return g.upEdgesNoCopy(v).Copy()
}

// DownEdges returns the vertices connected from the inward edges to Vertex v.
func (g *Graph[T]) DownEdges(v T) Set[T] {
	return g.downEdgesNoCopy(v).Copy()
}

// downEdgesNoCopy returns the outward edges from the source Vertex v as a Set.
// This Set is the same as used internally bu the Graph to prevent a copy, and
// must not be modified by the caller.
func (g *Graph[T]) downEdgesNoCopy(v T) Set[T] {
	g.init()
	return g.downEdges[v.Hashcode()]
}

// upEdgesNoCopy returns the inward edges to the destination Vertex v as a Set.
// This Set is the same as used internally bu the Graph to prevent a copy, and
// must not be modified by the caller.
func (g *Graph[T]) upEdgesNoCopy(v T) Set[T] {
	g.init()
	return g.upEdges[v.Hashcode()]
}

// Connect adds an edge with the given source and target. This is safe to
// call multiple times with the same value. Note that the same value is
// verified through pointer equality of the vertices, not through the
// value of the edge itself.
func (g *Graph[T]) Connect(edge Edge[T]) {
	g.init()

	source := edge.Source()
	target := edge.Target()
	sourceCode := source.Hashcode()
	targetCode := target.Hashcode()

	// Do we have this already? If so, don't add it again.
	if s, ok := g.downEdges[sourceCode]; ok && s.Include(target) {
		return
	}

	// Add the edge to the set
	g.edges.Add(edge)

	// Add the down edge
	s, ok := g.downEdges[sourceCode]
	if !ok {
		s = make(Set[T])
		g.downEdges[sourceCode] = s
	}
	s.Add(target)

	// Add the up edge
	s, ok = g.upEdges[targetCode]
	if !ok {
		s = make(Set[T])
		g.upEdges[targetCode] = s
	}
	s.Add(source)
}

// String outputs some human-friendly output for the graph structure.
func (g *Graph[T]) StringWithNodeTypes() string {
	var buf bytes.Buffer

	// Build the list of node names and a mapping so that we can more
	// easily alphabetize the output to remain deterministic.
	vertices := g.Vertices()
	names := make([]string, 0, len(vertices))
	mapping := make(map[string]T, len(vertices))
	for _, v := range vertices {
		name := VertexName(v)
		names = append(names, name)
		mapping[name] = v
	}
	sort.Strings(names)

	// Write each node in order...
	for _, name := range names {
		v := mapping[name]
		targets := g.downEdges[v.Hashcode()]

		buf.WriteString(fmt.Sprintf("%s - %T\n", name, v))

		// Alphabetize dependencies
		deps := make([]string, 0, targets.Len())
		targetNodes := make(map[string]T)
		for _, target := range targets {
			dep := VertexName(target)
			deps = append(deps, dep)
			targetNodes[dep] = target
		}
		sort.Strings(deps)

		// Write dependencies
		for _, d := range deps {
			buf.WriteString(fmt.Sprintf("  %s - %T\n", d, targetNodes[d]))
		}
	}

	return buf.String()
}

// String outputs some human-friendly output for the graph structure.
func (g *Graph[T]) String() string {
	var buf bytes.Buffer

	// Build the list of node names and a mapping so that we can more
	// easily alphabetize the output to remain deterministic.
	vertices := g.Vertices()
	names := make([]string, 0, len(vertices))
	mapping := make(map[string]T, len(vertices))
	for _, v := range vertices {
		name := VertexName(v)
		names = append(names, name)
		mapping[name] = v
	}
	sort.Strings(names)

	// Write each node in order...
	for _, name := range names {
		v := mapping[name]
		targets := g.downEdges[v.Hashcode()]

		buf.WriteString(fmt.Sprintf("%s\n", name))

		// Alphabetize dependencies
		deps := make([]string, 0, targets.Len())
		for _, target := range targets {
			deps = append(deps, VertexName(target))
		}
		sort.Strings(deps)

		// Write dependencies
		for _, d := range deps {
			buf.WriteString(fmt.Sprintf("  %s\n", d))
		}
	}

	return buf.String()
}

func (g *Graph[T]) init() {
	if g.vertices == nil {
		g.vertices = make(Set[T])
	}
	if g.edges == nil {
		g.edges = make(Set[Edge[T]])
	}
	if g.downEdges == nil {
		g.downEdges = make(map[string]Set[T])
	}
	if g.upEdges == nil {
		g.upEdges = make(map[string]Set[T])
	}
}

// Dot returns a dot-formatted representation of the Graph.
func (g *Graph[T]) Dot(opts *DotOpts) []byte {
	return newMarshalGraph("", g).Dot(opts)
}

// VertexName returns the name of a vertex.
func VertexName[T Hashable](raw T) string {

	var l interface{}
	l = raw
	n, ok := l.(Named)
	if ok {
		return n.Name()
	}

	return raw.Hashcode()
}
