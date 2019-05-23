package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with basic Greedy algorithm
*/
func BasicGreedy(graphData adjGraph.Graph) [][]adjGraph.Node {
	fmt.Println("[INFO] Called BasicGreedy Algorithm")

	//TODO implement basic greedy algorithm. (for details see issue #12 and issue #3)

	for _, edge := range graphData.Edges() {
		fmt.Println(edge)
	}

	return [][]adjGraph.Node{
		{adjGraph.Node(1)},
		{adjGraph.Node(2), adjGraph.Node(4)},
		{adjGraph.Node(3)},
	}
}
