package backend

import (
	"../backend/adjGraph"
	"../backend/algorithms"
	"fmt"
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

	var countNodes int=getCountOfNodes(data);


	fmt.Println("count Nodes", countNodes)


	graph :=adjGraph.NewGraphAdjMat(countNodes);

	//top=1, right=2, bottom=3, left=4
	//Lanes: right, middle, left

	var startNode adjGraph.Node=1;
	var endNode adjGraph.Node=startNode+1;
	if(data.Intersection.Top.LeftLane){
		if(data.Intersection.Top.Pedestrian==WITH_ISLAND) {
			endNode++; //with island at top
		}
		if(data.Intersection.Left.Pedestrian==WITH_ISLAND) {
			endNode++; //with island at top
		}
		graph.AddEdge(startNode, endNode); //with island
	}

	if(data.Intersection.Top.Pedestrian==WITH_ISLAND) {
		startNode++; //with island at top
	}
	startNode++;
	if(data.Intersection.Right.RightLane){
		graph.AddEdge(startNode, startNode-1);
	}

	if(data.Intersection.Right.LeftLane){
		endNode++;
		if(data.Intersection.Buttom.Pedestrian==WITH_ISLAND) {
			endNode++; //with island at top
			startNode++;
		}
		graph.AddEdge(startNode, endNode); //with island
	}
	if(data.Intersection.Top.StraightLane){
		graph.AddEdge(1, endNode); //with island
	}

	startNode++;
	if(data.Intersection.Right.Pedestrian==WITH_ISLAND) {
		startNode++; //with island at top
	}
	if(data.Intersection.Buttom.RightLane){
		graph.AddEdge(startNode, startNode-1);
	}
	if(data.Intersection.Buttom.StraightLane){
		var eN adjGraph.Node =1;
		if(data.Intersection.Top.Pedestrian==WITH_ISLAND) {
			eN++; //with island at top
		}
		graph.AddEdge(startNode, eN);
	}

	if(data.Intersection.Buttom.LeftLane){
		endNode++;
		if(data.Intersection.Left.Pedestrian==WITH_ISLAND) {
			endNode++; //with island at top
			startNode++;
		}
		graph.AddEdge(startNode, endNode); //with island
	}

	var right adjGraph.Node =2;
	if(data.Intersection.Right.StraightLane){
		if(data.Intersection.Top.Pedestrian==WITH_ISLAND) {
			right++;
		}
		graph.AddEdge(right, endNode); //with island
	}

	startNode++;
	if(data.Intersection.Left.Pedestrian==WITH_ISLAND) {
		startNode++; //with island at top
	}
	if(data.Intersection.Left.RightLane){
		graph.AddEdge(startNode, startNode-1);
	}

	if(data.Intersection.Left.StraightLane){
		if(data.Intersection.Right.Pedestrian==WITH_ISLAND) {
			right++; //with island at top
		}
		graph.AddEdge(startNode, right);
	}

	if(data.Intersection.Left.LeftLane){
		var eN adjGraph.Node =1;
		if(data.Intersection.Top.Pedestrian==WITH_ISLAND) {
			eN++; //with island at top
		}
		graph.AddEdge(startNode, endNode); //with island
	}

	if(data.Intersection.Top.RightLane){
		graph.AddEdge(1, endNode); //with island
	}


	fmt.Println("[DEBUG] buildGraphObjectFromJSON graphExport", graph)
	return graph
}


//returns the count of nodes (int-value between 0 and 8)
func getCountOfNodes(data GuiRequestData) int{
	var countNodes int=0;

	var top bool=false;
	//from top away
	if(data.Intersection.Top.RightLane ||
		data.Intersection.Top.StraightLane ||
		data.Intersection.Top.LeftLane){
			fmt.Println("from Top away", 1)
			top=true;
			countNodes++;
	}
	//inside top
	if(data.Intersection.Right.RightLane ||
		data.Intersection.Buttom.StraightLane||
		data.Intersection.Left.LeftLane){
		if(!top || data.Intersection.Top.Pedestrian == WITH_ISLAND){
			fmt.Println("inside Top", 1)
			countNodes++;
		}
	}

	var right bool=false;
	//from right away
	if(data.Intersection.Right.RightLane ||
		data.Intersection.Right.StraightLane ||
		data.Intersection.Right.LeftLane){
			fmt.Println("from Right way", 1)
			right=true;
			countNodes++;
	}
	//inside right
	if(data.Intersection.Top.LeftLane ||
		data.Intersection.Buttom.RightLane||
		data.Intersection.Left.StraightLane){
		if(!right || data.Intersection.Right.Pedestrian == WITH_ISLAND){
			fmt.Println("inside right", 1)
			countNodes++;
		}
	}

	var bottom bool=false;
	//from bottom away
	if(data.Intersection.Buttom.RightLane ||
		data.Intersection.Buttom.StraightLane ||
		data.Intersection.Buttom.LeftLane){
		fmt.Println("from bottom away", 1)
		bottom=true;
		countNodes++;
	}
	//inside bottom
	if(data.Intersection.Top.StraightLane ||
		data.Intersection.Right.LeftLane||
		data.Intersection.Left.RightLane){
		if(!bottom || data.Intersection.Buttom.Pedestrian == WITH_ISLAND){
			fmt.Println("inside bottom", 1)
			countNodes++;
		}
	}

	var left bool=false;
	//from left away
	if(data.Intersection.Left.RightLane ||
		data.Intersection.Left.StraightLane ||
		data.Intersection.Left.LeftLane){
			fmt.Println("from left away", 1)
			left=true;
			countNodes++;
	}
	//inside left
	if(data.Intersection.Top.RightLane ||
		data.Intersection.Right.StraightLane||
		data.Intersection.Buttom.LeftLane){
		if(!left || data.Intersection.Left.Pedestrian == WITH_ISLAND){
			fmt.Println("inside left", 1)
			countNodes++;
		}
	}


	return countNodes;
}
