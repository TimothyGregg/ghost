package graph

import (
	"errors"
	"fmt"
)


type Vertex struct {
	position []int
}

type Edge struct {
	vertices [2]*Vertex
}

type Graph struct {
	dimensions int
	vertices   []*Vertex
	edges      []*Edge
}

func (g *Graph) Add_Vertex_At(position []int) error {
	v := Vertex{position: position}
	return g.add_vertex(&v)
}

func (g *Graph) add_vertex(v *Vertex) error {
	// On the addition of the first vertex, set the graph dimensions
	if len(g.vertices) == 0 {
		g.dimensions = len(v.position)
	}
	if g.dimensions != 0 { // Vertices cannot be the same on a graph with no dimensions (position is irrelevant)
		if g.dimensions != len(v.position) {
			return errors.New(fmt.Sprintf("Cannot add a vertex of dimension %d to a graph of dimension %d", len(v.position), g.dimensions))
		}
		for _, v_test := range g.vertices {
			same, err := v_test.same_as(v)
			if err != nil {
				panic(err)
			}
			if same {
				return errors.New("Vertex already exists")
			}
		}
	}
	g.vertices = append(g.vertices, v)
	return nil
}

func (v1 *Vertex) same_as(v2 *Vertex) (bool, error) {
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

func (g *Graph) add_edge(e *Edge) error {
	// Any number of edges between vertices can be created
	found := 0
	for _, v_added := range e.vertices {
		for _, v := range g.vertices {
			if v == v_added {
				found += 1
			}
		}
	}
	if found < 2 {
		return errors.New("The edge specified does not connect two vertices within the graph")
	}
	g.edges = append(g.edges, e)
	return nil
}
