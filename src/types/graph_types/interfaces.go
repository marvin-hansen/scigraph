package graph_types

// Edge represents an edge in the graph, with a source and target vertex.
type Edge[T Hashable] interface {
	Source() T
	Target() T
	Hashcode() string
}

// A Grapher is any type that returns a Grapher, mainly used to identify
// dag.Graph and dag.AcyclicGraph.  In the case of Graph and AcyclicGraph, they
// return themselves.
type Grapher interface {
	DirectedGraph() Grapher
}

// Subgrapher allows a Vertex to be a Graph itself, by returning a Grapher.
type Subgrapher interface {
	Subgraph() Grapher
}

// Hashable is the interface used by set to get the hash code of a value.
// If this isn't given, then the value of the item being added to the set
// itself is used as the comparison value.
type Hashable interface {
	Hashcode() string
}

// Vertex of the graph.
// type Vertex[T Hashable] Hashable

// Named is an optional interface that can be implemented by Vertex
// to give it a human-friendly name that is used for outputting the graph.
type Named interface {
	Name() string
}
