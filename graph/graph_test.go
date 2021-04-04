package graph

import (
	"strings"
	"testing"
)

func TestGraph(t *testing.T) {
	var test_graph Graph

	v1 := vertex{position: []int{0, 0}}
	v2 := vertex{position: []int{0, 0, 0}}
	v3 := vertex{position: []int{0, 1}}
	v4 := vertex{position: []int{0, 0}}
	v5 := vertex{}
	v6 := vertex{}

	// Add Vertex
	err := test_graph.Add_Vertex_At([]int{0, 0})
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}

	// Add Vertex where another is already present
	err = test_graph.Add_Vertex_At([]int{0, 0})
	if err == nil {
		t.Log("error should be \"Vertex already exists\"")
		t.Fail()
	} else if strings.Compare(err.Error(), "Vertex already exists") != 0 {
		t.Log("Error should be \"Vertex already exists\"; instead, got", err)
		t.Fail()
	}

	// Add Vertex of incorrect dimensions
	err = test_graph.Add_Vertex_At([]int{0, 0, 0})
	if err == nil {
		t.Log("Error should be \"Cannot add a vertex of dimension 3 to a graph of dimension 2\"")
		t.Fail()
	} else if strings.Compare(err.Error(), "Cannot add a vertex of dimension 3 to a graph of dimension 2") != 0 {
		t.Log("Error should be \"Cannot add a vertex of dimension 3 to a graph of dimension 2\"; instead, got", err)
		t.Fail()
	}

	// Compare vertices of different dimensions
	result, err := v1.same_as(&v2)
	if err == nil {
		t.Log("\"Vertex of dimension 2 cannot be compared to vertex of dimension 3\"")
		t.Fail()
	} else if strings.Compare(err.Error(), "Vertex of dimension 2 cannot be compared to vertex of dimension 3") != 0 {
		t.Log("Error should be \"Vertex of dimension 2 cannot be compared to vertex of dimension 3\"; instead, got", err)
		t.Fail()
	}
	if result == true {
		t.Log("Vertex at (0, 0) compared to vertex at (0, 0, 0) and found to be the same.")
		t.Fail()
	}

	result, err = v3.same_as(&v5)
	if err == nil {
		t.Log("\"Vertex of dimension 2 cannot be compared to vertex of dimension 0\"")
		t.Fail()
	} else if strings.Compare(err.Error(), "Vertex of dimension 2 cannot be compared to vertex of dimension 0") != 0 {
		t.Log("Error should be \"Vertex of dimension 2 cannot be compared to vertex of dimension 0\"; instead, got", err)
		t.Fail()
	}
	if result == true {
		t.Log("Vertex at (0, 1) compared to vertex of zero dimensions and found to be the same.")
		t.Fail()
	}

	// Compare dissimilar vertices
	result, err = v1.same_as(&v3)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	if result == true {
		t.Log("Vertex at (0, 0) compared to vertex at (0, 1) and found to be the same.")
		t.Fail()
	}

	// Compare similar vertices
	result, err = v1.same_as(&v4)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	if result == false {
		t.Log("Vertex at (0, 0) compared to vertex at (0, 0) and found to be different.")
		t.Fail()
	}

	result, err = v5.same_as(&v6)
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	if result == true {
		t.Log("Vertex of zero dimensions compared to vertex of zero dimensions and found to be the same.")
		t.Fail()
	}
	

	// Logic for a graph of dimension 0
	test_graph = *new(Graph)
	err = test_graph.add_vertex(vertex{})
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
	err = test_graph.add_vertex(vertex{})
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
		// Vertices cannot be the same on a graph with no dimensions (position is irrelevant)
	err = test_graph.add_vertex(vertex{position: []int{0}})
	if err != nil {
		t.Log("error should be nil", err)
		t.Fail()
	}
}