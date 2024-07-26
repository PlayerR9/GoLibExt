package WeightedGraph

import (
	uc "github.com/PlayerR9/MyGoLib/Units/common"
	tn "github.com/PlayerR9/tree"
	trb "github.com/PlayerR9/tree/builder"
	tr "github.com/PlayerR9/tree/tree"
)

// WeightFunc is a function that calculates the weight of an edge.
//
// Parameters:
//   - from: the source vertex.
//   - to: the destination vertex.
//
// Returns:
//   - float64: the weight of the edge.
//   - bool: true if the edge exists, otherwise false.
type WeightFunc[T uc.Objecter] func(from, to T) (float64, bool)

// Graph represents a graph.
type Graph[T uc.Objecter] struct {
	// vertices in the graph.
	vertices []T

	// edges in the graph.
	edges [][]*float64
}

// NewGraph creates a new graph with the given vertices.
//
// Parameters:
//   - vertices: vertices in the graph.
//
// Returns:
//   - *WeightedGraph: the new graph.
func NewGraph[T uc.Objecter](vertices []T, f WeightFunc[T]) *Graph[T] {
	if len(vertices) == 0 {
		return &Graph[T]{
			vertices: make([]T, 0),
			edges:    make([][]*float64, 0),
		}
	}

	g := &Graph[T]{
		vertices: vertices,
		edges:    make([][]*float64, 0, len(vertices)),
	}

	for _, from := range vertices {
		edge := make([]*float64, 0, len(vertices))

		for _, to := range vertices {
			w, ok := f(from, to)
			if !ok {
				edge = append(edge, nil)
			} else {
				edge = append(edge, &w)
			}
		}

		g.edges = append(g.edges, edge)
	}

	return g
}

// IndexOf returns the index of the given element in the graph.
//
// Parameters:
//   - elem: the element to find.
//
// Returns:
//   - int: the index of the element, or -1 if not found.
func (g *Graph[T]) IndexOf(elem T) int {
	for i, x := range g.vertices {
		if x.Equals(elem) {
			return i
		}
	}

	return -1
}

// AdjacentOf returns the adjacent vertices of the given vertex.
//
// Parameters:
//   - from: the source vertex.
//
// Returns:
//   - []T: the adjacent vertices.
func (g *Graph[T]) AdjacentOf(from T) []T {
	index := g.IndexOf(from)
	if index == -1 {
		return nil
	}

	adj := make([]T, 0)

	for j, distance := range g.edges[index] {
		if distance != nil {
			adj = append(adj, g.vertices[j])
		}
	}

	return adj
}

// MakeTree creates a tree of the graph with the given root.
//
// Parameters:
//   - data: the data of the root.
//   - f: the nexts function.
//
// Returns:
//   - *WeightedGraphTree: the tree of the graph.
//   - error: an error if the tree creation fails.
func (g *Graph[T]) MakeTree(data T, info uc.Objecter, f trb.NextsFunc[T]) (*tr.Tree[*tn.TreeNode[T]], error) {
	var builder trb.Builder[T]

	builder.SetInfo(info)
	builder.SetNextFunc(f)

	return builder.Build(data)
}

// GetVertices returns the vertices in the graph.
//
// Returns:
//   - []T: the vertices.
func (g *Graph[T]) GetVertices() []T {
	return g.vertices
}

// GetEdges returns the edges in the graph.
//
// Returns:
//   - [][]*float64: the edges.
func (g *Graph[T]) GetEdges() [][]*float64 {
	return g.edges
}

// GetEdge returns the weight of the edge between the given vertices.
//
// Parameters:
//   - from: the source vertex.
//   - to: the destination vertex.
//
// Returns:
//   - float64: the weight of the edge.
//   - bool: true if the edge exists, otherwise false.
func (g *Graph[T]) GetEdge(from, to T) (float64, bool) {
	i := g.IndexOf(from)
	j := g.IndexOf(to)

	if i == -1 || j == -1 {
		return 0, false
	}

	w := g.edges[i][j]
	if w == nil {
		return 0, false
	}

	return *w, true
}
