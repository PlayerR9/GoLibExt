package Graph

import (
	"slices"

	uts "github.com/PlayerR9/MyGoLib/Utility/Sorting"
)

// remapVertices shifts all vertices to the right starting from the given position.
//
// Parameters:
//   - pos: The position to start shifting vertices.
func (g *Graph[T]) remapVertices(pos int) {
	for i := len(g.edges) - 1; i >= pos; i-- {
		to := g.edges[i]

		// Shift all elements to the right
		for j := len(to) - 1; j >= pos; j-- {
			to[j+1] = to[j]
		}

		g.edges[i+1] = to
	}
}

type Graph[T VertexElementer] struct {
	vertices []*Vertex[T]
	edges    map[int]map[int]bool
	sf       uts.SortFunc[*Vertex[T]]
}

func NewGraph[T VertexElementer](sf uts.SortFunc[*Vertex[T]]) *Graph[T] {
	if sf == nil {
		return nil
	}

	return &Graph[T]{
		vertices: make([]*Vertex[T], 0),
		edges:    make(map[int]map[int]bool),
		sf:       sf,
	}
}

func (g *Graph[T]) AddVertex(v *Vertex[T]) {
	if v == nil {
		return
	}

	pos, ok := slices.BinarySearchFunc(g.vertices, v, g.sf)
	if !ok {
		g.vertices = slices.Insert(g.vertices, pos, v)
		g.remapVertices(pos)
	}

	g.edges[pos] = make(map[int]bool)
}

func (g *Graph[T]) AddEdge(from, to *Vertex[T]) {
	if to == nil {
		if from == nil {
			return
		}

		from.isFinal = true
	} else if from == nil {
		to.isInitial = true
	}

	g.AddVertex(from)
	g.AddVertex(to)

	if from == nil {
		return
	} else if to == nil {
		return
	}

	fromIndex, _ := slices.BinarySearchFunc(g.vertices, from, g.sf)
	toIndex, _ := slices.BinarySearchFunc(g.vertices, to, g.sf)

	g.edges[fromIndex][toIndex] = true
}
