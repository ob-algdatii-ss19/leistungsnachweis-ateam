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

	//graphObject := buildGraphObjectFromJSON(receivedData)

	//TODO @SDouglas3 build intolerance graph here from graphObject.
	// Outsource logic in separate function (for details see issue #21)

	if receivedData.Settings.Algorithm == BASIC_GREEDY {
		resultGraph := algorithms.BasicGreedy(nil)
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

	return adjGraph.NewGraphAdjMat(2)
}
