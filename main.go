package main

import (
	"fmt"

	"github.com/TimothyGregg/ghost/graph"
)

func main() {
	fmt.Println("Ghost.")

	var g_test graph.Graph
	err := g_test.Add_Vertex_At([]int{0, 0, 0})
	if err != nil {
		panic(err)
	}
	fmt.Println("Step 1 Complete")
	err = g_test.Add_Vertex_At([]int{0, 0, 0})
	if err != nil {
		fmt.Println(err)
	}
	err = g_test.Add_Vertex_At([]int{0})
	if err != nil {
		fmt.Println(err)
	}
	err = g_test.Add_Vertex_At([]int{})
	if err != nil {
		panic(err)
	}
}
