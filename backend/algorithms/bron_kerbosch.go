package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

type Clique []adjGraph.TrafficEntry

var listOfMaxClique []Clique = make([]Clique, 0)

func getAllMaxCliques(returnType adjGraph.ConflictGraphPackage) []Clique {
	listOfMaxClique = make([]Clique, 0)
	var currentClique Clique = make([]adjGraph.TrafficEntry, 0)
	var nextPossibleExpansion Clique = make([]adjGraph.TrafficEntry, 0)
	var previousExpansion Clique = make([]adjGraph.TrafficEntry, 0)

	for _, trafficentry := range returnType.Entries {
		if trafficentry.ChosenByUser {
			nextPossibleExpansion = append(nextPossibleExpansion, trafficentry)
		}
	}
	/*listOfMaxClique = */ nextMaxCliques(currentClique, nextPossibleExpansion, previousExpansion, returnType)
	return listOfMaxClique
}

func nextMaxCliques(currentClique Clique, nextPossibleExpansion Clique, previousExpansion Clique, returnType adjGraph.ConflictGraphPackage) /* []Clique */ {

	if oneOfPIsNeighborOfThemAll(nextPossibleExpansion, previousExpansion, returnType) {
		fmt.Println("one is neighbor of them all")
		//Ende dieses Rekursionsbranches
	} else {
		length := len(nextPossibleExpansion)
		for i := 0; i < length; i++ {
			elementOfN := nextPossibleExpansion[0]
			//newNextPossibleExpansion := remove(nextPossibleExpansion,index) aus meiner Sicht nicht unbedingt notwendig
			currentClique2 := append(currentClique, elementOfN)
			nN := getAllNeighbors(elementOfN, nextPossibleExpansion, returnType)
			pN := getAllNeighbors(elementOfN, previousExpansion, returnType)
			if len(nN) == 0 && len(pN) == 0 {
				listOfMaxClique = append(listOfMaxClique, currentClique2)
			} else {
				nextMaxCliques(currentClique2, nN, pN, returnType)
			}
			previousExpansion = append(previousExpansion, elementOfN)
			nextPossibleExpansion = nextPossibleExpansion[1:]
		}
	}
}

func getAllNeighbors(element adjGraph.TrafficEntry, clique Clique, returnType adjGraph.ConflictGraphPackage) Clique {
	var neighbors Clique = make([]adjGraph.TrafficEntry, 0)
	for _, elementClique := range clique {
		var indexElementClique int = adjGraph.GetIndexInConflictGraph(returnType, elementClique)
		var indexElement int = adjGraph.GetIndexInConflictGraph(returnType, element)
		if returnType.ConflictGraph.GetMatrixEntryAtIndex(indexElement, indexElementClique) {
			neighbors = append(neighbors, elementClique)
		}
	}
	return neighbors
}

func oneOfPIsNeighborOfThemAll(N Clique, P Clique, returnType adjGraph.ConflictGraphPackage) bool {
	var pIsNeighborOfThemAll bool = true
	if len(P) != 0 {
		for _, element := range N {
			indexElementN := adjGraph.GetIndexInConflictGraph(returnType, element)
			indexElementP := adjGraph.GetIndexInConflictGraph(returnType, P[0])
			if !returnType.ConflictGraph.GetMatrixEntryAtIndex(indexElementN, indexElementP) {
				pIsNeighborOfThemAll = false
				break
			}
		}
	} else {
		pIsNeighborOfThemAll = false
	}
	return pIsNeighborOfThemAll
}
