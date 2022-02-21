package graph_types

import (
	"fmt"
	"sort"
	"strconv"
)

// the marshal* structs are for serialization of the graph data.
type marshalGraph struct {
	// Type is always "Graph", for identification as a top level object in the
	// JSON stream.
	Type string

	// Each marshal structure requires a unique ID so that it can be referenced
	// by other structures.
	ID string `json:",omitempty"`

	// Human readable name for this graph.
	Name string `json:",omitempty"`

	// Arbitrary attributes that can be added to the output.
	Attrs map[string]string `json:",omitempty"`

	// List of graph vertices, sorted by ID.
	Vertices []*marshalVertex `json:",omitempty"`

	// List of edges, sorted by Source ID.
	Edges []*marshalEdge `json:",omitempty"`

	// Any number of subgraphs. A subgraph itself is considered a vertex, and
	// may be referenced by either end of an edge.
	Subgraphs []*marshalGraph `json:",omitempty"`

	// Any lists of vertices that are included in cycles.
	Cycles [][]*marshalVertex `json:",omitempty"`
}

func (g *marshalGraph) vertexByID(id string) *marshalVertex {
	for _, v := range g.Vertices {
		if id == v.ID {
			return v
		}
	}
	return nil
}

type marshalVertex struct {
	// Unique ID, used to reference this vertex from other structures.
	ID string

	// Human readable name
	Name string `json:",omitempty"`

	Attrs map[string]string `json:",omitempty"`

	// This is to help transition from the old Dot interfaces. We record if the
	// node was a GraphNodeDotter here, so we can call it to get attributes.
	graphNodeDotter GraphNodeDotter
}

func newMarshalVertex[T Hashable](raw T) *marshalVertex {
	var v interface{}
	v = raw
	dn, ok := v.(GraphNodeDotter)
	if !ok {
		dn = nil
	}

	// the name will be quoted again later, so we need to ensure it's properly
	// escaped without quotes.
	name := strconv.Quote(VertexName(raw))
	name = name[1 : len(name)-1]

	return &marshalVertex{
		ID:              marshalVertexID(raw),
		Name:            name,
		Attrs:           make(map[string]string),
		graphNodeDotter: dn,
	}
}

// vertices is a sort.Interface implementation for sorting vertices by ID
type vertices []*marshalVertex

func (v vertices) Less(i, j int) bool { return v[i].Name < v[j].Name }
func (v vertices) Len() int           { return len(v) }
func (v vertices) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }

type marshalEdge struct {
	// Human readable name
	Name string

	// Source and Target Vertices by ID
	Source string
	Target string

	Attrs map[string]string `json:",omitempty"`
}

func newMarshalEdge[T Hashable](e Edge[T]) *marshalEdge {
	return &marshalEdge{
		Name:   fmt.Sprintf("%s|%s", VertexName(e.Source()), VertexName(e.Target())),
		Source: marshalVertexID(e.Source()),
		Target: marshalVertexID(e.Target()),
		Attrs:  make(map[string]string),
	}
}

// edges is a sort.Interface implementation for sorting edges by Source ID
type edges []*marshalEdge

func (e edges) Less(i, j int) bool { return e[i].Name < e[j].Name }
func (e edges) Len() int           { return len(e) }
func (e edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

// build a marshalGraph structure from a *Graph
func newMarshalGraph[T Hashable](name string, g *Graph[T]) *marshalGraph {
	mg := &marshalGraph{
		Type:  "Graph",
		Name:  name,
		Attrs: make(map[string]string),
	}

	for _, v := range g.Vertices() {
		id := marshalVertexID(v)
		if sg, ok := marshalSubgrapher(v); ok {
			smg := newMarshalGraph(VertexName(v), sg)
			smg.ID = id
			mg.Subgraphs = append(mg.Subgraphs, smg)
		}

		mv := newMarshalVertex(v)
		mg.Vertices = append(mg.Vertices, mv)
	}

	sort.Sort(vertices(mg.Vertices))

	for _, e := range g.Edges() {
		mg.Edges = append(mg.Edges, newMarshalEdge(e))
	}

	sort.Sort(edges(mg.Edges))

	for _, c := range (&AcyclicGraph[T]{*g}).Cycles() {
		var cycle []*marshalVertex
		for _, v := range c {
			mv := newMarshalVertex(v)
			cycle = append(cycle, mv)
		}
		mg.Cycles = append(mg.Cycles, cycle)
	}

	return mg
}

// Attempt to return a unique ID for any vertex.
func marshalVertexID[T Hashable](v T) string {

	return v.Hashcode()
}

// check for a Subgrapher, and return the underlying *Graph.
func marshalSubgrapher[T Hashable](raw T) (*Graph[T], bool) {
	var v interface{}
	v = raw
	sg, ok := v.(Subgrapher)
	if !ok {
		return nil, false
	}

	switch g := sg.Subgraph().DirectedGraph().(type) {
	case *Graph[T]:
		return g, true
	case *AcyclicGraph[T]:
		return &g.Graph, true
	}

	return nil, false
}
