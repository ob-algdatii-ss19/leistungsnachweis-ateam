package backend

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/algorithms"
)

/*
Handle all calls to the different algorithmns
*/
func HandleAlgorithmCalls(receivedData GuiRequestData) JsonResponse {
	fmt.Println("[INFO] Called algorithmHandler.go")

	graphObject := buildGraphObjectFromJSON(receivedData)

	//TODO @SDouglas3 build intolerance graph here from graphObject.
	// Outsource logic in separate function (for details see issue #21)

	if receivedData.Settings.Algorithm == BASIC_GREEDY {
		resultGraph := algorithms.BasicGreedy(graphObject)
		fmt.Println("[DEBUG] generated result graph with Basic Greedy Algorithm ", resultGraph)

		return JsonResponse{true, resultGraph}
	} else {
		return JsonResponse{false, nil}
	}
}

/*
Build a Graph-Object from the received JSON-Data
*/
func buildGraphObjectFromJSON(data GuiRequestData) adjGraph.Graph {

	//TODO @mike-la build graph object here (for details see issue #20)

	var countNodes int = getCountOfNodes(data)

	fmt.Println("count Nodes", countNodes)

	graph := adjGraph.NewGraphAdjMat(countNodes)

	//top=1, right=2, bottom=3, left=4
	//Lanes: right, middle, left

	var topOut adjGraph.Node = 1
	var topIn adjGraph.Node = topOut
	if (data.Intersection.Top.Pedestrian == WITH_ISLAND){
		topIn++
	}

	var leftOut adjGraph.Node = topIn +1;
	var leftIn adjGraph.Node =leftOut
	if (data.Intersection.Left.Pedestrian == WITH_ISLAND){
		leftIn++
	}

	var bottomOut adjGraph.Node = leftIn+1;
	var bottomIn adjGraph.Node = bottomOut
	if (data.Intersection.Buttom.Pedestrian == WITH_ISLAND){
		bottomIn++
	}

	var rightOut adjGraph.Node = bottomIn+1;
	var rightIn adjGraph.Node = rightOut
	if (data.Intersection.Right.Pedestrian == WITH_ISLAND){
		rightIn++
	}



	fmt.Println("[DEBUG] node Top", topOut, topIn);
	fmt.Println("[DEBUG] node left", leftOut, leftIn);
	fmt.Println("[DEBUG] node Bottom", bottomOut, bottomIn);
	fmt.Println("[DEBUG] node Right", rightOut, rightIn);


	//Top Node
	if data.Intersection.Top.RightLane {
		graph.AddEdge(topOut, rightIn)
	}
	if data.Intersection.Top.StraightLane {
		graph.AddEdge(topOut, bottomIn)
	}
	if data.Intersection.Top.LeftLane {
		graph.AddEdge(topOut, leftIn)
	}

	//right Node
	if data.Intersection.Right.RightLane {
		graph.AddEdge(rightOut, bottomIn)
	}
	if data.Intersection.Right.StraightLane {
		graph.AddEdge(rightOut, leftIn)
	}
	if data.Intersection.Right.LeftLane {
		graph.AddEdge(rightOut, topIn)
	}

	//bottom Node
	if data.Intersection.Buttom.RightLane {
		graph.AddEdge(bottomOut, leftIn)
	}
	if data.Intersection.Buttom.StraightLane {
		graph.AddEdge(bottomOut, topIn)
	}
	if data.Intersection.Buttom.LeftLane {
		graph.AddEdge(bottomOut, rightIn)
	}

	//left Node
	if data.Intersection.Left.RightLane {
		graph.AddEdge(leftOut, bottomIn)
	}
	if data.Intersection.Left.StraightLane {
		graph.AddEdge(bottomOut, rightIn)
	}
	if data.Intersection.Left.LeftLane {
		graph.AddEdge(bottomOut, topIn)
	}



	fmt.Println("[DEBUG] buildGraphObjectFromJSON graphExport", graph)
	return graph
}

//returns the count of nodes (int-value between 0 and 8)
func getCountOfNodes(data GuiRequestData) int {
	var countNodes int = 0

	var top bool = false
	//from top away
	if data.Intersection.Top.RightLane ||
		data.Intersection.Top.StraightLane ||
		data.Intersection.Top.LeftLane {
		fmt.Println("from Top away", 1)
		top = true
		countNodes++
	}
	//inside top
	if data.Intersection.Right.RightLane ||
		data.Intersection.Buttom.StraightLane ||
		data.Intersection.Left.LeftLane {
		if !top || data.Intersection.Top.Pedestrian == WITH_ISLAND {
			fmt.Println("inside Top", 1)
			countNodes++
		}
	}

	var right bool = false
	//from right away
	if data.Intersection.Right.RightLane ||
		data.Intersection.Right.StraightLane ||
		data.Intersection.Right.LeftLane {
		fmt.Println("from Right way", 1)
		right = true
		countNodes++
	}
	//inside right
	if data.Intersection.Top.LeftLane ||
		data.Intersection.Buttom.RightLane ||
		data.Intersection.Left.StraightLane {
		if !right || data.Intersection.Right.Pedestrian == WITH_ISLAND {
			fmt.Println("inside right", 1)
			countNodes++
		}
	}

	var bottom bool = false
	//from bottom away
	if data.Intersection.Buttom.RightLane ||
		data.Intersection.Buttom.StraightLane ||
		data.Intersection.Buttom.LeftLane {
		fmt.Println("from bottom away", 1)
		bottom = true
		countNodes++
	}
	//inside bottom
	if data.Intersection.Top.StraightLane ||
		data.Intersection.Right.LeftLane ||
		data.Intersection.Left.RightLane {
		if !bottom || data.Intersection.Buttom.Pedestrian == WITH_ISLAND {
			fmt.Println("inside bottom", 1)
			countNodes++
		}
	}

	var left bool = false
	//from left away
	if data.Intersection.Left.RightLane ||
		data.Intersection.Left.StraightLane ||
		data.Intersection.Left.LeftLane {
		fmt.Println("from left away", 1)
		left = true
		countNodes++
	}
	//inside left
	if data.Intersection.Top.RightLane ||
		data.Intersection.Right.StraightLane ||
		data.Intersection.Buttom.LeftLane {
		if !left || data.Intersection.Left.Pedestrian == WITH_ISLAND {
			fmt.Println("inside left", 1)
			countNodes++
		}
	}

	return countNodes
}
