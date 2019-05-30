package adjGraph

import (
	"fmt"
)

func ExampleUGraphMat_UAddEdge() {
	g := NewUGraph(9)
	edges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}
	for _, edge := range edges {
		g.UAddEdge(edge.from, edge.to)
	}
	fmt.Printf("%v\n", g.UEdges())
	// Output:
	// [{1 2} {1 3} {1 6} {1 7} {2 1} {3 1} {4 5} {4 6} {5 4} {5 6} {5 7} {6 1} {6 4} {6 5} {6 6} {7 1} {7 5} {8 9} {9 8}]
}

func ExampleUGraphMat_UAdj() {
	g := NewUGraph(9)
	edges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}
	for _, edge := range edges {
		g.UAddEdge(edge.from, edge.to)
	}
	fmt.Printf("%v\n", g.UAdj(1))
	// Output:
	// [2 3 6 7]
}

func ExampleUGraphMat_UAdj_NoAdjNodes() {
	g := NewUGraph(9)
	edges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}
	for _, edge := range edges {
		g.UAddEdge(edge.from, edge.to)
	}
	fmt.Printf("%v\n", g.UAdj(0))
	// Output:
	// []
}

func ExampleUGraphMat_UNumberOfNodes() {

	g := NewUGraph(9)
	edges := []Edge{
		{1, 2},
		{1, 3},
		{1, 7},
		{4, 6},
		{5, 4},
		{6, 1},
		{6, 5},
		{6, 6},
		{7, 5},
		{9, 8},
	}
	for _, edge := range edges {
		g.UAddEdge(edge.from, edge.to)
	}
	fmt.Printf("%v\n", g.UNumberOfNodes())
	// Output:
	// 9
}

func ExampleUGraphMat_UNumberOfNodes_EmptyMatrix() {

	g := NewUGraph(0)

	fmt.Printf("%v\n", g.UNumberOfNodes())
	// Output:
	// 0
}
