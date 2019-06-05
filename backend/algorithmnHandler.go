package backend

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/algorithms"
)

var edgeNamesToNodeLetter = map[string]string{
	string(adjGraph.ABC + "-" + adjGraph.MNO): "A",
	string(adjGraph.ABC + "-" + adjGraph.IJK): "B",
	string(adjGraph.ABC + "-" + adjGraph.EFG): "C",
	string(adjGraph.ABC + "-" + adjGraph.P1):  "D1",
	string(adjGraph.ABC + "-" + adjGraph.P2):  "D2",
	string(adjGraph.EFG + "-" + adjGraph.ABC): "E",
	string(adjGraph.EFG + "-" + adjGraph.MNO): "F",
	string(adjGraph.EFG + "-" + adjGraph.IJK): "G",
	string(adjGraph.EFG + "-" + adjGraph.P1):  "H1",
	string(adjGraph.EFG + "-" + adjGraph.P2):  "H2",
	string(adjGraph.IJK + "-" + adjGraph.EFG): "I",
	string(adjGraph.IJK + "-" + adjGraph.ABC): "J",
	string(adjGraph.IJK + "-" + adjGraph.MNO): "K",
	string(adjGraph.IJK + "-" + adjGraph.P1):  "L1",
	string(adjGraph.IJK + "-" + adjGraph.P2):  "L2",
	string(adjGraph.MNO + "-" + adjGraph.IJK): "M",
	string(adjGraph.MNO + "-" + adjGraph.EFG): "N",
	string(adjGraph.MNO + "-" + adjGraph.ABC): "O",
	string(adjGraph.MNO + "-" + adjGraph.P1):  "P1",
	string(adjGraph.MNO + "-" + adjGraph.P2):  "P2",
}

/*
Handle all calls to the different algorithms
*/
func HandleAlgorithmCalls(receivedData GuiRequestData) JsonResponse {
	fmt.Println("[INFO] Called algorithmHandler.go")

	graphObject := adjGraph.MakeConflictGraphOutOfConnectionGraph(buildGraphObjectFromJSON(receivedData))

	if receivedData.Settings.Algorithm == BASIC_GREEDY {
		resultGraphWithNodeNames := algorithms.BasicGreedy(graphObject)
		fmt.Println("[DEBUG] generated result graph with Basic Greedy Algorithm ", resultGraphWithNodeNames)

		resultGraphWithLetters := changeNodeNumbersToLetters(resultGraphWithNodeNames, graphObject.Entries)

		return JsonResponse{true, resultGraphWithLetters}
	} else {
		return JsonResponse{false, nil}
	}
}

/*
Change Node-Numbers in the Array with the traffic phases to capital letters.
*/
func changeNodeNumbersToLetters(resultGraph [][]adjGraph.Node, trafficEntries []adjGraph.TrafficEntry) [][]string {

	allPhases := make([][]string, 0)

	for _, nodeArray := range resultGraph {

		trafficPhase := make([]string, 0)

		for _, node := range nodeArray {

			nodeLetter := edgeNamesToNodeLetter[string(trafficEntries[node-1].From)+"-"+string(trafficEntries[node-1].To)]

			trafficPhase = append(trafficPhase, nodeLetter)

		}

		allPhases = append(allPhases, trafficPhase)

	}

	return allPhases
}

/*
Build a Graph-Object from the received JSON-Data
*/
func buildGraphObjectFromJSON(data GuiRequestData) adjGraph.AdjMat {

	var countNodes int = 6 // second last = with pedestrian, last: with island

	graph := adjGraph.NewGraphAdjMat(countNodes)

	//top=1, right=2, bottom=3, left=4
	//Lanes: right, middle, left

	var left adjGraph.Node = 1
	var bottom adjGraph.Node = 2
	var right adjGraph.Node = 3
	var top adjGraph.Node = 4

	var pedestrian adjGraph.Node = 5
	var pedestrianWithIsland adjGraph.Node = 6

	//left Node
	if data.Intersection.Left.RightLane {
		graph.AddEdge(left, bottom)
	}
	if data.Intersection.Left.StraightLane {
		graph.AddEdge(left, right)
	}
	if data.Intersection.Left.LeftLane {
		graph.AddEdge(left, top)
	}
	if data.Intersection.Left.Pedestrian == NORMAL {
		graph.AddEdge(left, pedestrian)
	}
	if data.Intersection.Left.Pedestrian == WITH_ISLAND {
		graph.AddEdge(left, pedestrian)
		graph.AddEdge(left, pedestrianWithIsland)
	}

	//bottom Node
	if data.Intersection.Buttom.RightLane {
		graph.AddEdge(bottom, right)
	}
	if data.Intersection.Buttom.StraightLane {
		graph.AddEdge(bottom, top)
	}
	if data.Intersection.Buttom.LeftLane {
		graph.AddEdge(bottom, left)
	}
	if data.Intersection.Buttom.Pedestrian == NORMAL {
		graph.AddEdge(bottom, pedestrian)
	}
	if data.Intersection.Buttom.Pedestrian == WITH_ISLAND {
		graph.AddEdge(bottom, pedestrian)
		graph.AddEdge(bottom, pedestrianWithIsland)
	}

	//right Node
	if data.Intersection.Right.RightLane {
		graph.AddEdge(right, top)
	}
	if data.Intersection.Right.StraightLane {
		graph.AddEdge(right, left)
	}
	if data.Intersection.Right.LeftLane {
		graph.AddEdge(right, bottom)
	}
	if data.Intersection.Right.Pedestrian == NORMAL {
		graph.AddEdge(right, pedestrian)
	}
	if data.Intersection.Right.Pedestrian == WITH_ISLAND {
		graph.AddEdge(right, pedestrian)
		graph.AddEdge(right, pedestrianWithIsland)
	}

	//Top Node
	if data.Intersection.Top.RightLane {
		graph.AddEdge(top, left)
	}
	if data.Intersection.Top.StraightLane {
		graph.AddEdge(top, bottom)
	}
	if data.Intersection.Top.LeftLane {
		graph.AddEdge(top, right)
	}
	if data.Intersection.Top.Pedestrian == NORMAL {
		graph.AddEdge(top, pedestrian)
	}
	if data.Intersection.Top.Pedestrian == WITH_ISLAND {
		graph.AddEdge(top, pedestrian)
		graph.AddEdge(top, pedestrianWithIsland)
	}

	//fmt.Println("[DEBUG] buildGraphObjectFromJSON graphExport", graph)
	return graph
}
