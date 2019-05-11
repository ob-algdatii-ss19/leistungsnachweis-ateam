package backend

import (
	"fmt"
	"github.com/mountainflo/leistungsnachweis-ateam/backend/adjGraph"
	"github.com/mountainflo/leistungsnachweis-ateam/backend/algorithms"
)

/*
Handle all calls to the different algorithmns
*/
func HandleAlgorithmCalls(receivedData GuiRequestData) JsonResponse {
	fmt.Println("[INFO] Called algorithmHandler.go")

	graphObject := buildGraphObjectFromJSON(receivedData)

	if receivedData.Settings.Algorithm == BASIC_GREEDY {
		resultGraph := algorithms.BasicGreedy(graphObject)
		fmt.Println("[DEBUG] generated result graph with Basic Greedy Algorithmn ", resultGraph)

		return JsonResponse{true} //TODO return graph object
	} else {
		return JsonResponse{false}
	}
}

/*
Build a Graph-Object from the received JSON-Data
*/
func buildGraphObjectFromJSON(data GuiRequestData) adjGraph.Graph {
	return adjGraph.NewGraphAdjMat(2)
}
