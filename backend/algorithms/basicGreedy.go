package algorithms

import (
	"fmt"
	"github.com/mountainflo/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with basic Greedy algorithm
*/
func BasicGreedy(graphData adjGraph.Graph) adjGraph.Graph {
	fmt.Println("[INFO] Called BasicGreedy Algorithmn")
	return adjGraph.NewGraphAdjMat(2)
}
