package algorithms

import (
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

type Clique []adjGraph.TrafficEntry

func GetMaxCliques(returnType adjGraph.ConflictGraphPackage) [][]adjGraph.Node {
	listOfMaxClique := make([]Clique, 0)
	var currentClique Clique = make([]adjGraph.TrafficEntry, 0)
	var nextPossibleExpansion Clique = make([]adjGraph.TrafficEntry, 0)
	var previousExpansion Clique = make([]adjGraph.TrafficEntry, 0)

	for _, trafficentry := range returnType.Entries {
		if trafficentry.ChosenByUser {
			nextPossibleExpansion = append(nextPossibleExpansion, trafficentry)
		}
	}
	listOfMaxClique = bronkerbIterative(currentClique, nextPossibleExpansion, previousExpansion, returnType, listOfMaxClique)
	listOfMaxClique = getMininmalMaxCliques(returnType, listOfMaxClique)
	maxNodes := translateResult(listOfMaxClique, returnType)
	return maxNodes
}

func translateResult(clique []Clique, returnType adjGraph.ConflictGraphPackage) [][]adjGraph.Node {
	result_nodes := make([][]adjGraph.Node, 0)
	for index, element := range clique {
		result_nodes = append(result_nodes, make([]adjGraph.Node, 0))
		for _, clique_element := range element {
			nodeInt := adjGraph.GetIndexInConflictGraph(returnType, clique_element)
			result_nodes[index] = append(result_nodes[index], adjGraph.Node(nodeInt))
		}
	}
	return result_nodes
}

func getMininmalMaxCliques(graphPackage adjGraph.ConflictGraphPackage, listOfMaxCliques []Clique) []Clique {

	optimal := isOptimal(listOfMaxCliques, graphPackage)

	if optimal {
		return listOfMaxCliques
	}
	copyUserinputList := make([]adjGraph.TrafficEntry, len(graphPackage.Entries))
	copy(copyUserinputList, graphPackage.Entries)
	for _, elementt := range graphPackage.Entries {
		if !elementt.ChosenByUser {
			copyUserinputList, _ = removeItemOfClique(copyUserinputList, elementt)
		}
	}
	finalList := make([]Clique, 0)

	for !optimal {
		largestIndex := -1
		length := -1
		//find largest Clique
		for index, element := range listOfMaxCliques {
			if len(element) > length {
				length = len(element)
				largestIndex = index
			}
		}
		finalList = append(finalList, listOfMaxCliques[largestIndex])

		for _, elementLargestClique := range listOfMaxCliques[largestIndex] {
			copyUserinputList, _ = removeItemOfClique(copyUserinputList, elementLargestClique)
		}

		isOwnClique := true
		for _, element := range copyUserinputList {
			isOwnClique = IsNeighborOfThemAll(copyUserinputList, element, graphPackage)
			if !isOwnClique {
				break
			}
		}
		if isOwnClique {
			finalList = append(finalList, copyUserinputList)
			//return finalList
		} else {
			//newList := getMininmalMaxCliques(graphPackage,bronkerbIterative(make([]adjGraph.TrafficEntry, 0),copyUserinputList,make([]adjGraph.TrafficEntry, 0),graphPackage,make([]Clique, 0)))
			listOfMaxCliques = bronkerbIterative(make([]adjGraph.TrafficEntry, 0), copyUserinputList, make([]adjGraph.TrafficEntry, 0), graphPackage, make([]Clique, 0))
			//newList = append(newList, finalList[0])
			//return newList
		}
		newGraphPackage := adjGraph.ConflictGraphPackage{copyUserinputList, graphPackage.ConflictGraph}
		optimal = isOptimal(listOfMaxCliques, newGraphPackage)
	}
	for _, elementList := range listOfMaxCliques {
		finalList = append(finalList, elementList)
	}
	return finalList
}

func isOptimal(listMax []Clique, graphPackage adjGraph.ConflictGraphPackage) bool {
	optimal := true
	toTest := make([]adjGraph.TrafficEntry, len(graphPackage.Entries))
	copy(toTest, graphPackage.Entries)
	for _, elementOuter := range listMax {
		for _, elementInner := range elementOuter {
			toTest, optimal = removeItemOfClique(toTest, elementInner)
			if !optimal {
				return optimal
			}
		}
	}
	for _, elementIdeal := range toTest {
		if elementIdeal.ChosenByUser == true {
			optimal = false
			break
		}
	}
	return optimal
}

func bronkerbIterative(current Clique, nextp Clique, prevEx Clique, graphPackage adjGraph.ConflictGraphPackage, listOfMaxClique []Clique) []Clique {
	stack := NewStack()
	stack = stack.Push(current, nextp, prevEx)
	for !stack.IsEmpty() {
		nstack, entry := stack.Pop()
		stack = nstack
		r := entry[0]
		p := entry[1]
		x := entry[2]
		if len(p) == 0 && len(x) == 0 {
			listOfMaxClique = append(listOfMaxClique, r)
		} else if len(p) != 0 {
			v := p[0]
			tmpp := copyRemoveItemOfClique(p, v)
			tmpx := copyAppend(x, v)
			stack = stack.Push(r, tmpp, tmpx)

			tmpc := copyAppend(r, v)
			nnc := getAllNeighbors(v, p, graphPackage)
			nnx := getAllNeighbors(v, x, graphPackage)
			stack = stack.Push(tmpc, nnc, nnx)

		}
	}
	return listOfMaxClique
}

func removeItemOfClique(clique Clique, entry adjGraph.TrafficEntry) (Clique, bool) {
	//get index of element
	//cut it out
	//return Clique
	i := -1
	everythingWentFine := true
	for index, element := range clique {
		if element == entry {
			i = index
			break
		}
	}
	if i == -1 {
		everythingWentFine = false
		return clique, everythingWentFine
	} else {
		copy(clique[i:], clique[i+1:])
		return clique[:len(clique)-1], everythingWentFine
	}
}

func copyRemoveItemOfClique(clique Clique, entry adjGraph.TrafficEntry) Clique {
	//get index of element
	//cut it out
	//return it in a new Clique
	i := -1

	tmp := make([]adjGraph.TrafficEntry, len(clique))
	copy(tmp, clique)
	for index, element := range tmp {
		if element == entry {
			i = index
			break
		}
	}
	copy(tmp[i:], tmp[i+1:])
	return tmp[:len(tmp)-1]
}

func copyAppend(clique Clique, entry adjGraph.TrafficEntry) Clique {
	tmp := make([]adjGraph.TrafficEntry, len(clique))
	copy(tmp, clique)
	tmp = append(tmp, entry)
	return tmp
}

func getAllNeighbors(element adjGraph.TrafficEntry, clique Clique, returnType adjGraph.ConflictGraphPackage) Clique {
	var neighbors Clique = make([]adjGraph.TrafficEntry, 0)
	//printClique(clique)
	//fmt.Print(" =clique")
	for _, elementClique := range clique {
		var indexElementClique int = adjGraph.GetIndexInConflictGraph(returnType, elementClique)
		var indexElement int = adjGraph.GetIndexInConflictGraph(returnType, element)
		if returnType.ConflictGraph.GetMatrixEntryAtIndex(indexElement, indexElementClique) {
			neighbors = append(neighbors, elementClique)
			//fmt.Println(elementClique)
		}
	}
	return neighbors
}

func IsNeighborOfThemAll(N Clique, entry adjGraph.TrafficEntry, returnType adjGraph.ConflictGraphPackage) bool {
	var pIsNeighborOfThemAll bool = true
	for _, element := range N {
		indexElementN := adjGraph.GetIndexInConflictGraph(returnType, element)
		indexElementP := adjGraph.GetIndexInConflictGraph(returnType, entry)
		if !returnType.ConflictGraph.GetMatrixEntryAtIndex(indexElementN, indexElementP) {
			pIsNeighborOfThemAll = false
			break
		}
	}

	return pIsNeighborOfThemAll
}
