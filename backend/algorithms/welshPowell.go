package algorithms

import (
	"fmt"
	"github.com/ob-algdatii-ss19/leistungsnachweis-ateam/backend/adjGraph"
)

/*
Calculate the optimization of the traffic lights with WelshPowell algorithm
*/
func WelshPowell(returnType adjGraph.ReturnType) [][]adjGraph.Node {
	fmt.Println("[INFO] Called WelshPowell Algorithm")

	var graphArray [][]int=graphTo2DimensionalArray(returnType)

	if graphArray == nil {
		return [][]adjGraph.Node{}
	}

	println("vor sortierung");
	for i := 0; i < len(graphArray); i++ {
		print("outputArray:", graphArray[i][0])
		for j := 1; j < len(graphArray[i]); j++ {
			print(" ",graphArray[i][j])
			print(" ")
		}
		println();
	}


	graphArray=sortNodesDescending(graphArray)

	println("NACH sortierung");
	for i := 0; i < len(graphArray); i++ {
		print("outputArray:", graphArray[i][0])
		for j := 1; j < len(graphArray[i]); j++ {
			print(" ",graphArray[i][j])
			print(" ")
		}
		println();
	}

	var coloredArray [][]int=giveColoredArray(graphArray)

	println("Färbung");
	for i := 0; i < len(coloredArray); i++ {
		print("outputArray:", coloredArray[i][0])
		for j := 1; j < len(coloredArray[i]); j++ {
			print(" ",coloredArray[i][j])
			print(" ")
		}
		println();
	}


	return nil
}

func graphTo2DimensionalArray(conGraph adjGraph.ReturnType) [][]int{
	graphData := conGraph.UGraph
	trafficEntries := conGraph.Entries

	if graphData == nil || trafficEntries == nil {
		return nil
	}

	var numberOfNodes = graphData.UNumberOfNodes()
	//var usedNodes []int

	var returnArray [][] int

	//make a Int Array out from the given conflict graph
	//go through all nodes (vertices)
	for i := 1; i <= numberOfNodes; i++ {

		//user selected the node via checkbox in the gui
		if trafficEntries[i-1].IsTrue {
			var innerArray=getAllConflictsOfThisNode(adjGraph.Node(i), graphData)
			returnArray=append(returnArray, innerArray)
		}
	}

	//println("getAllConflicts",returnArray )

	return returnArray
}

//first row in every column is actual node
func getAllConflictsOfThisNode(n adjGraph.Node, conGraph adjGraph.UGraph) []int{
	retArr := make([]int, 0)

	retArr = append(retArr, int(n))
	adjacentNodes := conGraph.UAdj(n)
	for i := 0; i < len(adjacentNodes); i++ {
		nodeNumber := adjacentNodes[i]
		retArr = append(retArr, int(nodeNumber))
	}


	//println("getAllConflictsOfThisNode",retArr )


	return retArr
}

//this function make the first step and order the nodes descending
//node with most edges is first, with less edges is last
func sortNodesDescending (nodeConflArray [][]int) [][]int{
	var result = nodeConflArray

	for i := 0; i < len(result); i++{
		//actElementOuter :=result[i];
		for j := (i+1); j <  len(result); j++{
			//actElementInner :=result[j];
			if(len(result[j])>len(result[i])){ //swap
				var tmp=result[j]
				result[j]=result[i];
				result[i]=tmp;

			}
		}
	}

	return result
}


//this function
func giveColoredArray(nodeConflArray [][]int) [][]int{

	var coloredArray [][]int

	var usedNodes []int =getUsedNodes(nodeConflArray)
	
	for i := 0; i < len(nodeConflArray); i++{

		var usedNodesThisRound=usedNodes; //which Nodes are possible? (when conflict in this round, delete element)

		//loop for all conflicts of the actual node
		for j := 0; j <  len(nodeConflArray[i]); j++{

			//loop over nodes with smaller weighting as actual node
			for k := i+1; k <  len(nodeConflArray[i]); j++{
				if(indexOf(nodeConflArray[i][0], usedNodesThisRound)>=0){
					continue; // this node is already in another phase
				}

				//is it possible, that this node can drive at same time a other node?

				//loop for all nodes in the actual compare-node
				for l := 0; l <  len(nodeConflArray[k]); l++{
					if(nodeConflArray[i][j] == nodeConflArray[k][l]){
						usedNodesThisRound=findAndRemove(k,usedNodesThisRound) //this node cannot drive at same time
						break; //so end this loop
					}
				}
			}
		}


		//delete this node, because this street cannot drive anymore
		for i := 0; i < len(nodeConflArray); i++ {
			usedNodes=findAndRemove(usedNodesThisRound[i],usedNodes)
		}

		//add all nodes with same color
		coloredArray=append(coloredArray, usedNodesThisRound)
	}

	println("in Färbung", len(coloredArray))
	return coloredArray
}

//returna all used nodes
//is always index 0 in nodeConflArray
func getUsedNodes(nodeConflArray [][]int)[]int{
	var usedNodes []int
	for i := 0; i < len(nodeConflArray); i++{
		usedNodes=append(usedNodes, nodeConflArray[i][0])
	}
	return usedNodes
}


//array help functions

//find element and delete this
func findAndRemove(element int, data []int) []int{
	var pos int=indexOf(element, data)
	if(pos<0){
		return data
	}
	return remove(data, pos)
}


func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1    //not found.
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}



//siehe http://mrsleblancsmath.pbworks.com/w/file/fetch/46119304/vertex%20coloring%20algorithm.pdf]
//Konfliktgraph
//drüber iterieren nicht erforderlich, siehe ugraph::UAdj
//Knoten mit meisten Konflikten fahren zuerst
//keine Gruppen bilden!!!
//dann siehe Video, nur AEHNLICH zu basic greedy
