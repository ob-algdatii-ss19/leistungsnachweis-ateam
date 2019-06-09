package algorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

type Clique []adjGraph.TrafficEntry

func getAllMaxCliques(returnType adjGraph.ConflictGraphPackage) {
	var listOfMaxClique []Clique = make([]Clique, 0)
	var currentClique Clique = make([]adjGraph.TrafficEntry, 0)
	var nextPossibleExpansion Clique = make([]adjGraph.TrafficEntry, 0)
	var previousExpansion Clique = make([]adjGraph.TrafficEntry, 0)

	for _, trafficentry := range returnType.Entries {
		if trafficentry.ChosenByUser {
			nextPossibleExpansion = append(nextPossibleExpansion, trafficentry)
		}
	}
	nextMaxCliques(listOfMaxClique, currentClique, nextPossibleExpansion, previousExpansion, returnType)

}

func nextMaxCliques(listOfMaxClique []Clique, currentClique Clique, nextPossibleExpansion Clique, previousExpansion Clique, returnType adjGraph.ConflictGraphPackage) []Clique {

	if oneOfPIsNeighborOfThemAll(nextPossibleExpansion, previousExpansion, returnType) {

		//Ende dieses Rekursionsbranches
	} else {
		for _, elementOfN := range nextPossibleExpansion {
			//newNextPossibleExpansion := remove(nextPossibleExpansion,index) aus meiner Sicht nicht unbedingt notwendig
			currentClique2 := append(currentClique, elementOfN)
			nN := getAllNeighbors(elementOfN, nextPossibleExpansion, returnType)
			pN := getAllNeighbors(elementOfN, previousExpansion, returnType)
			if len(nN) == 0 && len(pN) == 0 {
				listOfMaxClique = append(listOfMaxClique, currentClique2)
			} else {
				nextMaxCliques(listOfMaxClique, currentClique2, nN, pN, returnType)
			}
			previousExpansion = append(previousExpansion, elementOfN)
		}
	}

	return listOfMaxClique
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

func remove(s Clique, index int) Clique {
	s[index] = s[len(s)-1]
	return s[:len(s)-1]
}

func oneOfPIsNeighborOfThemAll(N Clique, P Clique, returnType adjGraph.ConflictGraphPackage) bool {
	var pIsNeighborOfThemAll bool = true
	for _, element := range N {
		indexElementN := adjGraph.GetIndexInConflictGraph(returnType, element)
		indexElementP := adjGraph.GetIndexInConflictGraph(returnType, P[0])
		if !returnType.ConflictGraph.GetMatrixEntryAtIndex(indexElementN, indexElementP) {
			pIsNeighborOfThemAll = false
			break
		}
	}
	return pIsNeighborOfThemAll
}
