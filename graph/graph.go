package graph

import (
	"errors"
	"fmt"
)


type vertex struct {
	position []int
}

type edge struct {
	vertices [2]*vertex
}

type Graph struct {
	dimensions int
	vertices   []*vertex
	edges      []*edge
}

func (g *Graph) Add_Vertex_At(position []int) error {
	return g.add_vertex(vertex{position: position})
}

func (g *Graph) add_vertex(v vertex) error {
	// On the addition of the first vertex, set the graph dimensions
	if len(g.vertices) == 0 {
		g.dimensions = len(v.position)
	}
	if g.dimensions != 0 { // Vertices cannot be the same on a graph with no dimensions (position is irrelevant)
		if g.dimensions != len(v.position) {
			return errors.New(fmt.Sprintf("Cannot add a vertex of dimension %d to a graph of dimension %d", len(v.position), g.dimensions))
		}
		for _, v_test := range g.vertices {
			same, err := v_test.same_as(&v)
			if err != nil {
				panic(err)
			}
			if same {
				return errors.New("Vertex already exists")
			}
		}
	}
	g.vertices = append(g.vertices, &v)
	return nil
}

func (v1 *vertex) same_as(v2 *vertex) (bool, error) {
	if len(v1.position) != len(v2.position) {
		return false, errors.New(fmt.Sprintf("Vertex of dimension %d cannot be compared to vertex of dimension %d", len(v1.position), len(v2.position)))
	}
	for ndx := range v1.position {
		if v1.position[ndx] != v2.position[ndx] {
			return false, nil
		}
	}
	// This allows for vertices of no dimensions to not be flagged as the same
	return (len(v1.position) > 0 && len(v2.position) > 0), nil
}

/*
// Any number of edges between vertices can be created
func (g *Graph) add_edge(e edge) error {
	for _, v_added := range e.vertices {
		found := false
		for _, v := range g.vertices {
			continue
		}
	}
	g.edges = append(g.edges, e)
	return nil
}
*/