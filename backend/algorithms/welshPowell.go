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

	/*println("vor sortierung");
	for i := 0; i < len(graphArray); i++ {
		print("outputArray:", graphArray[i][0])
		for j := 1; j < len(graphArray[i]); j++ {
			print(" ",graphArray[i][j])
			print(" ")
		}
		println();
	}*/


	graphArray=sortNodesDescending(graphArray)

	/*println("NACH sortierung");
	for i := 0; i < len(graphArray); i++ {
		print("outputArray:", graphArray[i][0])
		for j := 1; j < len(graphArray[i]); j++ {
			print(" ",graphArray[i][j])
			print(" ")
		}
		println();
	}*/

	var coloredArray [][]int=giveColoredArray(graphArray)

	/*println("Färbung");
	for i := 0; i < len(coloredArray); i++ {
		print("outputArray:", coloredArray[i][0])
		for j := 1; j < len(coloredArray[i]); j++ {
			print(" ",coloredArray[i][j])
			print(" ")
		}
		println();
	}*/


	return intArrayToNodeArray(coloredArray)
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

	var coloredArray [][]int;// = make([][]int, 0,0)

	var usedNodes []int =getUsedNodes(nodeConflArray)
	/*for j := 0; j < len(usedNodes); j++ {
		print(" ",usedNodes[j])
		print(" ")
	}
	println("")*/

	for i := 0; i < len(nodeConflArray); i++{

		if(len(usedNodes)==0){ //all nodes where used
			break;
		}

		//copy is very important, when usedNodesThisRound=usedNodes, usedNodes will changes
		usedNodesThisRound := make([]int, len(usedNodes))//which Nodes are possible? (when conflict in this round, delete element)
		copy(usedNodesThisRound, usedNodes)


		//loop for all conflicts of the actual node
		for j := 0; j <  len(nodeConflArray[i]); j++{
			
			//loop over nodes with smaller weighting as actual node
			for k := i+1; k <  len(nodeConflArray[i]); k++{
				if(nodeConflArray[i][j] == nodeConflArray[k][0]){

					usedNodesThisRound=findAndRemove(nodeConflArray[k][0],usedNodesThisRound) //this nodes are not allowed at same time //ÄNERUNG!

					break;
				}else{
					continue; // this node cannot find in this part of array
				}

				if(len(usedNodesThisRound)==0){ //all deleted this round
					break;
				}
			}
		}

		for x := 0; x < len(usedNodesThisRound); x++ {
			//print(" ",usedNodesThisRound[x])
			//print(" ")
			var searcInt =usedNodesThisRound[x];
			usedNodes=findAndRemove(searcInt,usedNodes)
		}


		//add all nodes with same color
		coloredArray=append(coloredArray, usedNodesThisRound)

	}

	//println("in Färbung", len(coloredArray))
	return coloredArray
}

//returna all used nodes
//is always index 0 in nodeConflArray
func getUsedNodes(nodeConflArray [][]int) []int{
	usedNodes := make([]int, 0)
	for i := 0; i < len(nodeConflArray); i++{
		usedNodes=append(usedNodes, nodeConflArray[i][0])
	}
	return usedNodes
}


//array help functions

//find element and delete this
func findAndRemove(element int, data []int) []int{
	var pos int=indexOf(element, data)
	//print("pos", pos)
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



func intArrayToNodeArray(intConflArray [][]int) [][]adjGraph.Node{
	var nodeConflArray [][]adjGraph.Node

	for i := 0; i < len(intConflArray); i++{
		var actInnerIntArray=intConflArray[i]
		var actInnerNodeArray []adjGraph.Node

		for j := 0; j < len(intConflArray[i]); j++{
			var n adjGraph.Node=adjGraph.Node(actInnerIntArray[j])
			actInnerNodeArray = append(actInnerNodeArray, n)
		}
		if(len(actInnerNodeArray)>0){ //insert no empty elements
			nodeConflArray = append(nodeConflArray, actInnerNodeArray)
		}
	}

	return nodeConflArray
}


//siehe http://mrsleblancsmath.pbworks.com/w/file/fetch/46119304/vertex%20coloring%20algorithm.pdf]
//Konfliktgraph
//drüber iterieren nicht erforderlich, siehe ugraph::UAdj
//Knoten mit meisten Konflikten fahren zuerst
//keine Gruppen bilden!!!
//dann siehe Video, nur AEHNLICH zu basic greedy
