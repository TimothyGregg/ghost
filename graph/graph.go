package graph

import "bytes"

type vertex struct {
	position []int
}

type edge struct {
	vertices [2]vertex
}

type Graph struct {
	dimensions int
	vertices   []vertex
	edges      []edge
}

func (g *Graph) addVertex(v vertex) error {
	for v_test := range (*g).vertices {
		if sameVertex(v_test, v) {

		}
	}
	(*g).vertices = append((*g).vertices, v)
	return nil
}

func sameVertex(v1, v2 *vertex) bool {
	return bytes.Equal(*v1, *v2)
	if (*v1).position == (*v2).position {
		return true
	}
	return false
}