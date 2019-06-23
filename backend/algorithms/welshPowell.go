package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with WelshPowell algorithm
*/
func WelshPowell(returnType adjGraph.ConflictGraphPackage) [][]adjGraph.Node {
	fmt.Println("[INFO] Called WelshPowell Algorithm")

	var graphArray [][]int = graphTo2DimensionalArray(returnType)

	if graphArray == nil {
		return [][]adjGraph.Node{}
	}

	println("vor sortierung")
	for i := 0; i < len(graphArray); i++ {
		print("outputArray:", graphArray[i][0])
		for j := 1; j < len(graphArray[i]); j++ {
			print(" ", graphArray[i][j])
			print(" ")
		}
		println()
	}

	graphArray = sortNodesDescending(graphArray)

	var coloredArray [][]int = giveColoredArray(graphArray)

	return intArrayToNodeArray(coloredArray)
}

func graphTo2DimensionalArray(conGraph adjGraph.ConflictGraphPackage) [][]int {
	graphData := conGraph.ConflictGraph
	trafficEntries := conGraph.Entries

	if graphData == nil || trafficEntries == nil {
		return nil
	}

	var numberOfNodes = graphData.UNumberOfNodes()

	var returnArray [][]int

	//make a Int Array out from the given conflict graph
	//go through all nodes (vertices)
	for i := 1; i <= numberOfNodes; i++ {

		//user selected the node via checkbox in the gui
		if trafficEntries[i-1].ChosenByUser {
			var innerArray = getAllConflictsOfThisNode(adjGraph.Node(i), graphData)
			returnArray = append(returnArray, innerArray)
		}
	}

	return returnArray
}

//first row in every column is actual node
func getAllConflictsOfThisNode(n adjGraph.Node, conGraph adjGraph.UGraph) []int {
	retArr := make([]int, 0)

	retArr = append(retArr, int(n))
	adjacentNodes := conGraph.UAdj(n)
	for i := 0; i < len(adjacentNodes); i++ {
		nodeNumber := adjacentNodes[i]
		retArr = append(retArr, int(nodeNumber))
	}

	return retArr
}

//this function make the first step and order the nodes descending
//node with most edges is first, with less edges is last
func sortNodesDescending(nodeConflArray [][]int) [][]int {
	var result = nodeConflArray

	for i := 0; i < len(result); i++ {
		//actElementOuter :=result[i];
		for j := (i + 1); j < len(result); j++ {
			//actElementInner :=result[j];
			if len(result[j]) > len(result[i]) { //swap
				var tmp = result[j]
				result[j] = result[i]
				result[i] = tmp

			}
		}
	}

	return result
}

//this function
func giveColoredArray(nodeConflArray [][]int) [][]int {

	var coloredArray [][]int // = make([][]int, 0,0)

	var usedNodes []int = getUsedNodes(nodeConflArray)
	/*for j := 0; j < len(usedNodes); j++ {
		print(" ",usedNodes[j])
		print(" ")
	}
	println("")*/

	for i := 0; i < len(nodeConflArray); i++ {

		if len(usedNodes) == 0 { //all nodes where used
			break
		}

		//copy is very important, when usedNodesThisRound=usedNodes, usedNodes will changes
		usedNodesThisRound := make([]int, len(usedNodes)) //which Nodes are possible? (when conflict in this round, delete element)
		copy(usedNodesThisRound, usedNodes)

		//loop for all conflicts of the actual node
		for j := 0; j < len(nodeConflArray[i]); j++ {

			//loop over nodes with smaller weighting as actual node
			for k := i + 1; k < len(nodeConflArray); k++ {
				if nodeConflArray[i][j] == nodeConflArray[k][0] {

					usedNodesThisRound = findAndRemove(nodeConflArray[k][0], usedNodesThisRound) //this nodes are not allowed at same time

				} else {
					continue // this node cannot find in this part of array
				}

				if len(usedNodesThisRound) == 0 { //all deleted this round
					break
				}
			}
		}

		//now delete Nodes, wich depends at subnode
		for x := 0; x < len(usedNodesThisRound); x++ {
			var indexIn2DimArray = findIndexIn2DimArray(usedNodesThisRound[x], nodeConflArray)

			//all conflicts of this node
			for y := 1; y < len(nodeConflArray[indexIn2DimArray]); y++ {
				for z := y; z < len(nodeConflArray); z++ {
					if nodeConflArray[indexIn2DimArray][y] == nodeConflArray[z][0] {
						usedNodesThisRound = findAndRemove(nodeConflArray[z][0], usedNodesThisRound) //this nodes are not allowed at same time
					}
				}
			}
		}

		for x := 0; x < len(usedNodesThisRound); x++ {
			var searcInt = usedNodesThisRound[x]
			usedNodes = findAndRemove(searcInt, usedNodes)
		}

		//try to group with other phase
		var added bool = false //added something?
		for colored := 0; colored < len(coloredArray); colored++ {
			var hasConflict bool = false
			for coloredSub := 0; coloredSub < len(coloredArray[colored]); coloredSub++ { //to compare with all nodes, belongs a phase
				var indexIn2DimArray = findIndexIn2DimArray(coloredArray[colored][coloredSub], nodeConflArray)
				for y := 1; y < len(nodeConflArray[indexIn2DimArray]); y++ {
					for z := y + 1; z < len(nodeConflArray); z++ {
						for useInThisRound := 0; useInThisRound < len(usedNodesThisRound); useInThisRound++ {
							var useInThisRoundVal = usedNodesThisRound[useInThisRound]
							if nodeConflArray[indexIn2DimArray][y] == useInThisRoundVal {
								hasConflict = true
							}
						}
					}
				}
			}

			if !hasConflict {
				var innerAr = coloredArray[colored]
				for useInThisRound := 0; useInThisRound < len(usedNodesThisRound); useInThisRound++ {
					var useInThisRoundVal = usedNodesThisRound[useInThisRound]
					innerAr = append(innerAr, useInThisRoundVal)
				}
				coloredArray[colored] = innerAr

				added = true
			}

		}

		if !added { //only, when usedNodesThisRound not added at other index
			coloredArray = append(coloredArray, usedNodesThisRound)
		}

		//add all nodes with same color
	}

	//println("in Färbung", len(coloredArray))
	return coloredArray
}

//return all used nodes
//is always index 0 in nodeConflArray
func getUsedNodes(nodeConflArray [][]int) []int {
	usedNodes := make([]int, 0)
	for i := 0; i < len(nodeConflArray); i++ {
		usedNodes = append(usedNodes, nodeConflArray[i][0])
	}
	return usedNodes
}

//array help functions

//find element and delete this
func findAndRemove(element int, data []int) []int {
	var pos int = indexOf(element, data)
	//print("pos", pos)
	if pos < 0 {
		return data
	}
	return remove(data, pos)
}

func indexOf(element int, data []int) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

func findIndexIn2DimArray(element int, data [][]int) int {
	for i := 0; i < len(data); i++ {
		if data[i][0] == element {
			return i
		}
	}
	return -1
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func intArrayToNodeArray(intConflArray [][]int) [][]adjGraph.Node {
	var nodeConflArray [][]adjGraph.Node

	for i := 0; i < len(intConflArray); i++ {
		var actInnerIntArray = intConflArray[i]
		var actInnerNodeArray []adjGraph.Node

		for j := 0; j < len(intConflArray[i]); j++ {
			var n adjGraph.Node = adjGraph.Node(actInnerIntArray[j])
			actInnerNodeArray = append(actInnerNodeArray, n)
		}
		if len(actInnerNodeArray) > 0 { //insert no empty elements
			nodeConflArray = append(nodeConflArray, actInnerNodeArray)
		}
	}

	return nodeConflArray
}
