package main

import (
	"./adjGraph"
	"fmt"
)

func main() {
	g := adjGraph.NewGraphAdjMat(10)
	g.AddEdge(1, 2)
	// print(g)

	myMap := map[int]string{1: "eins", 2: "zwei"}

	for i, s := range myMap {
		myMap[i] = fmt.Sprintf("alles falsch %s", s)
	}

	fmt.Printf("myMap = %v\n", myMap)

}

func print(g adjGraph.Graph) {
	fmt.Printf("edges = %v\n", g.Edges())
}
