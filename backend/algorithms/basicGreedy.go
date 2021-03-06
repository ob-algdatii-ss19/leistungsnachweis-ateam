package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"sort"
)

/*
Calculate the optimization of the traffic lights with basic Greedy algorithm
*/
func BasicGreedy(returnType adjGraph.ConflictGraphPackage) [][]adjGraph.Node {
	fmt.Println("[INFO] Called BasicGreedy Algorithm")

	graphData := returnType.ConflictGraph
	trafficEntries := returnType.Entries

	if graphData == nil || trafficEntries == nil {
		return [][]adjGraph.Node{}
	}

	numberOfNodes := graphData.UNumberOfNodes()
	coloredNodes := make([]int, numberOfNodes+1)
	listOfColorsAndNodes := make([][]adjGraph.Node, numberOfNodes+1)

	//go through all nodes (vertices)
	for i := 1; i <= numberOfNodes; i++ {

		//user selected the node via checkbox in the gui
		if trafficEntries[i-1].ChosenByUser {
			unavailableColors := getUsedColorsByAdjacentNodes(graphData, adjGraph.Node(i), coloredNodes)

			lowestFreeColor := getLowestUnusedColorOfAdjacentNodes(unavailableColors)

			coloredNodes[i] = lowestFreeColor
			listOfColorsAndNodes[lowestFreeColor] = append(listOfColorsAndNodes[lowestFreeColor], adjGraph.Node(i))
		}
	}

	return getTrimmedListOfColorsAndNodes(listOfColorsAndNodes)
}

func getUsedColorsByAdjacentNodes(graphData adjGraph.UGraph, node adjGraph.Node, coloredNodes []int) []int {

	unavailableColors := make([]int, 0)
	adjacentNodes := graphData.UAdj(node)

	for j := 0; j < len(adjacentNodes); j++ {

		nodeNumber := adjacentNodes[j]

		if coloredNodes[nodeNumber] != 0 {
			unavailableColors = append(unavailableColors, coloredNodes[nodeNumber])
		}
	}

	sort.Ints(unavailableColors)

	return unavailableColors
}

func getLowestUnusedColorOfAdjacentNodes(unavailableColors []int) int {

	if len(unavailableColors) == 0 {

		return 1

	} else {

		sort.Ints(unavailableColors)

		lowestFreeColor := -1
		previousColor := 0
		for _, color := range unavailableColors {
			if (color != previousColor+1 && color > previousColor+1) || (previousColor == 0 && color > 1) {
				lowestFreeColor = previousColor + 1
				break
			}
			previousColor = color
		}

		if lowestFreeColor == -1 {
			lowestFreeColor = unavailableColors[len(unavailableColors)-1] + 1
		}

		return lowestFreeColor
	}

}

func getTrimmedListOfColorsAndNodes(listOfColorsAndNodes [][]adjGraph.Node) [][]adjGraph.Node {

	result := make([][]adjGraph.Node, 0)

	for _, item := range listOfColorsAndNodes {
		if item != nil {
			result = append(result, item)
		}
	}

	return result
}
